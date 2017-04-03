package goNTCore

import "github.com/technomancers/goNTCore/util"

//NetworkTabler is the interface implemented by types that can read and write Network Table at a high level.
type NetworkTabler interface {
	ContainsKey(key string) bool
	ContainsSubTable(key string) bool
	Delete(key string)
	DeleteAll()
	GetFlags(key string) byte
	IsPersisted(key string) bool
	GetKeys() []string
	GetTable(key string) Table
	GetBoolean(key string, def bool) bool
	PutBoolean(key string, val bool) bool
	GetNumber(key string, def float64) float64
	PutNumber(key string, val float64) bool
	GetString(key string, def string) string
	PutString(key string, val string) bool
	GetRaw(key string, def []byte) []byte
	PutRaw(key string, val []byte) bool
	GetBooleanArray(key string, def []bool) []bool
	PutBooleanArray(key string, val []bool) bool
	GetNumberArray(key string, def []float64) []float64
	PutNumberArray(key string, val []float64) bool
	GetStringArray(key string, def []string) []string
	PutStringArray(key string, val []string) bool
}

//Table implements NetworkTabler based on a backend that needs to exist on create.
type Table struct {
	data Data
	root string
}

//NewTable creates a new Network Table that is based off of the data and root passed in.
//Pass in an empty string or "/" for root table.
func NewTable(d Data, root string) *Table {
	if root == "" {
		root = "/"
	}
	root = util.SanatizeKey(root)
	return &Table{
		data: d,
		root: root,
	}
}
