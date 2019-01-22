package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"

	"github.com/clockworksoul/cog2/data/rest"
	"github.com/spf13/cobra"
)

const (
	bootstrapUse   = "bootstrap"
	bootstrapShort = "Bootstrap a new server."
	bootstrapLong  = `Bootstrap can be used on a brand new server to create the administrative user (admin).
	
It will only work on an instance that doesn't yet have any users defined.`
)

var (
	flagBootstrapEmail    string
	flagBootstrapName     string
	flagBootstrapPassword string
	flagBootstrapProfile  string
	flagBootstrapUser     string
)

// GetBootstrapCmd bootstrap
func GetBootstrapCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   bootstrapUse,
		Short: bootstrapShort,
		Long:  bootstrapLong,
		RunE:  bootstrapCmd,
		Args:  cobra.ExactArgs(1),
	}

	cmd.Flags().StringVarP(&flagBootstrapEmail, "email", "e", "", "Email for the bootstrapped user")
	cmd.Flags().StringVarP(&flagBootstrapName, "name", "n", "", "Full name of the bootstrapped user")
	cmd.Flags().StringVarP(&flagBootstrapPassword, "password", "p", "", "Password for bootstrapped user")
	cmd.Flags().StringVarP(&flagBootstrapProfile, "profile", "r", "", "Name to use for the server profile")
	cmd.Flags().StringVarP(&flagBootstrapUser, "user", "u", "", "Username for the bootstrapped user")

	return cmd
}

func bootstrapCmd(cmd *cobra.Command, args []string) error {
	hostURL, err := bootstrapCmdParseServer(args[0])
	if err != nil {
		log.Fatalln(err)
	}

	// Does the user want a specific email?
	// If not, Cog will apply the default ("cog@localhost") and return it
	email := flagBootstrapEmail

	// Does the user want a specific password?
	// If not, Cog will generate one automatically and return it
	password := flagBootstrapPassword

	// Does the user want a specific name?
	// If not, Cog will generate one automatically and return it
	fullName := flagBootstrapName

	// The default profile name is the hostname of the Cog server
	profile := flagBootstrapProfile
	if profile == "" {
		profile = hostURL.Hostname()
	}

	// Does the user want a specific username?
	// If not, Cog will apply the default ("admin") and return it
	username := flagBootstrapUser

	endpointURL := fmt.Sprintf("%s/v2/bootstrap", hostURL.String())

	user := rest.User{
		Email:    email,
		FullName: fullName,
		Password: password,
		Username: username,
	}

	postBytes, err := json.Marshal(user)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post(endpointURL, "application/json", bytes.NewBuffer(postBytes))
	if err != nil {
		log.Fatalln(err)
	}

	switch resp.StatusCode {
	case http.StatusOK: // Everything is swell.
	case http.StatusConflict:
		fmt.Printf("Server %s has already been bootstrapped.\n", hostURL.String())
		return nil
	case http.StatusInternalServerError:
		fmt.Println("Internal server error. Check the server logs for details.")
		return nil
	default:
		bytes, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("Unexpected response: %d %v\n", resp.StatusCode, string(bytes))
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// Re-using "user" instance. Sorry.
	err = json.Unmarshal(body, &user)

	log.Println(user)

	return nil
}

func bootstrapCmdParseServer(serverURLArg string) (*url.URL, error) {
	serverURLString := serverURLArg

	// Does the URL have a prefix? If not, assume 'http://'
	matches, err := regexp.MatchString("^[a-z0-9]+://.*", serverURLString)
	if err != nil {
		return nil, err
	}
	if !matches {
		serverURLString = "http://" + serverURLString
	}

	// Parse the resulting URL
	serverURL, err := url.Parse(serverURLString)
	if err != nil {
		return nil, err
	}

	return serverURL, nil
}
