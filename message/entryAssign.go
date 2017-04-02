package message

import (
	"io"

	"github.com/technomancers/goNTCore/entry"
)

var (
	//NewEntryID is the ID of an element when a client is creating a new entry.
	NewEntryID = [2]byte{0xff, 0xff}
	//NewEntrySN is the sequence number of an element when a client is creating a new entry.
	NewEntrySN = [2]byte{0x00, 0x00}
)

//EntryAssign is used to inform others that a new entry was introduced into the network.
type EntryAssign struct {
	message
	entryName       *entry.String
	entryID         [2]byte
	entrySN         [2]byte
	entryPersistant bool
	entrier         entry.Entrier
}

//NewEntryAssin creates a new instance on EntryAssign.
func NewEntryAssin(entryName string, entrier entry.Entrier, persistant bool, id, sn [2]byte) *EntryAssign {
	return &EntryAssign{
		message: message{
			mType: mTypeEntryAssign,
		},
		entryName:       entry.NewString(entryName),
		entrier:         entrier,
		entryPersistant: persistant,
		entryID:         id,
		entrySN:         sn,
	}
}

//MarshalMessage implements Marshaler for Network Table Messages.
func (ea *EntryAssign) MarshalMessage(writer io.Writer) error {
	flags := byte(0x00)
	if ea.entryPersistant {
		flags = flags | flagPersistantMask
	}
	_, err := writer.Write([]byte{ea.Type()})
	if err != nil {
		return err
	}
	err = ea.entryName.MarshalEntry(writer)
	if err != nil {
		return err
	}
	_, err = writer.Write([]byte{ea.entrier.Type()})
	if err != nil {
		return err
	}
	_, err = writer.Write(ea.entryID[:])
	if err != nil {
		return err
	}
	_, err = writer.Write(ea.entrySN[:])
	if err != nil {
		return err
	}
	_, err = writer.Write([]byte{flags})
	if err != nil {
		return err
	}
	err = ea.entrier.MarshalEntry(writer)
	return err
}
