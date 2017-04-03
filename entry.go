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
