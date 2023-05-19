package match

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dogaakcinar/ball-possesion-rate/internal/team"
)

type Match struct {
	Team1              *team.Team
	Team2              *team.Team
	Team1TimeWithBall  time.Duration
	Team2TimeWithBall  time.Duration
	Team1WithPossesion bool
	Team2WithPossesion bool
	Team1PossesionRate float64
	Team2PossesionRate float64
	IsRunning          bool
	StartTime          time.Time
	CancelChan         chan bool
}

func NewMatch(team1Name string, team2Name string) *Match {
	return &Match{
		Team1:              team.NewTeam(team1Name),
		Team2:              team.NewTeam(team2Name),
		Team1TimeWithBall:  0,
		Team2TimeWithBall:  0,
		IsRunning:          false,
		Team1WithPossesion: true,
		Team2WithPossesion: false,
		Team1PossesionRate: 0,
		Team2PossesionRate: 0,
		StartTime:          time.Now(),
		CancelChan:         make(chan bool),
	}
}

func (m *Match) SwitchPossession() {
	m.Team1WithPossesion = !m.Team1WithPossesion
	m.Team2WithPossesion = !m.Team2WithPossesion
}

func (m *Match) PrintPossesion() {
	possessingTeam := m.Team1.Name
	if !m.Team1WithPossesion {
		possessingTeam = m.Team2.Name
	}

	fmt.Printf("Possession: %s - %.2f%% | %s - %.2f%%\n",
		m.Team1.Name, m.Team1PossesionRate, m.Team2.Name, m.Team2PossesionRate)
	fmt.Println("Team with possession:", possessingTeam)
}

func (m *Match) UpdateTimeWithBall(chTime time.Time) {
	if m.Team1WithPossesion {
		m.Team1TimeWithBall += time.Since(chTime)
	} else {
		m.Team2TimeWithBall += time.Since(chTime)
	}
}

func (m *Match) CalculatePossessionRate() {
	const maxPossessionRate = 100.0
	var totalTime = time.Since(m.StartTime).Seconds()
	m.Team1PossesionRate = (m.Team1TimeWithBall.Seconds() / totalTime) * maxPossessionRate
	m.Team2PossesionRate = 100.0 - m.Team1PossesionRate
}

func StartMatch() (*Match, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Press Enter to start the game")
	input, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	if input == "\n" {
		fmt.Println("Starting the game...")
	} else {
		fmt.Println("Invalid input. Please try again.")
		return StartMatch()
	}
	fmt.Println("Press Enter Team1's name: ")
	team1Name, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	fmt.Println("Press Enter Team2's name:")
	team2Name, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	match := NewMatch(strings.TrimSpace(team1Name), strings.TrimSpace(team2Name))
	return match, nil
}
