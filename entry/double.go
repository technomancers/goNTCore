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
			eType: ETypeDouble,
		},
		value: value,
	}
}

//GetValue gets the value of the string.
func (d *Double) GetValue() float64 {
	return d.value
}

//MarshalEntry implements Marshaler for Network Table Entry.
func (d *Double) MarshalEntry(writer io.Writer) error {
	val := math.Float64bits(d.value)
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, val)
	_, err := writer.Write(buf)
	return err
}

//UnmarshalEntry implements Unmarshaler for Network Table Entry.
func (d *Double) UnmarshalEntry(reader io.Reader) error {
	d.eType = ETypeDouble
	buf := make([]byte, 8)
	_, err := io.ReadFull(reader, buf)
	if err != nil {
		return err
	}
	double := binary.BigEndian.Uint64(buf)
	d.value = math.Float64frombits(double)
	return nil
}
