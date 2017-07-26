// Copyright (c) 2017, Technomancers. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goNTCore

//LogMessage is used to send information back to the running process about the server.
type LogMessage struct {
	Message string
	Err     error
}

//NewLogMessage creates a log message with the message filled out.
func NewLogMessage(msg string) LogMessage {
	return LogMessage{
		Message: msg,
	}
}

//NewErrorMessage creates a new log message with the error filled out.
func NewErrorMessage(err error) LogMessage {
	return LogMessage{
		Err: err,
	}
}
