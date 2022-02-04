// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmtypes from "wasmlib/wasmtypes";
import * as sc from "./index";

export class ArrayOfImmutableAddress extends wasmtypes.ScProxy {

    length(): u32 {
        return this.proxy.length();
    }

    getAddress(index: u32): wasmtypes.ScImmutableAddress {
        return new wasmtypes.ScImmutableAddress(this.proxy.index(index));
    }
}

export class MapAddressToImmutableUint64 extends wasmtypes.ScProxy {

    getUint64(key: wasmtypes.ScAddress): wasmtypes.ScImmutableUint64 {
        return new wasmtypes.ScImmutableUint64(this.proxy.key(wasmtypes.addressToBytes(key)));
    }
}

export class ImmutableDividendState extends wasmtypes.ScProxy {
    memberList(): sc.ArrayOfImmutableAddress {
		return new sc.ArrayOfImmutableAddress(this.proxy.root(sc.StateMemberList));
	}

    members(): sc.MapAddressToImmutableUint64 {
		return new sc.MapAddressToImmutableUint64(this.proxy.root(sc.StateMembers));
	}

    owner(): wasmtypes.ScImmutableAgentID {
		return new wasmtypes.ScImmutableAgentID(this.proxy.root(sc.StateOwner));
	}

    totalFactor(): wasmtypes.ScImmutableUint64 {
		return new wasmtypes.ScImmutableUint64(this.proxy.root(sc.StateTotalFactor));
	}
}

export class ArrayOfMutableAddress extends wasmtypes.ScProxy {

	appendAddress(): wasmtypes.ScMutableAddress {
		return new wasmtypes.ScMutableAddress(this.proxy.append());
	}

    clear(): void {
        this.proxy.clearArray();
    }

    length(): u32 {
        return this.proxy.length();
    }

    getAddress(index: u32): wasmtypes.ScMutableAddress {
        return new wasmtypes.ScMutableAddress(this.proxy.index(index));
    }
}

export class MapAddressToMutableUint64 extends wasmtypes.ScProxy {

    clear(): void {
        this.proxy.clearMap();
    }

    getUint64(key: wasmtypes.ScAddress): wasmtypes.ScMutableUint64 {
        return new wasmtypes.ScMutableUint64(this.proxy.key(wasmtypes.addressToBytes(key)));
    }
}

export class MutableDividendState extends wasmtypes.ScProxy {
    asImmutable(): sc.ImmutableDividendState {
		return new sc.ImmutableDividendState(this.proxy);
	}

    memberList(): sc.ArrayOfMutableAddress {
		return new sc.ArrayOfMutableAddress(this.proxy.root(sc.StateMemberList));
	}

    members(): sc.MapAddressToMutableUint64 {
		return new sc.MapAddressToMutableUint64(this.proxy.root(sc.StateMembers));
	}

    owner(): wasmtypes.ScMutableAgentID {
		return new wasmtypes.ScMutableAgentID(this.proxy.root(sc.StateOwner));
	}

    totalFactor(): wasmtypes.ScMutableUint64 {
		return new wasmtypes.ScMutableUint64(this.proxy.root(sc.StateTotalFactor));
	}
}
