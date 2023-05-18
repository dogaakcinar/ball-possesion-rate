// main.go

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// Create two teams
	team1 := NewTeam("Team 1", true)
	team2 := NewTeam("Team 2", !team1.IsHomeTeam)

	// The starting team has the kickoff possession
	team1.HasPossession = true

	// Read user input
	reader := bufio.NewReader(os.Stdin)
	for {
		// Print the possession rates
		team1.PrintPossession(team2)

		// Wait for the operator to change possession or timeout
		fmt.Print("Press Enter to switch possession or 'q' to quit: ")
		input, _ := reader.ReadString('\n')

		if strings.TrimSpace(input) == "q" {
			fmt.Println("Exiting...")
			break
		}

		// Start the timer
		startTime := time.Now()

		// Switch the possession to the other team
		if team1.HasPossession {
			team2.SwitchPossession(team1)
		} else {
			team1.SwitchPossession(team2)
		}

		// Calculate the elapsed time
		elapsedTime := time.Since(startTime)

		// Calculate possession rates based on elapsed time
		team1.PossessionRate = calculatePossessionRate(elapsedTime, team1.IsHomeTeam)
		team2.PossessionRate = calculatePossessionRate(elapsedTime, team2.IsHomeTeam)
	}

	fmt.Println("Thank you for using the football possession calculator!")
}
