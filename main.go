// main.go

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Create two teams
	team1 := NewTeam("Team 1")
	team2 := NewTeam("Team 2")

	// The starting team has the kickoff possession
	team1.HasPossession = true

	// Read user input
	reader := bufio.NewReader(os.Stdin)
	for {
		// Print the possession rates
		team1.PrintPossession(team2)

		// Wait for the operator to change possession
		fmt.Print("Press Enter to switch possession or 'q' to quit: ")
		input, _ := reader.ReadString('\n')

		if strings.TrimSpace(input) == "q" {
			fmt.Println("Exiting...")
			break
		}

		// Switch the possession to the other team
		if team1.HasPossession {
			team2.SwitchPossession(team1)
		} else {
			team1.SwitchPossession(team2)
		}
	}

	fmt.Println("Thank you for using the football possession calculator!")
}
