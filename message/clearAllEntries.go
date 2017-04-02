package message

import (
	"io"
)

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
func (cae *ClearAllEntries) MarshalMessage(writer io.Writer) error {
	_, err := writer.Write([]byte{cae.Type()})
	if err != nil {
		return err
	}
	_, err = writer.Write(clearAllMagic[:])
	return err
}
