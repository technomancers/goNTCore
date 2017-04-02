package entry

import "io"

//DoubleArray is a Network Table Entry that holds the value of type Array of Doubles.
type DoubleArray struct {
	entry
	value []*Double
}

//NewDoubleArray creates an instance of DoubleArray.
func NewDoubleArray(value []float64) *DoubleArray {
	da := &DoubleArray{
		entry: entry{
			eType: eTypeDoubleArray,
		},
	}
	for _, d := range value {
		double := NewDouble(d)
		da.value = append(da.value, double)
	}
	return da
}

//MarshalEntry implements Marshaler for Network Table Entry.
func (da *DoubleArray) MarshalEntry(writer io.Writer) error {
	lenArray := byte(len(da.value))
	_, err := writer.Write([]byte{lenArray})
	if err != nil {
		return err
	}
	for _, d := range da.value {
		err = d.MarshalEntry(writer)
		if err != nil {
			return err
		}
	}
	return nil
}
