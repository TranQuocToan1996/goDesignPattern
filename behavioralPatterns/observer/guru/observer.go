package behavioralPatterns

type Observer interface {
	update(string)
	getID() string
}
