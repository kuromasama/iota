package vmtxbuilder

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/iotaledger/wasp/packages/vm"

	"github.com/ethereum/go-ethereum/accounts/abi"
	iotago "github.com/iotaledger/iota.go/v3"
	"github.com/iotaledger/wasp/packages/iscp"
	"github.com/iotaledger/wasp/packages/parameters"
	"github.com/iotaledger/wasp/packages/transaction"
	"github.com/iotaledger/wasp/packages/util"
	"github.com/iotaledger/wasp/packages/vm/vmcontext/vmexceptions"
	"golang.org/x/xerrors"
)

// tokenOutputLoader externally supplied function which loads stored output from the state
// Should return nil if does not exist
type tokenOutputLoader func(*iotago.NativeTokenID) (*iotago.BasicOutput, *iotago.UTXOInput)

// foundryLoader externally supplied function which returns foundry output and id by its serial number
// Should return nil if foundry does not exist
type foundryLoader func(uint32) (*iotago.FoundryOutput, *iotago.UTXOInput)

// NFTOutputLoader externally supplied function which returns the stored NFT output from the state
// Should return nil if NFT is not accounted for
type NFTOutputLoader func(id *iotago.NFTID) (*iotago.NFTOutput, *iotago.UTXOInput)

// AnchorTransactionBuilder represents structure which handles all the data needed to eventually
// build an essence of the anchor transaction
type AnchorTransactionBuilder struct {
	// anchorOutput output of the chain
	anchorOutput *iotago.AliasOutput
	// anchorOutputID is the ID of the anchor output
	anchorOutputID iotago.OutputID
	// already consumed outputs, specified by entire Request. It is needed for checking validity
	consumed []iscp.Request
	// iotas which are on-chain. It does not include dust deposits on anchor and on internal outputs
	totalIotasInL2Accounts uint64
	// minimum dust deposit assumption for internal outputs. It is used as constants. Assumed real dust cost never grows
	dustDepositAssumption transaction.DustDepositAssumption
	// balance loader for native tokens
	loadTokenOutput tokenOutputLoader
	// foundry loader
	loadFoundry foundryLoader
	// NFToutput loader
	loadNFTOutput NFTOutputLoader
	// balances of native tokens loaded during the batch run
	balanceNativeTokens map[iotago.NativeTokenID]*nativeTokenBalance
	// all nfts loaded during the batch run
	nftsIncluded map[iotago.NFTID]*nftIncluded
	// invoked foundries. Foundry serial number is used as a key
	invokedFoundries map[uint32]*foundryInvoked
	// requests posted by smart contracts
	postedOutputs []iotago.Output
	// parameters coming from the L1 node
	l1Params *parameters.L1
}

// NewAnchorTransactionBuilder creates new AnchorTransactionBuilder object
func NewAnchorTransactionBuilder(
	anchorOutput *iotago.AliasOutput,
	anchorOutputID iotago.OutputID,
	tokenBalanceLoader tokenOutputLoader,
	foundryLoader foundryLoader,
	nftLoader NFTOutputLoader,
	dustDepositAssumptions transaction.DustDepositAssumption,
	l1Params *parameters.L1,
) *AnchorTransactionBuilder {
	if anchorOutput.Amount < dustDepositAssumptions.AnchorOutput {
		panic("internal inconsistency")
	}
	return &AnchorTransactionBuilder{
		anchorOutput:           anchorOutput,
		anchorOutputID:         anchorOutputID,
		totalIotasInL2Accounts: anchorOutput.Amount - dustDepositAssumptions.AnchorOutput,
		dustDepositAssumption:  dustDepositAssumptions,
		loadTokenOutput:        tokenBalanceLoader,
		loadFoundry:            foundryLoader,
		loadNFTOutput:          nftLoader,
		consumed:               make([]iscp.Request, 0, iotago.MaxInputsCount-1),
		balanceNativeTokens:    make(map[iotago.NativeTokenID]*nativeTokenBalance),
		postedOutputs:          make([]iotago.Output, 0, iotago.MaxOutputsCount-1),
		invokedFoundries:       make(map[uint32]*foundryInvoked),
		nftsIncluded:           make(map[iotago.NFTID]*nftIncluded),
		l1Params:               l1Params,
	}
}

