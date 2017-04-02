package entry

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
			eType: eTypeBooleanArray,
		},
	}
	for _, b := range value {
		boolean := NewBoolean(b)
		ba.value = append(ba.value, boolean)
	}
	return ba
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
