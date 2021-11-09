// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "wasmlib";
import * as sc from "./index";

export class ImmutableInitParams extends wasmlib.ScMapID {
    owner(): wasmlib.ScImmutableAgentID {
		return new wasmlib.ScImmutableAgentID(this.mapID, sc.idxMap[sc.IdxParamOwner]);
	}
}

export class MutableInitParams extends wasmlib.ScMapID {
    owner(): wasmlib.ScMutableAgentID {
		return new wasmlib.ScMutableAgentID(this.mapID, sc.idxMap[sc.IdxParamOwner]);
	}
}

export class ImmutableMemberParams extends wasmlib.ScMapID {
    address(): wasmlib.ScImmutableAddress {
		return new wasmlib.ScImmutableAddress(this.mapID, sc.idxMap[sc.IdxParamAddress]);
	}

    factor(): wasmlib.ScImmutableInt64 {
		return new wasmlib.ScImmutableInt64(this.mapID, sc.idxMap[sc.IdxParamFactor]);
	}
}

export class MutableMemberParams extends wasmlib.ScMapID {
    address(): wasmlib.ScMutableAddress {
		return new wasmlib.ScMutableAddress(this.mapID, sc.idxMap[sc.IdxParamAddress]);
	}

    factor(): wasmlib.ScMutableInt64 {
		return new wasmlib.ScMutableInt64(this.mapID, sc.idxMap[sc.IdxParamFactor]);
	}
}

export class ImmutableSetOwnerParams extends wasmlib.ScMapID {
    owner(): wasmlib.ScImmutableAgentID {
		return new wasmlib.ScImmutableAgentID(this.mapID, sc.idxMap[sc.IdxParamOwner]);
	}
}

export class MutableSetOwnerParams extends wasmlib.ScMapID {
    owner(): wasmlib.ScMutableAgentID {
		return new wasmlib.ScMutableAgentID(this.mapID, sc.idxMap[sc.IdxParamOwner]);
	}
}

export class ImmutableGetFactorParams extends wasmlib.ScMapID {
    address(): wasmlib.ScImmutableAddress {
		return new wasmlib.ScImmutableAddress(this.mapID, sc.idxMap[sc.IdxParamAddress]);
	}
}

export class MutableGetFactorParams extends wasmlib.ScMapID {
    address(): wasmlib.ScMutableAddress {
		return new wasmlib.ScMutableAddress(this.mapID, sc.idxMap[sc.IdxParamAddress]);
	}
}