package structuralPatterns

import "sync"

func NewTeamFactory() teamFlyweightFactory {
	return teamFlyweightFactory{
		createdTeams: make(map[int]*Team),
	}
}

type teamFlyweightFactory struct {
	l            sync.RWMutex
	createdTeams map[int]*Team
}

func (t *teamFlyweightFactory) GetTeam(teamID int) *Team {
	t.l.Lock()
	defer t.l.Unlock()

	if t.createdTeams[teamID] != nil {
		return t.createdTeams[teamID]
	}

	team := getTeamFactory(teamID)

	t.createdTeams[teamID] = &team

	return t.createdTeams[teamID]
}

func (t *teamFlyweightFactory) GetNumberOfObjects() int {
	t.l.Lock()
	defer t.l.Unlock()
	return len(t.createdTeams)
}

func getTeamFactory(teamID int) Team {
	switch teamID {
	case TeamB:
		return Team{
			ID:   2,
			Name: TeamB,
		}
	default:
		return Team{
			ID:   1,
			Name: TeamA,
		}
	}
}
