// Copyright (c) 2017, Technomancers. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package message

import (
	"errors"
	"io"
)

const (
	MTypeKeepAlive           byte = 0x00
	MTypeClientHello         byte = 0x01
	MTypeProtoUnsupported    byte = 0x02
	MTypeServerHelloComplete byte = 0x03
	MTypeServerHello         byte = 0x04
	MTypeClientHelloComplete byte = 0x05
	MTypeEntryAssign         byte = 0x10
	MTypeEntryUpdate         byte = 0x11
	MTypeEntryFlagUpdate     byte = 0x12
	MTypeEntryDelete         byte = 0x13
	MTypeClearAllEntries     byte = 0x14
	MTypeRPCExecute          byte = 0x20
	MTypeRPCResponse         byte = 0x21
)

type message struct {
	mType byte
}

func (m message) Type() byte {
	return m.mType
}

//Marshaler is the interface implemented by types that can marshal themselves into valid Network Table Message.
type Marshaler interface {
	MarshalMessage(io.Writer) error
}

//Unmarshaler is the interface implemented by types that can unmarshal a Network Table Message of themselves.
type Unmarshaler interface {
	UnmarshalMessage(io.Reader) error
}

//Messager is the interface implemented by types that can communicate on the network.
type Messager interface {
	Type() byte
	Marshaler
	Unmarshaler
}

//Unmarshal takes the type passed in and tries to unmarshal the next bytes from reader based on the type.
//It returns an instance messager.
func Unmarshal(t byte, reader io.Reader) (Messager, error) {
	var msg Messager
	switch t {
	case MTypeKeepAlive:
		msg = new(KeepAlive)
	case MTypeClientHello:
		msg = new(ClientHello)
	case MTypeProtoUnsupported:
		msg = new(ProtoUnsupported)
	case MTypeServerHelloComplete:
		msg = new(ServerHelloComplete)
	case MTypeServerHello:
		msg = new(ServerHello)
	case MTypeClientHelloComplete:
		msg = new(ClientHelloComplete)
	case MTypeEntryAssign:
		msg = new(EntryAssign)
	case MTypeEntryUpdate:
		msg = new(EntryUpdate)
	case MTypeEntryFlagUpdate:
		msg = new(EntryFlagUpdate)
	case MTypeEntryDelete:
		msg = new(EntryDelete)
	case MTypeClearAllEntries:
		msg = new(ClearAllEntries)
	default:
		return nil, errors.New("Unmarshal Message: Could not find appropropriate type")
	}
	err := msg.UnmarshalMessage(reader)
	return msg, err
}
