// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the schema definition file instead

#![allow(dead_code)]
#![allow(unused_imports)]

use crate::*;
use crate::coreaccounts::*;

#[derive(Clone)]
pub struct ImmutableFoundryCreateNewParams {
    pub(crate) proxy: Proxy,
}

impl ImmutableFoundryCreateNewParams {
    pub fn new() -> ImmutableFoundryCreateNewParams {
        ImmutableFoundryCreateNewParams {
            proxy: params_proxy(),
        }
    }

    pub fn token_scheme(&self) -> ScImmutableBytes {
        ScImmutableBytes::new(self.proxy.root(PARAM_TOKEN_SCHEME))
    }
}

#[derive(Clone)]
pub struct MutableFoundryCreateNewParams {
    pub(crate) proxy: Proxy,
}

impl MutableFoundryCreateNewParams {
    pub fn token_scheme(&self) -> ScMutableBytes {
        ScMutableBytes::new(self.proxy.root(PARAM_TOKEN_SCHEME))
    }
}

#[derive(Clone)]
pub struct ImmutableFoundryDestroyParams {
    pub(crate) proxy: Proxy,
}

impl ImmutableFoundryDestroyParams {
    pub fn new() -> ImmutableFoundryDestroyParams {
        ImmutableFoundryDestroyParams {
            proxy: params_proxy(),
        }
    }

    pub fn foundry_sn(&self) -> ScImmutableUint32 {
        ScImmutableUint32::new(self.proxy.root(PARAM_FOUNDRY_SN))
    }
}

#[derive(Clone)]
pub struct MutableFoundryDestroyParams {
    pub(crate) proxy: Proxy,
}

impl MutableFoundryDestroyParams {
    pub fn foundry_sn(&self) -> ScMutableUint32 {
        ScMutableUint32::new(self.proxy.root(PARAM_FOUNDRY_SN))
    }
}

#[derive(Clone)]
pub struct ImmutableFoundryModifySupplyParams {
    pub(crate) proxy: Proxy,
}

impl ImmutableFoundryModifySupplyParams {
    pub fn new() -> ImmutableFoundryModifySupplyParams {
        ImmutableFoundryModifySupplyParams {
            proxy: params_proxy(),
        }
    }

    pub fn destroy_tokens(&self) -> ScImmutableBool {
        ScImmutableBool::new(self.proxy.root(PARAM_DESTROY_TOKENS))
    }

    pub fn foundry_sn(&self) -> ScImmutableUint32 {
        ScImmutableUint32::new(self.proxy.root(PARAM_FOUNDRY_SN))
    }

    pub fn supply_delta_abs(&self) -> ScImmutableBigInt {
        ScImmutableBigInt::new(self.proxy.root(PARAM_SUPPLY_DELTA_ABS))
    }
}

#[derive(Clone)]
pub struct MutableFoundryModifySupplyParams {
    pub(crate) proxy: Proxy,
}

impl MutableFoundryModifySupplyParams {
    pub fn destroy_tokens(&self) -> ScMutableBool {
        ScMutableBool::new(self.proxy.root(PARAM_DESTROY_TOKENS))
    }

    pub fn foundry_sn(&self) -> ScMutableUint32 {
        ScMutableUint32::new(self.proxy.root(PARAM_FOUNDRY_SN))
    }

    pub fn supply_delta_abs(&self) -> ScMutableBigInt {
        ScMutableBigInt::new(self.proxy.root(PARAM_SUPPLY_DELTA_ABS))
    }
}

#[derive(Clone)]
pub struct ImmutableHarvestParams {
    pub(crate) proxy: Proxy,
}

impl ImmutableHarvestParams {
    pub fn new() -> ImmutableHarvestParams {
        ImmutableHarvestParams {
            proxy: params_proxy(),
        }
    }

    pub fn force_minimum_base_tokens(&self) -> ScImmutableUint64 {
        ScImmutableUint64::new(self.proxy.root(PARAM_FORCE_MINIMUM_BASE_TOKENS))
    }
}

#[derive(Clone)]
pub struct MutableHarvestParams {
    pub(crate) proxy: Proxy,
}

impl MutableHarvestParams {
    pub fn force_minimum_base_tokens(&self) -> ScMutableUint64 {
        ScMutableUint64::new(self.proxy.root(PARAM_FORCE_MINIMUM_BASE_TOKENS))
    }
}

#[derive(Clone)]
pub struct ImmutableTransferAllowanceToParams {
    pub(crate) proxy: Proxy,
}

impl ImmutableTransferAllowanceToParams {
    pub fn new() -> ImmutableTransferAllowanceToParams {
        ImmutableTransferAllowanceToParams {
            proxy: params_proxy(),
        }
    }

    pub fn agent_id(&self) -> ScImmutableAgentID {
        ScImmutableAgentID::new(self.proxy.root(PARAM_AGENT_ID))
    }
}

#[derive(Clone)]
pub struct MutableTransferAllowanceToParams {
    pub(crate) proxy: Proxy,
}

impl MutableTransferAllowanceToParams {
    pub fn agent_id(&self) -> ScMutableAgentID {
        ScMutableAgentID::new(self.proxy.root(PARAM_AGENT_ID))
    }
}

#[derive(Clone)]
pub struct ImmutableAccountNFTsParams {
    pub(crate) proxy: Proxy,
}

impl ImmutableAccountNFTsParams {
    pub fn new() -> ImmutableAccountNFTsParams {
        ImmutableAccountNFTsParams {
            proxy: params_proxy(),
        }
    }

    pub fn agent_id(&self) -> ScImmutableAgentID {
        ScImmutableAgentID::new(self.proxy.root(PARAM_AGENT_ID))
    }
}

