package structuralPatterns

import (
	"fmt"
	"testing"
)

func TestFlyWeight(t *testing.T) {

	factory := NewTeamFactory()
	teamA1 := factory.GetTeam(TeamA)
	if teamA1 == nil {
		t.Error("The pointer to the TEAM_A was nil")
	}
	teamA2 := factory.GetTeam(TeamA)
	if teamA2 == nil {
		t.Error("The pointer to the TEAM_A was nil")
	}
	if teamA1 != teamA2 {
		t.Error("TEAM_A pointers weren't the same")
	}
	if factory.GetNumberOfObjects() != 1 {
		t.Errorf("The number of objects created was not 1: %d\n",
			factory.GetNumberOfObjects())
	}
}

func Test_HighVolume(t *testing.T) {
	noViewers := 500000
	factory := NewTeamFactory()

	// Load 2 teams to the map
	factory.GetTeam(TeamA)
	factory.GetTeam(TeamB)

	noTeams := factory.GetNumberOfObjects()
	teams := make([]*Team, noViewers*noTeams)
	for i := 0; i < noViewers; i++ {
		teams[i] = factory.GetTeam(TeamA)
	}
	for i := noViewers; i < noTeams*noViewers; i++ {
		teams[i] = factory.GetTeam(TeamB)
	}
	if getTeams := factory.GetNumberOfObjects(); getTeams != noTeams {
		t.Errorf("The number of objects created was not 2: %d\n", getTeams)
	}
	// first 50 pointer
	for i := 0; i < 50; i++ {
		fmt.Printf("Pointer %d points to %p and is located in %p\n", i, teams[i],
			&teams[i])
	}
}
