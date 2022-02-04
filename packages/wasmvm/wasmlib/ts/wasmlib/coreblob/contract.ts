// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "wasmlib";
import * as sc from "./index";

export class StoreBlobCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncStoreBlob);
	params: sc.MutableStoreBlobParams = new sc.MutableStoreBlobParams(wasmlib.ScView.nilProxy);
	results: sc.ImmutableStoreBlobResults = new sc.ImmutableStoreBlobResults(wasmlib.ScView.nilProxy);
}

export class GetBlobFieldCall {
	func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewGetBlobField);
	params: sc.MutableGetBlobFieldParams = new sc.MutableGetBlobFieldParams(wasmlib.ScView.nilProxy);
	results: sc.ImmutableGetBlobFieldResults = new sc.ImmutableGetBlobFieldResults(wasmlib.ScView.nilProxy);
}

export class GetBlobInfoCall {
	func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewGetBlobInfo);
	params: sc.MutableGetBlobInfoParams = new sc.MutableGetBlobInfoParams(wasmlib.ScView.nilProxy);
	results: sc.ImmutableGetBlobInfoResults = new sc.ImmutableGetBlobInfoResults(wasmlib.ScView.nilProxy);
}

export class ListBlobsCall {
	func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewListBlobs);
	results: sc.ImmutableListBlobsResults = new sc.ImmutableListBlobsResults(wasmlib.ScView.nilProxy);
}

export class ScFuncs {
    static storeBlob(_ctx: wasmlib.ScFuncCallContext): StoreBlobCall {
        const f = new StoreBlobCall();
		f.params = new sc.MutableStoreBlobParams(wasmlib.newCallParamsProxy(f.func));
		f.results = new sc.ImmutableStoreBlobResults(wasmlib.newCallResultsProxy(f.func));
        return f;
    }

    static getBlobField(_ctx: wasmlib.ScViewCallContext): GetBlobFieldCall {
        const f = new GetBlobFieldCall();
		f.params = new sc.MutableGetBlobFieldParams(wasmlib.newCallParamsProxy(f.func));
		f.results = new sc.ImmutableGetBlobFieldResults(wasmlib.newCallResultsProxy(f.func));
        return f;
    }

    static getBlobInfo(_ctx: wasmlib.ScViewCallContext): GetBlobInfoCall {
        const f = new GetBlobInfoCall();
		f.params = new sc.MutableGetBlobInfoParams(wasmlib.newCallParamsProxy(f.func));
		f.results = new sc.ImmutableGetBlobInfoResults(wasmlib.newCallResultsProxy(f.func));
        return f;
    }

    static listBlobs(_ctx: wasmlib.ScViewCallContext): ListBlobsCall {
        const f = new ListBlobsCall();
		f.results = new sc.ImmutableListBlobsResults(wasmlib.newCallResultsProxy(f.func));
        return f;
    }
}
