package entry

const (
	eTypeBoolean      byte = 0x00
	eTypeDouble       byte = 0x01
	eTypeString       byte = 0x02
	eTypeRawData      byte = 0x03
	eTypeBooleanArray byte = 0x10
	eTypeDoubleArray  byte = 0x11
	eTypeStringArray  byte = 0x12
	eTypeRPCDef       byte = 0x20
)

type entry struct {
	eType byte
}

func (e entry) Type() byte {
	return e.eType
}

//Marshaler is the interface implemented by types that can marshal themselves into valid Network Table Entry Value.
type Marshaler interface {
	MarshalEntry() ([]byte, error)
}

//Unmarshaler is the interface implemented by types that can unmarshal a Network Table Entry Value of themselves.
type Unmarshaler interface {
	UnmarshalEntry([]byte) error
}

//Entrier is the interface implemented by types that can be an Entry in the Network Tables.
type Entrier interface {
	Type() byte
	Marshaler
}
