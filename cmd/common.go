package cmd

import (
	"fmt"
	"os"
	"strings"

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

func printError(a interface{}) error {
	switch v := a.(type) {
	case error:
		// Remove that annoying extra newline
		msg := strings.TrimSpace(v.Error())
		fmt.Fprintf(os.Stderr, "%s\n", msg)
	default:
		fmt.Fprintf(os.Stderr, "%v\n", a)
	}
	return nil
}

func printErrorf(format string, a interface{}) error {
	_, err := fmt.Fprintf(os.Stderr, format, a)
	return err
}

func userNames(users []rest.User) []string {
	names := make([]string, 0, 0)

	for _, u := range users {
		names = append(names, u.Username)
	}

	return names
}
