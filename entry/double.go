package entry

import (
	"encoding/binary"
	"io"
	"math"
)

//Double is a Network Table Entry that holds the value of type double(float64).
type Double struct {
	entry
	value float64
}

//NewDouble creates a new instance of double.
func NewDouble(value float64) *Double {
	return &Double{
		entry: entry{
			eType: eTypeDouble,
		},
		value: value,
	}
}

//MarshalEntry implements Marshaler for Network Table Entry.
func (d *Double) MarshalEntry(writer io.Writer) error {
	val := math.Float64bits(d.value)
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, val)
	_, err := writer.Write(buf)
	return err
}
