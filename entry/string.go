package entry

import "github.com/technomancers/goNTCore/util"
import "io"

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
func (s *String) MarshalEntry(writer io.Writer) error {
	valueBytes := []byte(s.value)
	err := util.EncodeULeb128(uint32(len(valueBytes)), writer)
	if err != nil {
		return err
	}
	_, err = writer.Write(valueBytes)
	return err
}

//UnmarshalEntry implements Unmarshaler for Network Table Entry.
func (s *String) UnmarshalEntry(reader io.Reader) error {
	s.eType = eTypeString
	lengthString, err := util.DecodeULeb128(reader)
	if err != nil {
		return err
	}
	buf := make([]byte, lengthString)
	_, err = io.ReadFull(reader, buf)
	if err != nil {
		return err
	}
	s.value = string(buf)
	return nil
}
