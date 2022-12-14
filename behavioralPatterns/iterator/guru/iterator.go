package behavioralPatterns

type Iterator interface {
	HasNext() bool
	GetNext() *User
}
