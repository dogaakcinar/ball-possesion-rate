// team.go

package main

import (
	"fmt"
	"time"
)

// Team represents a football team
type Team struct {
	Name           string
	PossessionRate float64
	PreviousRate   float64
	HasPossession  bool
	IsHomeTeam     bool
}

func NewTeam(name string, isHomeTeam bool) *Team {
	team := &Team{
		Name:          name,
		HasPossession: isHomeTeam,
		IsHomeTeam:    isHomeTeam,
	}

	if isHomeTeam {
		team.PossessionRate = 100
	} else {
		team.PossessionRate = 0
	}

	return team
}

// SwitchPossession switches the possession to the current team
func (t *Team) SwitchPossession(previousTeam *Team) {
	t.HasPossession = true
	t.PreviousRate = previousTeam.PossessionRate

	previousTeam.HasPossession = false
	previousTeam.PossessionRate = 100.0 - t.PreviousRate
}

// PrintPossession prints the possession rates of both teams
func (t *Team) PrintPossession(otherTeam *Team) {
	fmt.Printf("Possession: %s - %.2f%% | %s - %.2f%%\n",
		t.Name, t.PossessionRate, otherTeam.Name, otherTeam.PossessionRate)
}

func calculatePossessionRate(elapsedTime time.Duration, isHomeTeam bool) float64 {
	// Calculate the possession rate based on elapsed time
	const maxPossessionRate = 100.0
	const totalTime = 10 * time.Second

	if elapsedTime >= totalTime {
		return maxPossessionRate
	}

	elapsedSeconds := float64(elapsedTime.Seconds())
	elapsedRate := elapsedSeconds / totalTime.Seconds()
	possessionRate := maxPossessionRate * elapsedRate

	if !isHomeTeam {
		possessionRate = maxPossessionRate - possessionRate
	}

	return possessionRate
}
