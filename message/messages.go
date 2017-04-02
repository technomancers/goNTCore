package message

import (
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
