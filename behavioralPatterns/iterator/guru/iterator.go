package behavioralPatterns

type Iterator interface {
	hasNext() bool
	getNext() *User
}
