package message

import (
	"io"

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
func (eu *EntryUpdate) MarshalMessage(writer io.Writer) error {
	_, err := writer.Write([]byte{eu.Type()})
	if err != nil {
		return err
	}
	_, err = writer.Write(eu.entryID[:])
	if err != nil {
		return err
	}
	_, err = writer.Write(eu.entrySN[:])
	if err != nil {
		return err
	}
	_, err = writer.Write([]byte{eu.entrier.Type()})
	if err != nil {
		return err
	}
	err = eu.entrier.MarshalEntry(writer)
	return err
}

//UnmarshalMessage implements Unmarshaler for Network Table Messages and assumes the message type byte has already been read.
func (eu *EntryUpdate) UnmarshalMessage(reader io.Reader) error {
	eu.mType = mTypeEntryUpdate
	idBuf := make([]byte, 2)
	snBuf := make([]byte, 2)
	typeBuf := make([]byte, 1)

	_, err := io.ReadFull(reader, idBuf)
	if err != nil {
		return err
	}
	_, err = io.ReadFull(reader, snBuf)
	if err != nil {
		return err
	}
	_, err = io.ReadFull(reader, typeBuf)
	if err != nil {
		return err
	}
	ent, err := entry.Unmarshal(typeBuf[0], reader)
	if err != nil {
		return err
	}

	eu.entrier = ent
	for i := 0; i < len(idBuf); i++ {
		eu.entryID[i] = idBuf[i]
	}
	for j := 0; j < len(snBuf); j++ {
		eu.entrySN[j] = snBuf[j]
	}
	return nil
}
