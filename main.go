package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"github.com/dogaakcinar/ball-possesion-rate/internal/match"
)

func main() {
	m, err := match.StartMatch()
	if err != nil {
		fmt.Println(err)
	}

	// Read user input
	reader := bufio.NewReader(os.Stdin)
	for {
		// Print the possession rates
		m.CalculatePossessionRate()
		chTime := time.Now()
		fmt.Println(m)
		// Wait for the operator to change possession or quit
		fmt.Print("Press Enter to switch possession or 'q' to quit: ")
		input, _ := reader.ReadString('\n')

		if strings.TrimSpace(input) == "q" {
			fmt.Println("Exiting...")
			break
		}
		m.UpdateTimeWithBall(chTime)
		m.SwitchPossession()
		m.PrintPossesion()
	}

	fmt.Printf("Match summary: \n"+
		"Total Time = %.2f seconds ", time.Since(m.StartTime).Seconds())

	fmt.Println("Thank you for using the football possession calculator!")
}
