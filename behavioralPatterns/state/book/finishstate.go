package behavioralPatterns

type FinishState struct{}

func (f *FinishState) executeState(c *GameContext) bool {
	if c.Won {
		println("Congrats, you won")
	} else {
		println("You lose")
	}
	return false
}
