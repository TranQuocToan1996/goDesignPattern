package behavioralPatterns

type Mediator interface {
	CanArrive(Train) bool
	NotifyAboutDeparture()
}
