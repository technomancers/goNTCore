package entry

import "io"

//Boolean is a Network Table Entry that holds the value of type boolean.
type Boolean struct {
	entry
	value bool
}

//NewBoolean creates a new instance of a Boolean entry
func NewBoolean(value bool) *Boolean {
	return &Boolean{
		entry: entry{
			eType: eTypeBoolean,
		},
		value: value,
	}
}

//MarshalEntry implements Marshaler for Network Table Entry.
func (b *Boolean) MarshalEntry(writer io.Writer) error {
	val := byte(0x00)
	if b.value {
		val = val | 0x01
	}
	_, err := writer.Write([]byte{val})
	return err
}