#[derive(Clone)]
pub struct MutableAccountNFTsParams {
    pub(crate) proxy: Proxy,
}

impl MutableAccountNFTsParams {
    pub fn agent_id(&self) -> ScMutableAgentID {
        ScMutableAgentID::new(self.proxy.root(PARAM_AGENT_ID))
    }
}

#[derive(Clone)]
pub struct ImmutableBalanceParams {
    pub(crate) proxy: Proxy,
}

impl ImmutableBalanceParams {
    pub fn new() -> ImmutableBalanceParams {
        ImmutableBalanceParams {
            proxy: params_proxy(),
        }
    }

    pub fn agent_id(&self) -> ScImmutableAgentID {
        ScImmutableAgentID::new(self.proxy.root(PARAM_AGENT_ID))
    }
}

#[derive(Clone)]
pub struct MutableBalanceParams {
    pub(crate) proxy: Proxy,
}

impl MutableBalanceParams {
    pub fn agent_id(&self) -> ScMutableAgentID {
        ScMutableAgentID::new(self.proxy.root(PARAM_AGENT_ID))
    }
}

#[derive(Clone)]
pub struct ImmutableBalanceBaseTokenParams {
    pub(crate) proxy: Proxy,
}

impl ImmutableBalanceBaseTokenParams {
    pub fn new() -> ImmutableBalanceBaseTokenParams {
        ImmutableBalanceBaseTokenParams {
            proxy: params_proxy(),
        }
    }

    pub fn agent_id(&self) -> ScImmutableAgentID {
        ScImmutableAgentID::new(self.proxy.root(PARAM_AGENT_ID))
    }
}

#[derive(Clone)]
pub struct MutableBalanceBaseTokenParams {
    pub(crate) proxy: Proxy,
}

impl MutableBalanceBaseTokenParams {
    pub fn agent_id(&self) -> ScMutableAgentID {
        ScMutableAgentID::new(self.proxy.root(PARAM_AGENT_ID))
    }
}

#[derive(Clone)]
pub struct ImmutableBalanceNativeTokenParams {
    pub(crate) proxy: Proxy,
}

impl ImmutableBalanceNativeTokenParams {
    pub fn new() -> ImmutableBalanceNativeTokenParams {
        ImmutableBalanceNativeTokenParams {
            proxy: params_proxy(),
        }
    }

    pub fn agent_id(&self) -> ScImmutableAgentID {
        ScImmutableAgentID::new(self.proxy.root(PARAM_AGENT_ID))
    }

    pub fn token_id(&self) -> ScImmutableTokenID {
        ScImmutableTokenID::new(self.proxy.root(PARAM_TOKEN_ID))
    }
}

#[derive(Clone)]
pub struct MutableBalanceNativeTokenParams {
    pub(crate) proxy: Proxy,
}

impl MutableBalanceNativeTokenParams {
    pub fn agent_id(&self) -> ScMutableAgentID {
        ScMutableAgentID::new(self.proxy.root(PARAM_AGENT_ID))
    }

    pub fn token_id(&self) -> ScMutableTokenID {
        ScMutableTokenID::new(self.proxy.root(PARAM_TOKEN_ID))
    }
}

#[derive(Clone)]
pub struct ImmutableFoundryOutputParams {
    pub(crate) proxy: Proxy,
}

impl ImmutableFoundryOutputParams {
    pub fn new() -> ImmutableFoundryOutputParams {
        ImmutableFoundryOutputParams {
            proxy: params_proxy(),
        }
    }

    pub fn foundry_sn(&self) -> ScImmutableUint32 {
        ScImmutableUint32::new(self.proxy.root(PARAM_FOUNDRY_SN))
    }
}

#[derive(Clone)]
pub struct MutableFoundryOutputParams {
    pub(crate) proxy: Proxy,
}

impl MutableFoundryOutputParams {
    pub fn foundry_sn(&self) -> ScMutableUint32 {
        ScMutableUint32::new(self.proxy.root(PARAM_FOUNDRY_SN))
    }
}

#[derive(Clone)]
pub struct ImmutableGetAccountNonceParams {
    pub(crate) proxy: Proxy,
}

impl ImmutableGetAccountNonceParams {
    pub fn new() -> ImmutableGetAccountNonceParams {
        ImmutableGetAccountNonceParams {
            proxy: params_proxy(),
        }
    }

    pub fn agent_id(&self) -> ScImmutableAgentID {
        ScImmutableAgentID::new(self.proxy.root(PARAM_AGENT_ID))
    }
}

#[derive(Clone)]
pub struct MutableGetAccountNonceParams {
    pub(crate) proxy: Proxy,
}

impl MutableGetAccountNonceParams {
    pub fn agent_id(&self) -> ScMutableAgentID {
        ScMutableAgentID::new(self.proxy.root(PARAM_AGENT_ID))
    }
}

#[derive(Clone)]
pub struct ImmutableNftDataParams {
    pub(crate) proxy: Proxy,
}

impl ImmutableNftDataParams {
    pub fn new() -> ImmutableNftDataParams {
        ImmutableNftDataParams {
            proxy: params_proxy(),
        }
    }

    pub fn nft_id(&self) -> ScImmutableNftID {
        ScImmutableNftID::new(self.proxy.root(PARAM_NFT_ID))
    }
}

#[derive(Clone)]
pub struct MutableNftDataParams {
    pub(crate) proxy: Proxy,
}

impl MutableNftDataParams {
    pub fn nft_id(&self) -> ScMutableNftID {
        ScMutableNftID::new(self.proxy.root(PARAM_NFT_ID))
    }
}
