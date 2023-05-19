// team.go

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// Team represents a football team
type Team struct {
	Name           string
	PossessionRate float64
	HasPossession  bool
}

func NewTeam(name string) *Team {
	team := &Team{
		Name:          name,
		HasPossession: false,
	}
	return team
}

// SwitchPossession switches the possession to the current team
func (t *Team) SwitchPossession(previousTeam *Team) {
	t.HasPossession = true
	previousTeam.HasPossession = false
}

// PrintPossession prints the possession rates of both teams
func (t *Team) PrintPossession(otherTeam *Team) {
	fmt.Printf("Possession: %s - %.2f%% | %s - %.2f%%\n",
		t.Name, t.PossessionRate, otherTeam.Name, otherTeam.PossessionRate)
	if t.HasPossession {
		fmt.Println("Team with possesion: ", t.Name)
	} else {
		fmt.Println("Team with possesion: ", otherTeam.Name)
	}
}


