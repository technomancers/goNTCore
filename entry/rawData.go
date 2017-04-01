package entry

import "github.com/technomancers/goNTCore/util"

//RawData is any type that is not covered in the protocole.
//It is recommended to add a byte or two to the beginning to describe what type it is.
type RawData struct {
	entry
	data []byte
}

//NewRawData creates and instance of RawData.
func NewRawData(data []byte) *RawData {
	return &RawData{
		entry: entry{
			eType: eTypeRawData,
		},
		data: data,
	}
}

//MarshalEntry implements Marshaler for Network Table Entry.
func (rd *RawData) MarshalEntry() ([]byte, error) {
	valueLen := util.EncodeULeb128(uint32(len(rd.data)))
	var output []byte
	output = append(output, valueLen...)
	output = append(output, rd.data...)
	return output, nil
}
