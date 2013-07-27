
package colt

// Interface for all colt backends.
type BaseData interface {
	IsView() bool
	Elements() interface{}
}

type CoreData struct {
	isView bool // Whether the receiver is a view or not.
}

// Returns whether the receiver is a view or not.
func (v CoreData) IsView() bool {
	return v.isView
}
