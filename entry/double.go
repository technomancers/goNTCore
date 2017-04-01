package entry

import (
	"encoding/binary"
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
func (d *Double) MarshalEntry() ([]byte, error) {
	val := math.Float64bits(d.value)
	var output []byte
	binary.BigEndian.PutUint64(output, val)
	return output, nil
}
