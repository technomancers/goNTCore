// Copyright (c) 2017, Technomancers. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goNTCore

//Data is the interface implemented by types that can read and write Network Table at a low level.
type Data interface {
	PutEntry(ent *Entry) error //Creates if new, updates otherwise
	GetEntries(root string) ([]string, error)
	GetEntry(key string) (*Entry, error)
	DeleteEntry(key string) error
	DeleteAll(root string) error
	IsTable(key string) bool
	IsKey(key string) bool
}
