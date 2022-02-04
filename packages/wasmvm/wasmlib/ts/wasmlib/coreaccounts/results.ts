// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmtypes from "wasmlib/wasmtypes";
import * as sc from "./index";

export class MapAgentIDToImmutableBytes extends wasmtypes.ScProxy {

    getBytes(key: wasmtypes.ScAgentID): wasmtypes.ScImmutableBytes {
        return new wasmtypes.ScImmutableBytes(this.proxy.key(wasmtypes.agentIDToBytes(key)));
    }
}

export class ImmutableAccountsResults extends wasmtypes.ScProxy {
    agents(): sc.MapAgentIDToImmutableBytes {
		return new sc.MapAgentIDToImmutableBytes(this.proxy);
	}
}

export class MapAgentIDToMutableBytes extends wasmtypes.ScProxy {

    clear(): void {
        this.proxy.clearMap();
    }

    getBytes(key: wasmtypes.ScAgentID): wasmtypes.ScMutableBytes {
        return new wasmtypes.ScMutableBytes(this.proxy.key(wasmtypes.agentIDToBytes(key)));
    }
}

export class MutableAccountsResults extends wasmtypes.ScProxy {
    agents(): sc.MapAgentIDToMutableBytes {
		return new sc.MapAgentIDToMutableBytes(this.proxy);
	}
}

export class MapColorToImmutableInt64 extends wasmtypes.ScProxy {

    getInt64(key: wasmtypes.ScColor): wasmtypes.ScImmutableInt64 {
        return new wasmtypes.ScImmutableInt64(this.proxy.key(wasmtypes.colorToBytes(key)));
    }
}

export class ImmutableBalanceResults extends wasmtypes.ScProxy {
    balances(): sc.MapColorToImmutableInt64 {
		return new sc.MapColorToImmutableInt64(this.proxy);
	}
}

export class MapColorToMutableInt64 extends wasmtypes.ScProxy {

    clear(): void {
        this.proxy.clearMap();
    }

    getInt64(key: wasmtypes.ScColor): wasmtypes.ScMutableInt64 {
        return new wasmtypes.ScMutableInt64(this.proxy.key(wasmtypes.colorToBytes(key)));
    }
}

export class MutableBalanceResults extends wasmtypes.ScProxy {
    balances(): sc.MapColorToMutableInt64 {
		return new sc.MapColorToMutableInt64(this.proxy);
	}
}

export class ImmutableGetAccountNonceResults extends wasmtypes.ScProxy {
    accountNonce(): wasmtypes.ScImmutableInt64 {
		return new wasmtypes.ScImmutableInt64(this.proxy.root(sc.ResultAccountNonce));
	}
}

export class MutableGetAccountNonceResults extends wasmtypes.ScProxy {
    accountNonce(): wasmtypes.ScMutableInt64 {
		return new wasmtypes.ScMutableInt64(this.proxy.root(sc.ResultAccountNonce));
	}
}

export class ImmutableTotalAssetsResults extends wasmtypes.ScProxy {
    balances(): sc.MapColorToImmutableInt64 {
		return new sc.MapColorToImmutableInt64(this.proxy);
	}
}

export class MutableTotalAssetsResults extends wasmtypes.ScProxy {
    balances(): sc.MapColorToMutableInt64 {
		return new sc.MapColorToMutableInt64(this.proxy);
	}
}
