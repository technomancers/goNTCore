// Copyright (c) 2017, Technomancers. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goNTCore

//Entry is a key value in the network table.
type Entry struct {
	ID        uint16
	SN        uint16
	Persitant bool
	Key       string
	Type      byte
	Value     interface{}
}
