package root

import (
	"github.com/iotaledger/wasp/packages/iscp/coreutil"
	"github.com/iotaledger/wasp/packages/vm/core/errors/commonerrors"
)

var (
	Contract = coreutil.NewContract(coreutil.CoreContractRoot, "Root Contract")
)

// state variables
const (
	StateVarContractRegistry         = "r"
	StateVarDeployPermissionsEnabled = "a"
	StateVarDeployPermissions        = "p"
	StateVarStateInitialized         = "i"
)

// param variables
const (
	ParamDeployer                  = "dp"
	ParamHname                     = "hn"
	ParamName                      = "nm"
	ParamProgramHash               = "ph"
	ParamContractRecData           = "dt"
	ParamContractFound             = "cf"
	ParamDescription               = "ds"
	ParamDeployPermissionsEnabled  = "de"
	ParamDustDepositAssumptionsBin = "db"
)

// function names
var (
	FuncDeployContract           = coreutil.Func("deployContract")
	FuncGrantDeployPermission    = coreutil.Func("grantDeployPermission")
	FuncRevokeDeployPermission   = coreutil.Func("revokeDeployPermission")
	FuncRequireDeployPermissions = coreutil.Func("requireDeployPermissions")
	FuncFindContract             = coreutil.ViewFunc("findContract")
	FuncGetContractRecords       = coreutil.ViewFunc("getContractRecords")
)

var (
	ErrChainInitConditionsFailed = commonerrors.RegisterGlobalError("root.init can't be called in this state").CreateTyped()
)
