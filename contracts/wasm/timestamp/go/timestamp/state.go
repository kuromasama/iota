// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package timestamp

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type ImmutabletimestampState struct {
	proxy wasmtypes.Proxy
}

type MutabletimestampState struct {
	proxy wasmtypes.Proxy
}

func (s MutabletimestampState) AsImmutable() ImmutabletimestampState {
	return ImmutabletimestampState(s)
}
