package structuralPatterns

import "time"

type Team struct {
	ID             uint64
	Name           int
	Shield         []byte
	Players        []Player
	HistoricalData []HistoricalData
}

const (
	TeamA = 1
	TeamB = 2
)

type Player struct {
	Name         string
	Surname      string
	PreviousTeam uint64
	Photo        []byte
}
type HistoricalData struct {
	Year          uint8
	LeagueResults []Match
}

type Match struct {
	Date          time.Time
	VisitorID     uint64
	LocalID       uint64
	LocalScore    byte
	VisitorScore  byte
	LocalShoots   uint16
	VisitorShoots uint16
}
