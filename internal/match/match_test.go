package match

import (
	"testing"
	"time"

	"github.com/dogaakcinar/ball-possesion-rate/internal/team"
)

func TestMatch(t *testing.T) {
	// Create two teams
	team1 := team.NewTeam("Team 1")
	team2 := team.NewTeam("Team 2")
	refreshRate := time.Second / 3
	ch := make(chan int)

	// Create a match
	m := NewMatch(team1.Name, team2.Name)

	go func(chan int) {
		for {
			select {
			case <-ch:
				time.Sleep(refreshRate)
				m.UpdateTimeWithBall(refreshRate)
				t.Logf("in chan")
			default:
				time.Sleep(refreshRate)
				m.UpdateTimeWithBall(refreshRate)
				t.Logf("in default")
			}
		}
	}(ch)

	// Switch possession and update time with ball
	time.Sleep(time.Millisecond * 2000)
	ch <- 1
	m.SwitchPossession()
	time.Sleep(time.Millisecond * 2000)
	ch <- 1

	t.Logf("team1:%2f", m.Team1TimeWithBall.Seconds())
	t.Logf("team2:%2f", m.Team2TimeWithBall.Seconds())
	t.Logf("total: %2f", time.Since(m.StartTime).Seconds()) // Verify the possession rates
	expectedMinRate := 49.8
	expectedMaxRate := 50.2
	m.CalculatePossessionRate()
	if m.Team1PossesionRate < expectedMinRate || m.Team1PossesionRate > expectedMaxRate {
		t.Errorf("Team possession rate out of range. Expected range: %.2f - %.2f, Got: %.2f", expectedMinRate, expectedMaxRate, m.Team2PossesionRate)
		m.PrintPossesion()
	} else {
		t.Logf("Team 1 possession rate: %.3f\nTeam 2 possession rate: %.3f", m.Team1PossesionRate, m.Team2PossesionRate)
	}
}
