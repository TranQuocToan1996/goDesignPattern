package behavioralPatterns

type Mediator interface {
	canArrive(Train) bool
	notifyAboutDeparture()
}
