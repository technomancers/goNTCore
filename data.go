package goNTCore

//Data is the interface implemented by types that can read and write Network Table at a low level.
type Data interface {
	PutEntry(ent *Entry) error //Creates if new, updates otherwise
	GetEntries() ([]*Entry, error)
	GetEntry(key string) (*Entry, error)
	DeleteEntry(key string) error
	DeleteAll() error
	IsTable(key string) bool
	IsKey(key string) bool
}
