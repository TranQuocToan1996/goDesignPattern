package behavioralPatterns

type Shape interface {
	getType() string
	accept(Visitor)
}
