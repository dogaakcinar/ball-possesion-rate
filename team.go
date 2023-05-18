// team.go

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// Team represents a football team
type Team struct {
	Name           string
	PossessionRate float64
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

func calculatePossessionRate(elapsedTime time.Duration, totaltime time.Duration) float64 {
	const maxPossessionRate = 100.0
	var totalTime = totaltime

	elapsedSeconds := elapsedTime.Seconds()
	elapsedRate := elapsedSeconds / totalTime.Seconds()
	possessionRate := maxPossessionRate * elapsedRate

	return possessionRate
}

func startGame() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Press Enter to start the game")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	if input == "\n" {
		fmt.Println("Starting the game...")
	} else {
		fmt.Println("Invalid input. Please try again.")
		startGame()
	}
}