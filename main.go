package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dogaakcinar/ball-possesion-rate/internal/match"
	"github.com/dogaakcinar/ball-possesion-rate/internal/team"
)

func main() {
	newMatch, err := match.StartMatch()
	if err != nil {
		fmt.Println(err)
	}

	// Read user input
	reader := bufio.NewReader(os.Stdin)
	for {
		// Print the possession rates
		match.CalculatePossessionRate(newMatch)
		
		// Wait for the operator to change possession or quit
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
			team1.SwitchPossession(team2)
		} else {
			team2Time += elapsedTime
			team2.SwitchPossession(team1)
		}

		// Calculate the elapsed time

		// Calculate possession rates based on elapsed time
		team1PossessionRate, team2PossessionRate := match.CalculatePossessionRate(team1Time, team2Time)
		team1.PossessionRate = team1PossessionRate
		team2.PossessionRate = team2PossessionRate
	}

	fmt.Printf("Match summary: \n"+
		"Total Time = %.2f \n"+
		"Team1 Time = %.2f |"+
		"Team2 Time = %.2f\n", team1Time.Seconds()+team2Time.Seconds(), team1Time.Seconds(), team2Time.Seconds())
	team1.PrintPossession(team2)

	fmt.Println("Thank you for using the football possession calculator!")
}
