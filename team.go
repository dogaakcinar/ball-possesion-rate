// team.go

package main

import "fmt"

// Team represents a football team
type Team struct {
	Name            string
	PossessionRate  float64
	PreviousRate    float64
	HasPossession   bool
}

// NewTeam creates a new team with initial possession rate
func NewTeam(name string) *Team {
	return &Team{
		Name:           name,
		PossessionRate: 100.0,
		HasPossession:  false,
	}
}

// SwitchPossession switches the possession to the current team
func (t *Team) SwitchPossession(previousTeam *Team) {
	t.HasPossession = true
	t.PreviousRate = previousTeam.PossessionRate
	previousTeam.HasPossession = false
	previousTeam.PossessionRate /= 2
}

// PrintPossession prints the possession rates of both teams
func (t *Team) PrintPossession(otherTeam *Team) {
	fmt.Printf("Possession: %s - %.2f%% | %s - %.2f%%\n",
		t.Name, t.PossessionRate, otherTeam.Name, otherTeam.PossessionRate)
}
