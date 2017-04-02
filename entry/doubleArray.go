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
			eType: ETypeDoubleArray,
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

//UnmarshalEntry implements Unmarshaler for Network Table Entry.
func (da *DoubleArray) UnmarshalEntry(reader io.Reader) error {
	da.eType = ETypeDoubleArray
	lenBuf := make([]byte, 1)
	_, err := io.ReadFull(reader, lenBuf)
	if err != nil {
		return err
	}
	numEle := int(lenBuf[0])
	for i := 0; i < numEle; i++ {
		double := new(Double)
		err = double.UnmarshalEntry(reader)
		if err != nil {
			return err
		}
		da.value = append(da.value, double)
	}
	return nil
}
