package entry

import "github.com/technomancers/goNTCore/util"

//String is a Network Table Entry that holds the value of type string.
type String struct {
	entry
	value string
}

//NewString creates a new Network Table Entry of type string.
func NewString(value string) *String {
	return &String{
		entry: entry{
			eType: eTypeString,
		},
		value: value,
	}
}

//MarshalEntry implements Marshaler for Network Table Entry.
func (s *String) MarshalEntry() ([]byte, error) {
	var output []byte
	valueBytes := []byte(s.value)
	valueLen := util.EncodeULeb128(uint32(len(valueBytes)))
	output = append(output, valueLen...)
	output = append(output, valueBytes...)
	return output, nil
}
