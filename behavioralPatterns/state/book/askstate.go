package behavioralPatterns

import (
	"fmt"
	"os"
)

type AskState struct{}

func (a *AskState) executeState(c *GameContext) bool {
	fmt.Printf("Introduce a number between 0 and 10, you have %d tries left\n", c.Retries)
	var n int
	fmt.Fscanf(os.Stdin, "%d", &n)
	c.Retries = c.Retries - 1
	if n == c.SecretNumber {
		c.Won = true
		c.Next = &FinishState{}
	}
	if c.Retries == 0 {
		c.Next = &FinishState{}
	}
	return true
}
