package entry

//StringArray is a Network Table Entry that holds the value of type Array of Strings.
type StringArray struct {
	entry
	value []*String
}

//NewStringArray creates an instance of StringArray.
func NewStringArray(value []string) *StringArray {
	sa := &StringArray{
		entry: entry{
			eType: eTypeStringArray,
		},
	}
	for _, s := range value {
		st := NewString(s)
		sa.value = append(sa.value, st)
	}
	return sa
}

//MarshalEntry implements Marshaler for Network Table Entry.
func (sa *StringArray) MarshalEntry() ([]byte, error) {
	lenArray := byte(len(sa.value))
	var output []byte
	output = append(output, lenArray)
	for _, s := range sa.value {
		val, err := s.MarshalEntry()
		if err != nil {
			return nil, err
		}
		output = append(output, val...)
	}
	return output, nil
}
