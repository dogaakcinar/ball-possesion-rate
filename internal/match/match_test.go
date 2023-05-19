package match

import (
	"testing"
	"time"

	"github.com/dogaakcinar/ball-possesion-rate/internal/team"
)

func TestCalculatePossessionRate(t *testing.T) {
	team1 := team.NewTeam("Team 1")
	team2 := team.NewTeam("Team 2")

	// Create a new match
	match := NewMatch(team1, team2)
	match.StartTime = time.Now()

	// Simulate time passing
	time.Sleep(time.Second)

	// Update time with ball for Team 1
	match.UpdateTimeWithBall(time.Now())

	// Calculate possession rates
	match.CalculatePossessionRate()

	// Verify that possession rates are correctly calculated
	expectedTeam1PossessionRate := 50.0
	expectedTeam2PossessionRate := 50.0
	if match.Team1PossesionRate != expectedTeam1PossessionRate {
		t.Errorf("Expected Team 1 possession rate: %.2f, got: %.2f", expectedTeam1PossessionRate, match.Team1PossesionRate)
	}
	if match.Team2PossesionRate != expectedTeam2PossessionRate {
		t.Errorf("Expected Team 2 possession rate: %.2f, got: %.2f", expectedTeam2PossessionRate, match.Team2PossesionRate)
	}
}
