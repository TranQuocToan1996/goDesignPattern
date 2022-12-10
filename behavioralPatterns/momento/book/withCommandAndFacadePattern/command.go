package withcommandandfacadepattern

// use a Command pattern to encapsulate a set of different types of states (those that implement a Command interface)
type Command interface {
	GetValue() interface{}
}
