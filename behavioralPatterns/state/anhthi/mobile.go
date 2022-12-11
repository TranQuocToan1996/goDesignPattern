package anhthi

import "fmt"

type MobileAlertState interface {
	alert()
	// TODO: way to set another state
}

type AlertStateContext struct {
	// TODO: May composite Vibration, Silence and populate in MobileAlertState var
	currentState MobileAlertState
}

func NewAlertStateContext() *AlertStateContext {
	return &AlertStateContext{currentState: &Vibration{}}
}

func (ctx *AlertStateContext) SetState(state MobileAlertState) {
	ctx.currentState = state
}

func (ctx *AlertStateContext) Alert() {
	ctx.currentState.alert()
}

type Vibration struct{}

func (v *Vibration) alert() {
	fmt.Println("vibrating....")
}

type Silence struct{}

func (s *Silence) alert() {
	fmt.Println("silent ....")
}
