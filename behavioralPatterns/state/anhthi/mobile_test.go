package anhthi

import "testing"

func TestMobible(t *testing.T) {
	stateContext := NewAlertStateContext()
	stateContext.Alert()
	stateContext.Alert()
	stateContext.Alert()
	stateContext.SetState(&Silence{})
	stateContext.Alert()
	stateContext.Alert()
	stateContext.Alert()
}