// Clone clones the AnchorTransactionBuilder object. Used to snapshot/recover
func (txb *AnchorTransactionBuilder) Clone() *AnchorTransactionBuilder {
	ret := &AnchorTransactionBuilder{
		anchorOutput:           txb.anchorOutput,
		anchorOutputID:         txb.anchorOutputID,
		totalIotasInL2Accounts: txb.totalIotasInL2Accounts,
		dustDepositAssumption:  txb.dustDepositAssumption,
		loadTokenOutput:        txb.loadTokenOutput,
		loadFoundry:            txb.loadFoundry,
		consumed:               make([]iscp.Request, 0, cap(txb.consumed)),
		balanceNativeTokens:    make(map[iotago.NativeTokenID]*nativeTokenBalance),
		postedOutputs:          make([]iotago.Output, 0, cap(txb.postedOutputs)),
		invokedFoundries:       make(map[uint32]*foundryInvoked),
		nftsIncluded:           make(map[iotago.NFTID]*nftIncluded),
		l1Params:               txb.l1Params,
	}

	ret.consumed = append(ret.consumed, txb.consumed...)
	for k, v := range txb.balanceNativeTokens {
		ret.balanceNativeTokens[k] = v.clone()
	}
	for k, v := range txb.invokedFoundries {
		ret.invokedFoundries[k] = v.clone()
	}
	for k, v := range txb.nftsIncluded {
		ret.nftsIncluded[k] = v.clone()
	}
	ret.postedOutputs = append(ret.postedOutputs, txb.postedOutputs...)
	return ret
}

// TotalIotasInL2Accounts returns number of on-chain iotas.
// It does not include minimum dust deposit needed for anchor output and other internal UTXOs
func (txb *AnchorTransactionBuilder) TotalIotasInL2Accounts() uint64 {
	return txb.totalIotasInL2Accounts
}

// Consume adds an input to the transaction.
// It panics if transaction cannot hold that many inputs
// All explicitly consumed inputs will hold fixed index in the transaction
// It updates total assets held by the chain. So it may panic due to exceed output counts
// Returns delta of iotas needed to adjust the common account due to dust deposit requirement for internal UTXOs
// NOTE: if call panics with ErrNotEnoughFundsForInternalDustDeposit, the state of the builder becomes inconsistent
// It means, in the caller context it should be rolled back altogether
func (txb *AnchorTransactionBuilder) Consume(req iscp.Request) int64 {
	if DebugTxBuilder {
		txb.MustBalanced("txbuilder.Consume IN")
	}
	if req.IsOffLedger() {
		panic("txbuilder.Consume: must be UTXO")
	}
	if txb.InputsAreFull() {
		panic(vmexceptions.ErrInputLimitExceeded)
	}

	defer txb.mustCheckTotalNativeTokensExceeded()

	txb.consumed = append(txb.consumed, req)

	// first we add all iotas arrived with the output to anchor balance
	txb.addDeltaIotasToTotal(req.AsOnLedger().Output().Deposit())
	// then we add all arriving native tokens to corresponding internal outputs
	deltaIotasDustDepositAdjustment := int64(0)
	for _, nt := range req.Assets().Tokens {
		deltaIotasDustDepositAdjustment += txb.addNativeTokenBalanceDelta(&nt.ID, nt.Amount)
	}
	if DebugTxBuilder {
		txb.MustBalanced("txbuilder.Consume OUT")
	}
	if req.NFT() != nil {
		txb.includeNFT(req.AsOnLedger().Output().(*iotago.NFTOutput))
	}
	return deltaIotasDustDepositAdjustment
}

// AddOutput adds an information about posted request. It will produce output
// Return adjustment needed for the L2 ledger (adjustment on iotas related to dust protection)
func (txb *AnchorTransactionBuilder) AddOutput(o iotago.Output) int64 {
	if txb.outputsAreFull() {
		panic(vmexceptions.ErrOutputLimitExceeded)
	}

	defer txb.mustCheckTotalNativeTokensExceeded()

	requiredDustDeposit := o.VByteCost(txb.l1Params.RentStructure(), nil)
	if o.Deposit() < requiredDustDeposit {
		panic(xerrors.Errorf("%v: available %d < required %d iotas",
			transaction.ErrNotEnoughIotasForDustDeposit, o.Deposit(), requiredDustDeposit))
	}
	assets := transaction.AssetsFromOutput(o)
	txb.subDeltaIotasFromTotal(assets.Iotas)
	bi := new(big.Int)
	iotaAdjustmentL2 := int64(0)
	for _, nt := range assets.Tokens {
		bi.Neg(nt.Amount)
		iotaAdjustmentL2 += txb.addNativeTokenBalanceDelta(&nt.ID, bi)
	}
	if nftout, ok := o.(*iotago.NFTOutput); ok {
		txb.includeNFT(nftout)
	}
	txb.postedOutputs = append(txb.postedOutputs, o)
	return iotaAdjustmentL2
}

