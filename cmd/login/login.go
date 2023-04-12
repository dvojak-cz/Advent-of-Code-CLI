package login

import (
	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(githubCmd)
	Cmd.AddCommand(googleCmd)
	Cmd.AddCommand(redditCmd)
	Cmd.AddCommand(twitterCmd)
}

var Cmd = &cobra.Command{
	Use:   "login",
	Short: "Login to Advent of Code",
	Long: `Login to Advent of Code
		you can login to via [github|google|reddit|twitter]`,
}
