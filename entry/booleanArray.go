package entry

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
func (ba *BooleanArray) MarshalEntry() ([]byte, error) {
	lenArray := byte(len(ba.value))
	var output []byte
	output = append(output, lenArray)
	for _, b := range ba.value {
		val, err := b.MarshalEntry()
		if err != nil {
			return nil, err
		}
		output = append(output, val...)
	}
	return output, nil
}
