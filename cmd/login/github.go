package login

import (
	"github.com/dvojak-cz/Advent-of-Code-CLI/pkg/cred"
	"github.com/spf13/cobra"
)

var githubCmd = &cobra.Command{
	Use:     "github",
	Aliases: []string{"gh", "g"},
	Short:   "Login to Advent of Code via Github",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cred.StoreSessionId("gh")
	},
}
