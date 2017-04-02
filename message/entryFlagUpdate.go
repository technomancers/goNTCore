package message

import "io"

const (
	flagPersistantMask byte = 0x01
)

//EntryFlagUpdate is used to update the flag of an entry.
type EntryFlagUpdate struct {
	message
	entryID   [2]byte
	persitant bool
}

//NewEntryFlagUpdate creates a new instance of EntryFlagUpdate.
func NewEntryFlagUpdate(id [2]byte, persistant bool) *EntryFlagUpdate {
	return &EntryFlagUpdate{
		message: message{
			mType: mTypeEntryFlagUpdate,
		},
		entryID:   id,
		persitant: persistant,
	}
}

//MarshalMessage implements Marshaler for Network Table Messages.
func (efu *EntryFlagUpdate) MarshalMessage(writer io.Writer) error {
	flags := byte(0x00)
	if efu.persitant {
		flags = flags | flagPersistantMask
	}
	_, err := writer.Write([]byte{efu.Type()})
	if err != nil {
		return err
	}
	_, err = writer.Write(efu.entryID[:])
	if err != nil {
		return err
	}
	_, err = writer.Write([]byte{flags})
	return err
}
