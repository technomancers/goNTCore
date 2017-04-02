package entryType

import "io"

//BooleanArray is a Network Table Entry that holds the value of type Array of Booleans.
type BooleanArray struct {
	entry
	value []*Boolean
}

//NewBooleanArray creates an instance of BooleanArray.
func NewBooleanArray(value []bool) *BooleanArray {
	ba := &BooleanArray{
		entry: entry{
			eType: ETypeBooleanArray,
		},
	}
	for _, b := range value {
		boolean := NewBoolean(b)
		ba.value = append(ba.value, boolean)
	}
	return ba
}

//GetValue gets the value of the string.
func (ba *BooleanArray) GetValue() []bool {
	var out []bool
	for _, b := range ba.value {
		out = append(out, b.GetValue())
	}
	return out
}

//MarshalEntry implements Marshaler for Network Table Entry.
func (ba *BooleanArray) MarshalEntry(writer io.Writer) error {
	lenArray := byte(len(ba.value))
	_, err := writer.Write([]byte{lenArray})
	if err != nil {
		return err
	}
	for _, b := range ba.value {
		err = b.MarshalEntry(writer)
		if err != nil {
			return err
		}
	}
	return nil
}

//UnmarshalEntry implements Unmarshaler for Network Table Entry.
func (ba *BooleanArray) UnmarshalEntry(reader io.Reader) error {
	ba.eType = ETypeBooleanArray
	lenBuf := make([]byte, 1)
	_, err := io.ReadFull(reader, lenBuf)
	if err != nil {
		return err
	}
	numEle := int(lenBuf[0])
	for i := 0; i < numEle; i++ {
		boolean := new(Boolean)
		err = boolean.UnmarshalEntry(reader)
		if err != nil {
			return err
		}
		ba.value = append(ba.value, boolean)
	}
	return nil
}
