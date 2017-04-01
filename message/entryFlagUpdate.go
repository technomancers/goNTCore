package message

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
func (efu *EntryFlagUpdate) MarshalMessage() ([]byte, error) {
	flags := byte(0x00)
	if efu.persitant {
		flags = flags | flagPersistantMask
	}
	var output []byte
	output = append(output, efu.Type())
	output = append(output, efu.entryID[:]...)
	output = append(output, flags)
	return output, nil
}
