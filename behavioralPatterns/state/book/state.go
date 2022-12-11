package behavioralPatterns

type GameState interface {
	executeState(*GameContext) bool
}
type GameContext struct {
	SecretNumber int
	Retries      int
	Won          bool
	Next         GameState
}
