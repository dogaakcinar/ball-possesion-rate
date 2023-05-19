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
	refreshRate := time.Second / 20

	// Create a match
	m := NewMatch(team1.Name, team2.Name)

	go func() {
		for {
			time.Sleep(refreshRate)
			m.UpdateTimeWithBall(refreshRate)
		}
	}()

	// Switch possession and update time with ball
	time.Sleep(time.Millisecond * 5000)
	t.Logf("%2f", m.Team1TimeWithBall.Seconds())
	m.SwitchPossession()
	time.Sleep(time.Millisecond * 5000)
	t.Logf("%2f", m.Team2TimeWithBall.Seconds())
	t.Logf("%2f", time.Since(m.StartTime).Seconds())
	// Verify the possession rates

		// Sleep for a while to allow the goroutine to update the values
	time.Sleep(refreshRate)
	
	expectedMinRate := 49.8
	expectedMaxRate := 50.2
	m.CalculatePossessionRate()
	if m.Team1PossesionRate < expectedMinRate || m.Team1PossesionRate > expectedMaxRate {
		t.Errorf("Team possession rate out of range. Expected range: %.2f - %.2f, Got: %.2f", expectedMinRate, expectedMaxRate, m.Team2PossesionRate)
		m.PrintPossesion()
	} else {
		t.Logf("Team 1 possesion rate: %.3f \n Team 2 possesion rate: %.3f ", m.Team1PossesionRate, m.Team2PossesionRate)
	}

}
