package entry

import "io"

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
func (sa *StringArray) MarshalEntry(writer io.Writer) error {
	lenArray := byte(len(sa.value))
	_, err := writer.Write([]byte{lenArray})
	if err != nil {
		return err
	}
	for _, s := range sa.value {
		err = s.MarshalEntry(writer)
		if err != nil {
			return err
		}
	}
	return nil
}