// InputsAreFull returns if transaction cannot hold more inputs
func (txb *AnchorTransactionBuilder) InputsAreFull() bool {
	return txb.numInputs() >= iotago.MaxInputsCount
}

// BuildTransactionEssence builds transaction essence from tx builder data
func (txb *AnchorTransactionBuilder) BuildTransactionEssence(stateData *iscp.StateData) (*iotago.TransactionEssence, []byte) {
	txb.MustBalanced("BuildTransactionEssence IN")
	inputs, inputIDs := txb.inputs()
	essence := &iotago.TransactionEssence{
		NetworkID: txb.l1Params.NetworkID,
		Inputs:    inputIDs.UTXOInputs(),
		Outputs:   txb.outputs(stateData),
		Payload:   nil,
	}
	return essence, inputIDs.OrderedSet(inputs).MustCommitment()
}

// inputIDs generates a deterministic list of inputs for the transaction essence
// - index 0 is always alias output
// - then goes consumed external BasicOutput UTXOs, the requests, in the order of processing
// - then goes inputs of native token UTXOs, sorted by token id
// - then goes inputs of foundries sorted by serial number
func (txb *AnchorTransactionBuilder) inputs() (iotago.OutputSet, iotago.OutputIDs) {
	ids := make(iotago.OutputIDs, 0, len(txb.consumed)+len(txb.balanceNativeTokens)+len(txb.invokedFoundries))
	inputs := make(iotago.OutputSet, 0)

	// alias output
	ids = append(ids, txb.anchorOutputID)
	inputs[txb.anchorOutputID] = txb.anchorOutput

	// consumed on-ledger requests
	for i := range txb.consumed {
		id := txb.consumed[i].ID().OutputID()
		ids = append(ids, id)
		inputs[id] = txb.consumed[i].AsOnLedger().Output()
	}

	// internal native token outputs
	for _, nt := range txb.nativeTokenOutputsSorted() {
		if nt.requiresInput() {
			id := nt.input.ID()
			ids = append(ids, id)
			inputs[id] = nt.in
		}
	}

	// foundries
	for _, f := range txb.foundriesSorted() {
		if f.requiresInput() {
			id := f.input.ID()
			ids = append(ids, id)
			inputs[id] = f.in
		}
	}

	// nfts
	for _, nft := range txb.nftsSorted() {
		if nft.in != nil {
			id := nft.input.ID()
			ids = append(ids, id)
			inputs[id] = nft.in
		}
	}

	if len(ids) != txb.numInputs() {
		panic("AnchorTransactionBuilder.inputs: internal inconsistency")
	}
	return inputs, ids
}

// outputs generates outputs for the transaction essence
func (txb *AnchorTransactionBuilder) outputs(stateData *iscp.StateData) iotago.Outputs {
	ret := make(iotago.Outputs, 0, 1+len(txb.balanceNativeTokens)+len(txb.postedOutputs))
	// creating the anchor output
	aliasID := txb.anchorOutput.AliasID
	if aliasID.Empty() {
		aliasID = iotago.AliasIDFromOutputID(txb.anchorOutputID)
	}
	anchorOutput := &iotago.AliasOutput{
		Amount:         txb.totalIotasInL2Accounts + txb.dustDepositAssumption.AnchorOutput,
		NativeTokens:   nil, // anchor output does not contain native tokens
		AliasID:        aliasID,
		StateIndex:     txb.anchorOutput.StateIndex + 1,
		StateMetadata:  stateData.Bytes(),
		FoundryCounter: txb.nextFoundryCounter(),
		Conditions: iotago.UnlockConditions{
			&iotago.StateControllerAddressUnlockCondition{Address: txb.anchorOutput.StateController()},
			&iotago.GovernorAddressUnlockCondition{Address: txb.anchorOutput.GovernorAddress()},
		},
		Blocks: iotago.FeatureBlocks{
			&iotago.SenderFeatureBlock{
				Address: aliasID.ToAddress(),
			},
		},
	}
	ret = append(ret, anchorOutput)

	// creating outputs for updated internal accounts
	nativeTokensToBeUpdated, _ := txb.NativeTokenRecordsToBeUpdated()
	for _, id := range nativeTokensToBeUpdated {
		// create one output for each token ID of internal account
		ret = append(ret, txb.balanceNativeTokens[id].out)
	}
	// creating outputs for updated foundries
	foundriesToBeUpdated, _ := txb.FoundriesToBeUpdated()
	for _, sn := range foundriesToBeUpdated {
		ret = append(ret, txb.invokedFoundries[sn].out)
	}
	// creating outputs for posted on-ledger requests
	ret = append(ret, txb.postedOutputs...)
	return ret
}

