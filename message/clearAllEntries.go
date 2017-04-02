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
	magic [4]byte
	valid bool
}

//NewClearAllEntries creates an instance of clear all entries.
func NewClearAllEntries() *ClearAllEntries {
	return &ClearAllEntries{
		message: message{
			mType: MTypeClearAllEntries,
		},
		magic: clearAllMagic,
		valid: true,
	}
}

//MarshalMessage implements Marshaler for Network Table Messages.
func (cae *ClearAllEntries) MarshalMessage(writer io.Writer) error {
	_, err := writer.Write([]byte{cae.Type()})
	if err != nil {
		return err
	}
	_, err = writer.Write(cae.magic[:])
	return err
}

//UnmarshalMessage implements Unmarshaler for Network Table Messages and assumes the message type bit has already been read.
func (cae *ClearAllEntries) UnmarshalMessage(reader io.Reader) error {
	cae.mType = MTypeClearAllEntries
	buf := make([]byte, 4)

	_, err := io.ReadFull(reader, buf)
	if err != nil {
		return err
	}

	cae.valid = true
	for i := 0; i < len(buf); i++ {
		cae.magic[i] = buf[i]
		if cae.magic[i] != clearAllMagic[i] {
			cae.valid = false
		}
	}
	return nil
}
