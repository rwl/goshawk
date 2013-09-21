package goshawk

// Interface for all backends.
type Base interface {
	IsView() bool
	Elements() interface{}
}

type Core struct {
	isView bool // Whether the receiver is a view or not.
}

// Returns whether the receiver is a view or not.
func (v *Core) IsView() bool {
	return v.isView
}
