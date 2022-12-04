package behavioralPatterns

type Department interface {
	execute(*Patient)
	setNext(Department)
}
