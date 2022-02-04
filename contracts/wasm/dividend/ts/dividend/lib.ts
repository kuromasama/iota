// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "wasmlib";
import * as sc from "./index";

export function on_call(index: i32): void {
    return wasmlib.onCall(index);
}

export function on_load(): void {
    let exports = new wasmlib.ScExports();
    exports.addFunc(sc.FuncDivide,    funcDivideThunk);
    exports.addFunc(sc.FuncInit,      funcInitThunk);
    exports.addFunc(sc.FuncMember,    funcMemberThunk);
    exports.addFunc(sc.FuncSetOwner,  funcSetOwnerThunk);
    exports.addView(sc.ViewGetFactor, viewGetFactorThunk);
    exports.addView(sc.ViewGetOwner,  viewGetOwnerThunk);
}

function funcDivideThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("dividend.funcDivide");
	let f = new sc.DivideContext();
	sc.funcDivide(ctx, f);
	ctx.log("dividend.funcDivide ok");
}

function funcInitThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("dividend.funcInit");
	let f = new sc.InitContext();
	sc.funcInit(ctx, f);
	ctx.log("dividend.funcInit ok");
}

function funcMemberThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("dividend.funcMember");
	let f = new sc.MemberContext();

	// only defined owner of contract can add members
	const access = f.state.owner();
	ctx.require(access.exists(), "access not set: owner");
	ctx.require(ctx.caller().equals(access.value()), "no permission");

	ctx.require(f.params.address().exists(), "missing mandatory address");
	ctx.require(f.params.factor().exists(), "missing mandatory factor");
	sc.funcMember(ctx, f);
	ctx.log("dividend.funcMember ok");
}

function funcSetOwnerThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("dividend.funcSetOwner");
	let f = new sc.SetOwnerContext();

	// only defined owner of contract can change owner
	const access = f.state.owner();
	ctx.require(access.exists(), "access not set: owner");
	ctx.require(ctx.caller().equals(access.value()), "no permission");

	ctx.require(f.params.owner().exists(), "missing mandatory owner");
	sc.funcSetOwner(ctx, f);
	ctx.log("dividend.funcSetOwner ok");
}

function viewGetFactorThunk(ctx: wasmlib.ScViewContext): void {
	ctx.log("dividend.viewGetFactor");
	let f = new sc.GetFactorContext();
    const results = new wasmlib.ScDict([]);
	f.results = new sc.MutableGetFactorResults(results.asProxy());
	ctx.require(f.params.address().exists(), "missing mandatory address");
	sc.viewGetFactor(ctx, f);
	ctx.results(results);
	ctx.log("dividend.viewGetFactor ok");
}

function viewGetOwnerThunk(ctx: wasmlib.ScViewContext): void {
	ctx.log("dividend.viewGetOwner");
	let f = new sc.GetOwnerContext();
    const results = new wasmlib.ScDict([]);
	f.results = new sc.MutableGetOwnerResults(results.asProxy());
	sc.viewGetOwner(ctx, f);
	ctx.results(results);
	ctx.log("dividend.viewGetOwner ok");
}
