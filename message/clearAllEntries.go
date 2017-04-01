package message

var (
	clearAllMagic = [4]byte{0xd0, 0x6c, 0xb2, 0x7a}
)

//ClearAllEntries clears all entries from the network.
type ClearAllEntries struct {
	message
}

//NewClearAllEntries creates an instance of clear all entries.
func NewClearAllEntries() *ClearAllEntries {
	return &ClearAllEntries{
		message{
			mType: mTypeClearAllEntries,
		},
	}
}

//MarshalMessage implements Marshaler for Network Table Messages.
func (cae *ClearAllEntries) MarshalMessage() ([]byte, error) {
	var output []byte
	output = append(output, cae.Type())
	output = append(output, clearAllMagic[:]...)
	return output, nil
}
