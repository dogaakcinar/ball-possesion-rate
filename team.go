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

func calculatePossessionRate(elapsedTime time.Duration, totaltime time.Duration) float64 {
	const maxPossessionRate = 100.0
	var totalTime = totaltime
	elapsedSeconds := elapsedTime.Seconds()
	elapsedRate := elapsedSeconds / totalTime.Seconds()
	possessionRate := maxPossessionRate * elapsedRate

	return possessionRate
}

func startGame() ([2]string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Press Enter to start the game")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return [2]string{}, err
	}
	if input == "\n" {
		fmt.Println("Starting the game...")
	} else {
		fmt.Println("Invalid input. Please try again.")
		startGame()
	}
	fmt.Println("Press Enter Team1's name: ")
	team1Name, err := reader.ReadString('\n')
	if err != nil {
		return [2]string{}, err
	}

	fmt.Println("Press Enter Team2's name:")
	team2Name, err := reader.ReadString('\n')
	if err != nil {
		return [2]string{}, err
	}

	return [2]string{strings.TrimSpace(team1Name), strings.TrimSpace(team2Name)}, nil
}
