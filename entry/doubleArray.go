package entry

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
func (da *DoubleArray) MarshalEntry() ([]byte, error) {
	lenArray := byte(len(da.value))
	var output []byte
	output = append(output, lenArray)
	for _, d := range da.value {
		val, err := d.MarshalEntry()
		if err != nil {
			return nil, err
		}
		output = append(output, val...)
	}
	return output, nil
}
