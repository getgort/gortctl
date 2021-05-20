package cmd

import (
	"github.com/clockworksoul/gort/data/rest"
)

var (
	// FlagGortProfile is a persistent flag
	FlagGortProfile string
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
