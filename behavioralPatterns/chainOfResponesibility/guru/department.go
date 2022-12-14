package behavioralPatterns

type Department interface {
	Execute(*Patient)
	SetNext(Department)
}
