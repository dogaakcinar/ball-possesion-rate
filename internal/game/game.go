package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func CalculatePossessionRate(elapsedTime time.Duration, totaltime time.Duration) float64 {
	const maxPossessionRate = 100.0
	var totalTime = totaltime
	elapsedSeconds := elapsedTime.Seconds()
	elapsedRate := elapsedSeconds / totalTime.Seconds()
	possessionRate := maxPossessionRate * elapsedRate

	return possessionRate
}

func StartGame() ([2]string, error) {
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
		StartGame()
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
