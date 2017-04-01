package message

import (
	"github.com/technomancers/goNTCore/entry"
)

//EntryUpdate is used to tell the network that an entry has been updated.
type EntryUpdate struct {
	message
	entryID [2]byte
	entrySN [2]byte
	entrier entry.Entrier
}

//NewEntryUpdate creates a new instance on EntryUpdate.
func NewEntryUpdate(id, sn [2]byte, entrier entry.Entrier) *EntryUpdate {
	return &EntryUpdate{
		message: message{
			mType: mTypeEntryUpdate,
		},
		entryID: id,
		entrySN: sn,
		entrier: entrier,
	}
}

//MarshalMessage implements Marshaler for Network Table Messages.
func (eu *EntryUpdate) MarshalMessage() ([]byte, error) {
	entryValue, err := eu.entrier.MarshalEntry()
	if err != nil {
		return nil, err
	}
	var output []byte
	output = append(output, eu.Type())
	output = append(output, eu.entryID[:]...)
	output = append(output, eu.entrySN[:]...)
	output = append(output, eu.entrier.Type())
	output = append(output, entryValue...)
	return nil, nil
}
