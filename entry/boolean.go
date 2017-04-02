package entry

import "io"

const (
	booleanTrue  byte = 0x01
	booleanFalse byte = 0x00
)

//Boolean is a Network Table Entry that holds the value of type boolean.
type Boolean struct {
	entry
	value bool
}

//NewBoolean creates a new instance of a Boolean entry
func NewBoolean(value bool) *Boolean {
	return &Boolean{
		entry: entry{
			eType: ETypeBoolean,
		},
		value: value,
	}
}

//MarshalEntry implements Marshaler for Network Table Entry.
func (b *Boolean) MarshalEntry(writer io.Writer) error {
	val := booleanFalse
	if b.value {
		val = val | booleanTrue
	}
	_, err := writer.Write([]byte{val})
	return err
}

//UnmarshalEntry implements Unmarshaler for Network Table Entry.
func (b *Boolean) UnmarshalEntry(reader io.Reader) error {
	b.eType = ETypeBoolean
	buf := make([]byte, 1)
	_, err := io.ReadFull(reader, buf)
	if err != nil {
		return err
	}
	b.value = buf[0] == booleanTrue
	return nil
}
