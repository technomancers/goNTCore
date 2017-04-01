package message

import (
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
func (ea *EntryAssign) MarshalMessage() ([]byte, error) {
	nameBytes, err := ea.entryName.MarshalEntry()
	if err != nil {
		return nil, err
	}
	entryValue, err := ea.entrier.MarshalEntry()
	if err != nil {
		return nil, err
	}
	flags := byte(0x00)
	if ea.entryPersistant {
		flags = flags | flagPersistantMask
	}
	var output []byte
	output = append(output, ea.Type())
	output = append(output, nameBytes...)
	output = append(output, ea.entrier.Type())
	output = append(output, ea.entryID[:]...)
	output = append(output, ea.entrySN[:]...)
	output = append(output, flags)
	output = append(output, entryValue...)
	return output, nil
}
