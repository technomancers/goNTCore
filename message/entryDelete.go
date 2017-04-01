package message

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
func (ed *EntryDelete) MarshalMessage() ([]byte, error) {
	var output []byte
	output = append(output, ed.Type())
	output = append(output, ed.entryID[:]...)
	return output, nil
}
