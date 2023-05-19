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
	var team1Time time.Duration
	var team2Time time.Duration
	teamNames, _ := startGame()
	team1 := NewTeam(teamNames[0])
	team2 := NewTeam(teamNames[1])
	// The starting team has the kickoff possession
	team1.HasPossession = true

	// Read user input
	reader := bufio.NewReader(os.Stdin)
	for {
		// Print the possession rates
		team1.PrintPossession(team2)
		startTime := time.Now()

		// Wait for the operator to change possession or timeout
		fmt.Print("Press Enter to switch possession or 'q' to quit: ")
		input, _ := reader.ReadString('\n')

		if strings.TrimSpace(input) == "q" {
			fmt.Println("Exiting...")
			break
		}

		elapsedTime := time.Since(startTime)

		// Switch the possession to the other team
		if team1.HasPossession {
			team1Time += elapsedTime
			team2.SwitchPossession(team1)
		} else {
			team2Time += elapsedTime
			team1.SwitchPossession(team2)
		}

		// Calculate the elapsed time

		// Calculate possession rates based on elapsed time
		team1.PossessionRate = calculatePossessionRate(team1Time, team1Time+team2Time)
		team2.PossessionRate = calculatePossessionRate(team2Time, team1Time+team2Time)
	}
	fmt.Printf("Match summary: \n"+
		"Total Time = %.2f \n"+
		"Team1 Time = %.2f |"+
		"Team2 Time = %.2f", team1Time.Seconds()+team2Time.Seconds(), team1Time.Seconds(), team2Time.Seconds())
	team1.PrintPossession(team2)
	fmt.Println("Thank you for using the football possession calculator!")
}