// numInputs number of inputs in the future transaction
func (txb *AnchorTransactionBuilder) numInputs() int {
	ret := len(txb.consumed) + 1 // + 1 for anchor UTXO
	for _, v := range txb.balanceNativeTokens {
		if v.requiresInput() {
			ret++
		}
	}
	for _, f := range txb.invokedFoundries {
		if f.requiresInput() {
			ret++
		}
	}
	return ret
}

// numOutputs in the transaction
func (txb *AnchorTransactionBuilder) numOutputs() int {
	ret := 1 // for chain output
	for _, v := range txb.balanceNativeTokens {
		if v.producesOutput() {
			ret++
		}
	}
	ret += len(txb.postedOutputs)
	for _, f := range txb.invokedFoundries {
		if f.producesOutput() {
			ret++
		}
	}
	return ret
}

// outputsAreFull return if transaction cannot bear more outputs
func (txb *AnchorTransactionBuilder) outputsAreFull() bool {
	return txb.numOutputs() >= iotago.MaxOutputsCount
}

func (txb *AnchorTransactionBuilder) mustCheckTotalNativeTokensExceeded() {
	num := 0
	for _, nt := range txb.balanceNativeTokens {
		if nt.requiresInput() || nt.producesOutput() {
			num++
		}
		if num > iotago.MaxNativeTokensCount {
			panic(vmexceptions.ErrTotalNativeTokensLimitExceeded)
		}
	}
}

// addDeltaIotasToTotal increases number of on-chain main account iotas by delta
func (txb *AnchorTransactionBuilder) addDeltaIotasToTotal(delta uint64) {
	if delta == 0 {
		return
	}
	// safe arithmetics
	n := txb.totalIotasInL2Accounts + delta
	if n+txb.dustDepositAssumption.AnchorOutput < txb.totalIotasInL2Accounts {
		panic(xerrors.Errorf("addDeltaIotasToTotal: %w", vm.ErrOverflow))
	}
	txb.totalIotasInL2Accounts = n
}

// subDeltaIotasFromTotal decreases number of on-chain main account iotas
func (txb *AnchorTransactionBuilder) subDeltaIotasFromTotal(delta uint64) {
	if delta == 0 {
		return
	}
	// safe arithmetics
	if delta > txb.totalIotasInL2Accounts {
		panic(vm.ErrNotEnoughIotaBalance)
	}
	txb.totalIotasInL2Accounts -= delta
}

// ensureNativeTokenBalance makes sure that cached output is in the builder
// if not, it asks for the in balance by calling the loader function
// Panics if the call results to exceeded limits
func (txb *AnchorTransactionBuilder) ensureNativeTokenBalance(id *iotago.NativeTokenID) *nativeTokenBalance {
	if b, ok := txb.balanceNativeTokens[*id]; ok {
		return b
	}
	in, input := txb.loadTokenOutput(id) // output will be nil if no such token id accounted yet
	if in != nil && txb.InputsAreFull() {
		panic(vmexceptions.ErrInputLimitExceeded)
	}
	if in != nil && txb.outputsAreFull() {
		panic(vmexceptions.ErrOutputLimitExceeded)
	}

	var out *iotago.BasicOutput
	if in == nil {
		out = txb.newInternalTokenOutput(txb.anchorOutput.AliasID, *id)
	} else {
		out = cloneInternalBasicOutputOrNil(in)
	}
	b := &nativeTokenBalance{
		tokenID:            out.NativeTokens[0].ID,
		in:                 in,
		out:                out,
		dustDepositCharged: in != nil,
	}
	if input != nil {
		b.input = *input
	}
	txb.balanceNativeTokens[*id] = b
	return b
}

