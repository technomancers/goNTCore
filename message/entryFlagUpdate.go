// Copyright (c) 2017, Technomancers. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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
			mType: MTypeEntryFlagUpdate,
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

//UnmarshalMessage implements Unmarshaler for Network Table Messages and assumes the message type byte has already been read.
func (efu *EntryFlagUpdate) UnmarshalMessage(reader io.Reader) error {
	efu.mType = MTypeEntryFlagUpdate
	idBuf := make([]byte, 2)
	flagBuf := make([]byte, 1)

	_, err := io.ReadFull(reader, idBuf)
	if err != nil {
		return err
	}
	_, err = io.ReadFull(reader, flagBuf)
	if err != nil {
		return err
	}

	copy(efu.entryID[:], idBuf)
	efu.persitant = flagBuf[0]&flagPersistantMask == flagPersistantMask
	return nil
}
