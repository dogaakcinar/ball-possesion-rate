// team.go

package team

// Team represents a football team
type Team struct {
	Name string
}

func NewTeam(name string) *Team {
	team := &Team{
		Name: name,
	}
	return team
}