// addNativeTokenBalanceDelta adds delta to the token balance. Use negative delta to subtract.
// The call may result in adding new token ID to the ledger or disappearing one
// This impacts dust amount locked in the internal UTXOs which keep respective balances
// Returns delta of required dust deposit
func (txb *AnchorTransactionBuilder) addNativeTokenBalanceDelta(id *iotago.NativeTokenID, delta *big.Int) int64 {
	if util.IsZeroBigInt(delta) {
		return 0
	}
	nt := txb.ensureNativeTokenBalance(id)
	tmp := new(big.Int).Add(nt.getOutValue(), delta)
	if tmp.Sign() < 0 {
		panic(xerrors.Errorf("addNativeTokenBalanceDelta (id: %s, delta: %d): %v",
			id, delta, vm.ErrNotEnoughNativeAssetBalance))
	}
	if tmp.Cmp(abi.MaxUint256) > 0 {
		panic(xerrors.Errorf("addNativeTokenBalanceDelta: %v", vm.ErrOverflow))
	}
	nt.setOutValue(tmp)
	switch {
	case nt.identicalInOut():
		return 0
	case nt.dustDepositCharged && !nt.producesOutput():
		// this is an old token in the on-chain ledger. Now it disappears and dust deposit
		// is released and delta of anchor is positive
		nt.dustDepositCharged = false
		txb.addDeltaIotasToTotal(txb.dustDepositAssumption.NativeTokenOutput)
		return int64(txb.dustDepositAssumption.NativeTokenOutput)
	case !nt.dustDepositCharged && nt.producesOutput():
		// this is a new token in the on-chain ledger
		// There's a need for additional dust deposit on the respective UTXO, so delta for the anchor is negative
		nt.dustDepositCharged = true
		if txb.dustDepositAssumption.NativeTokenOutput > txb.totalIotasInL2Accounts {
			panic(vmexceptions.ErrNotEnoughFundsForInternalDustDeposit)
		}
		txb.subDeltaIotasFromTotal(txb.dustDepositAssumption.NativeTokenOutput)
		return -int64(txb.dustDepositAssumption.NativeTokenOutput)
	}
	return 0
}

func (txb *AnchorTransactionBuilder) includeNFT(o *iotago.NFTOutput) {
	if txb.nftsIncluded[o.NFTID] != nil {
		// NFT comes in and out in the same block
		txb.nftsIncluded[o.NFTID].sentOutside = true
		txb.nftsIncluded[o.NFTID].out = o
		return
	}
	in, input := txb.loadNFTOutput(&o.NFTID) // output will be nil if no such token id accounted yet
	if in != nil && txb.InputsAreFull() {
		panic(vmexceptions.ErrInputLimitExceeded)
	}
	if in != nil && txb.outputsAreFull() {
		panic(vmexceptions.ErrOutputLimitExceeded)
	}

	b := &nftIncluded{
		ID:                 o.NFTID,
		in:                 in,
		out:                o,
		dustDepositCharged: in != nil,
		sentOutside:        false,
	}
	if input != nil {
		b.input = *input
	}

	txb.nftsIncluded[o.NFTID] = b
}

func stringUTXOInput(inp *iotago.UTXOInput) string {
	return fmt.Sprintf("[%d]%s", inp.TransactionOutputIndex, hex.EncodeToString(inp.TransactionID[:]))
}

func stringNativeTokenID(id *iotago.NativeTokenID) string {
	return hex.EncodeToString(id[:])
}

func (txb *AnchorTransactionBuilder) String() string {
	ret := ""
	ret += fmt.Sprintf("%s\n", stringUTXOInput(txb.anchorOutputID.UTXOInput()))
	ret += fmt.Sprintf("in IOTA balance: %d\n", txb.anchorOutput.Amount)
	ret += fmt.Sprintf("current IOTA balance: %d\n", txb.totalIotasInL2Accounts)
	ret += fmt.Sprintf("Native tokens (%d):\n", len(txb.balanceNativeTokens))
	for id, ntb := range txb.balanceNativeTokens {
		initial := "0"
		if ntb.in != nil {
			initial = ntb.getOutValue().String()
		}
		current := ntb.getOutValue().String()
		ret += fmt.Sprintf("      %s: %s --> %s, dust deposit charged: %v\n",
			stringNativeTokenID(&id), initial, current, ntb.dustDepositCharged)
	}
	ret += fmt.Sprintf("consumed inputs (%d):\n", len(txb.consumed))
	//for _, inp := range txb.consumed {
	//	ret += fmt.Sprintf("      %s\n", inp.ID().String())
	//}
	ret += fmt.Sprintf("added outputs (%d):\n", len(txb.postedOutputs))
	ret += ">>>>>> TODO. Not finished....."
	return ret
}
