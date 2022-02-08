// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "wasmlib";
import * as sc from "./index";

export class ForcePayoutCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncForcePayout);
}

export class ForcePayoutContext {
	events: sc.FairRouletteEvents = new sc.FairRouletteEvents();
	state: sc.MutableFairRouletteState = new sc.MutableFairRouletteState(wasmlib.ScState.proxy());
}

export class ForceResetCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncForceReset);
}

export class ForceResetContext {
	events: sc.FairRouletteEvents = new sc.FairRouletteEvents();
	state: sc.MutableFairRouletteState = new sc.MutableFairRouletteState(wasmlib.ScState.proxy());
}

export class PayWinnersCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncPayWinners);
}

export class PayWinnersContext {
	events: sc.FairRouletteEvents = new sc.FairRouletteEvents();
	state: sc.MutableFairRouletteState = new sc.MutableFairRouletteState(wasmlib.ScState.proxy());
}

export class PlaceBetCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncPlaceBet);
	params: sc.MutablePlaceBetParams = new sc.MutablePlaceBetParams(wasmlib.ScView.nilProxy);
}

export class PlaceBetContext {
	events: sc.FairRouletteEvents = new sc.FairRouletteEvents();
	params: sc.ImmutablePlaceBetParams = new sc.ImmutablePlaceBetParams(wasmlib.paramsProxy());
	state: sc.MutableFairRouletteState = new sc.MutableFairRouletteState(wasmlib.ScState.proxy());
}

export class PlayPeriodCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncPlayPeriod);
	params: sc.MutablePlayPeriodParams = new sc.MutablePlayPeriodParams(wasmlib.ScView.nilProxy);
}

export class PlayPeriodContext {
	events: sc.FairRouletteEvents = new sc.FairRouletteEvents();
	params: sc.ImmutablePlayPeriodParams = new sc.ImmutablePlayPeriodParams(wasmlib.paramsProxy());
	state: sc.MutableFairRouletteState = new sc.MutableFairRouletteState(wasmlib.ScState.proxy());
}

export class LastWinningNumberCall {
	func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewLastWinningNumber);
	results: sc.ImmutableLastWinningNumberResults = new sc.ImmutableLastWinningNumberResults(wasmlib.ScView.nilProxy);
}

export class LastWinningNumberContext {
	results: sc.MutableLastWinningNumberResults = new sc.MutableLastWinningNumberResults(wasmlib.ScView.nilProxy);
	state: sc.ImmutableFairRouletteState = new sc.ImmutableFairRouletteState(wasmlib.ScState.proxy());
}

export class RoundNumberCall {
	func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewRoundNumber);
	results: sc.ImmutableRoundNumberResults = new sc.ImmutableRoundNumberResults(wasmlib.ScView.nilProxy);
}

export class RoundNumberContext {
	results: sc.MutableRoundNumberResults = new sc.MutableRoundNumberResults(wasmlib.ScView.nilProxy);
	state: sc.ImmutableFairRouletteState = new sc.ImmutableFairRouletteState(wasmlib.ScState.proxy());
}

export class RoundStartedAtCall {
	func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewRoundStartedAt);
	results: sc.ImmutableRoundStartedAtResults = new sc.ImmutableRoundStartedAtResults(wasmlib.ScView.nilProxy);
}

export class RoundStartedAtContext {
	results: sc.MutableRoundStartedAtResults = new sc.MutableRoundStartedAtResults(wasmlib.ScView.nilProxy);
	state: sc.ImmutableFairRouletteState = new sc.ImmutableFairRouletteState(wasmlib.ScState.proxy());
}

export class RoundStatusCall {
	func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewRoundStatus);
	results: sc.ImmutableRoundStatusResults = new sc.ImmutableRoundStatusResults(wasmlib.ScView.nilProxy);
}

export class RoundStatusContext {
	results: sc.MutableRoundStatusResults = new sc.MutableRoundStatusResults(wasmlib.ScView.nilProxy);
	state: sc.ImmutableFairRouletteState = new sc.ImmutableFairRouletteState(wasmlib.ScState.proxy());
}

export class ScFuncs {
	static forcePayout(_ctx: wasmlib.ScFuncCallContext): ForcePayoutCall {
		return new ForcePayoutCall();
	}

	static forceReset(_ctx: wasmlib.ScFuncCallContext): ForceResetCall {
		return new ForceResetCall();
	}

	static payWinners(_ctx: wasmlib.ScFuncCallContext): PayWinnersCall {
		return new PayWinnersCall();
	}

	static placeBet(_ctx: wasmlib.ScFuncCallContext): PlaceBetCall {
		const f = new PlaceBetCall();
		f.params = new sc.MutablePlaceBetParams(wasmlib.newCallParamsProxy(f.func));
		return f;
	}

	static playPeriod(_ctx: wasmlib.ScFuncCallContext): PlayPeriodCall {
		const f = new PlayPeriodCall();
		f.params = new sc.MutablePlayPeriodParams(wasmlib.newCallParamsProxy(f.func));
		return f;
	}

	static lastWinningNumber(_ctx: wasmlib.ScViewCallContext): LastWinningNumberCall {
		const f = new LastWinningNumberCall();
		f.results = new sc.ImmutableLastWinningNumberResults(wasmlib.newCallResultsProxy(f.func));
		return f;
	}

	static roundNumber(_ctx: wasmlib.ScViewCallContext): RoundNumberCall {
		const f = new RoundNumberCall();
		f.results = new sc.ImmutableRoundNumberResults(wasmlib.newCallResultsProxy(f.func));
		return f;
	}

	static roundStartedAt(_ctx: wasmlib.ScViewCallContext): RoundStartedAtCall {
		const f = new RoundStartedAtCall();
		f.results = new sc.ImmutableRoundStartedAtResults(wasmlib.newCallResultsProxy(f.func));
		return f;
	}

	static roundStatus(_ctx: wasmlib.ScViewCallContext): RoundStatusCall {
		const f = new RoundStatusCall();
		f.results = new sc.ImmutableRoundStatusResults(wasmlib.newCallResultsProxy(f.func));
		return f;
	}
}
