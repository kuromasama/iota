// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmtypes from "wasmlib/wasmtypes";

export const ScName        = "accounts";
export const ScDescription = "Chain account ledger contract";
export const HScName       = new wasmtypes.ScHname(0x3c4b5e02);

export const ParamAgentID           = "a";
export const ParamDestroyTokens     = "y";
export const ParamForceMinimumIotas = "f";
export const ParamForceOpenAccount  = "c";
export const ParamFoundrySN         = "s";
export const ParamNftID             = "z";
export const ParamSupplyDeltaAbs    = "d";
export const ParamTokenScheme       = "t";
export const ParamTokenTag          = "g";

export const ResultAccountNonce     = "n";
export const ResultAllAccounts      = "this";
export const ResultAssets           = "this";
export const ResultBalances         = "this";
export const ResultFoundryOutputBin = "b";
export const ResultFoundrySN        = "s";
export const ResultMapping          = "G";
export const ResultNftData          = "e";
export const ResultNftIDs           = "i";

export const FuncDeposit                  = "deposit";
export const FuncFoundryCreateNew         = "foundryCreateNew";
export const FuncFoundryDestroy           = "foundryDestroy";
export const FuncFoundryModifySupply      = "foundryModifySupply";
export const FuncHarvest                  = "harvest";
export const FuncTransferAllowanceTo      = "transferAllowanceTo";
export const FuncWithdraw                 = "withdraw";
export const ViewAccountNFTs              = "accountNFTs";
export const ViewAccounts                 = "accounts";
export const ViewBalance                  = "balance";
export const ViewFoundryOutput            = "foundryOutput";
export const ViewGetAccountNonce          = "getAccountNonce";
export const ViewGetNativeTokenIDRegistry = "getNativeTokenIDRegistry";
export const ViewNftData                  = "nftData";
export const ViewTotalAssets              = "totalAssets";

export const HFuncDeposit                  = new wasmtypes.ScHname(0xbdc9102d);
export const HFuncFoundryCreateNew         = new wasmtypes.ScHname(0x41822f5f);
export const HFuncFoundryDestroy           = new wasmtypes.ScHname(0x85e4c893);
export const HFuncFoundryModifySupply      = new wasmtypes.ScHname(0x76a5868b);
export const HFuncHarvest                  = new wasmtypes.ScHname(0x7b40efbd);
export const HFuncTransferAllowanceTo      = new wasmtypes.ScHname(0x23f4e3a1);
export const HFuncWithdraw                 = new wasmtypes.ScHname(0x9dcc0f41);
export const HViewAccountNFTs              = new wasmtypes.ScHname(0x27422359);
export const HViewAccounts                 = new wasmtypes.ScHname(0x3c4b5e02);
export const HViewBalance                  = new wasmtypes.ScHname(0x84168cb4);
export const HViewFoundryOutput            = new wasmtypes.ScHname(0xd9647be3);
export const HViewGetAccountNonce          = new wasmtypes.ScHname(0x529d7df9);
export const HViewGetNativeTokenIDRegistry = new wasmtypes.ScHname(0x2ad8a59f);
export const HViewNftData                  = new wasmtypes.ScHname(0x83c5c4da);
export const HViewTotalAssets              = new wasmtypes.ScHname(0xfab0f8d2);
