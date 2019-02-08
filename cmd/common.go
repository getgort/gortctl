package cmd

import (
	"github.com/clockworksoul/cog2/data/rest"
)

var (
	// FlagCogProfile is a persistent flag
	FlagCogProfile string
)

func groupNames(groups []rest.Group) []string {
	names := make([]string, 0, 0)

	for _, g := range groups {
		names = append(names, g.Name)
	}

	return names
}

func userNames(users []rest.User) []string {
	names := make([]string, 0, 0)

	for _, u := range users {
		names = append(names, u.Username)
	}

	return names
}
