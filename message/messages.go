package message

const (
	mTypeKeepAlive           byte = 0x00
	mTypeClientHello         byte = 0x01
	mTypeProtoUnsupported    byte = 0x02
	mTypeServerHelloComplete byte = 0x03
	mTypeServerHello         byte = 0x04
	mTypeClientHelloComplete byte = 0x05
	mTypeEntryAssign         byte = 0x10
	mTypeEntryUpdate         byte = 0x11
	mTypeEntryFlagUpdate     byte = 0x12
	mTypeEntryDelete         byte = 0x13
	mTypeClearAllEntries     byte = 0x14
	mTypeRPCExecute          byte = 0x20
	mTypeRPCResponse         byte = 0x21
)

type message struct {
	mType byte
}

func (m message) Type() byte {
	return m.mType
}

//Marshaler is the interface implemented by types that can marshal themselves into valid Network Table Message.
type Marshaler interface {
	MarshalMessage() ([]byte, error)
}

//Unmarshaler is the interface implemented by types that can unmarshal a Network Table Message of themselves.
type Unmarshaler interface {
	UnmarshalMessage([]byte) error
}
