package message

import "io"

//EntryDelete is used to delete an entry from the network.
type EntryDelete struct {
	message
	entryID [2]byte
}

//NewEntryDelete creates a new instance of EntryDelete.
func NewEntryDelete(id [2]byte) *EntryDelete {
	return &EntryDelete{
		message: message{
			mType: mTypeEntryDelete,
		},
		entryID: id,
	}
}

//MarshalMessage implements Marshaler for Network Table Messages.
func (ed *EntryDelete) MarshalMessage(writer io.Writer) error {
	_, err := writer.Write([]byte{ed.Type()})
	if err != nil {
		return err
	}
	_, err = writer.Write(ed.entryID[:])
	return err
}

//UnmarshalMessage implements Unmarshaler for Network Table Messages and assumes the message type byte has already been read.
func (ed *EntryDelete) UnmarshalMessage(reader io.Reader) error {
	ed.mType = mTypeEntryDelete
	idBuf := make([]byte, 2)
	_, err := io.ReadFull(reader, idBuf)
	if err != nil {
		return err
	}
	for i := 0; i < len(idBuf); i++ {
		ed.entryID[i] = idBuf[i]
	}
	return nil
}
