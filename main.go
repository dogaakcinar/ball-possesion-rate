package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dogaakcinar/ball-possesion-rate/internal/match"
	"github.com/gin-gonic/gin"
)

func main() {
	m, err := match.StartMatch()
	if err != nil {
		fmt.Println(err)
		return
	}
	refreshRate := time.Second

	router := gin.Default()

	router.GET("/possession-rates", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			m.Team1.Name: m.Team1PossesionRate,
			m.Team2.Name: m.Team2PossesionRate,
		})
	})

	// Start a goroutine to update the possession rates periodically
	go func() {
		for {
			time.Sleep(refreshRate)
			m.UpdateTimeWithBall(refreshRate)
			m.CalculatePossessionRate()
			m.PrintPossesion()
		}
	}()

	go func() {
		router.Run(":8080")
	}()

	// Read user input
	reader := bufio.NewReader(os.Stdin)
	for {
		// Wait for the operator to change possession or quit
		fmt.Print("Press Enter to switch possession or 'q' to quit: ")
		input, _ := reader.ReadString('\n')

		if strings.TrimSpace(input) == "q" {
			fmt.Println("Exiting...")
			break
		}

		m.SwitchPossession()
	}

	fmt.Printf("Match summary:\n"+
		"Total Time = %.2f seconds\n", time.Since(m.StartTime).Seconds())

	fmt.Println("Thank you for using the football possession calculator!")
}
