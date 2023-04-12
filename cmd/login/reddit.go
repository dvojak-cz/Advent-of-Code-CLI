package login

import (
	"github.com/dvojak-cz/Advent-of-Code-CLI/pkg/cred"
	"github.com/spf13/cobra"
)

var redditCmd = &cobra.Command{
	Use:     "reddit",
	Aliases: []string{"rd", "r"},
	Short:   "Login to Advent of Code via Reddit",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cred.StoreSessionId("redditCmd")
	},
}
