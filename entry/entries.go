package entry

import (
	"errors"
	"io"
)

const (
	ETypeBoolean      byte = 0x00
	ETypeDouble       byte = 0x01
	ETypeString       byte = 0x02
	ETypeRawData      byte = 0x03
	ETypeBooleanArray byte = 0x10
	ETypeDoubleArray  byte = 0x11
	ETypeStringArray  byte = 0x12
	ETypeRPCDef       byte = 0x20
)

type entry struct {
	eType byte
}

func (e entry) Type() byte {
	return e.eType
}

//Marshaler is the interface implemented by types that can marshal themselves into valid Network Table Entry Value.
type Marshaler interface {
	MarshalEntry(io.Writer) error
}

//Unmarshaler is the interface implemented by types that can unmarshal a Network Table Entry Value of themselves.
type Unmarshaler interface {
	UnmarshalEntry(io.Reader) error
}

//Entrier is the interface implemented by types that can be an Entry in the Network Tables.
type Entrier interface {
	Type() byte
	Marshaler
	Unmarshaler
}

//Unmarshal takes the type passed in and tries to unmarshal the next bytes from reader based on the type.
//It returns an instance entry.
func Unmarshal(t byte, reader io.Reader) (Entrier, error) {
	var ent Entrier
	switch t {
	case ETypeBoolean:
		ent = new(Boolean)
	case ETypeDouble:
		ent = new(Double)
	case ETypeString:
		ent = new(String)
	case ETypeRawData:
		ent = new(RawData)
	case ETypeBooleanArray:
		ent = new(BooleanArray)
	case ETypeDoubleArray:
		ent = new(DoubleArray)
	case ETypeStringArray:
		ent = new(StringArray)
	default:
		return nil, errors.New("Unmarshal Entry: Could not find appropropriate type")
	}
	err := ent.UnmarshalEntry(reader)
	if err != nil {
		return nil, err
	}
	return ent, nil
}
