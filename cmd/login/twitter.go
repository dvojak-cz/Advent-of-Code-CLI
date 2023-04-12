package login

import (
	"github.com/dvojak-cz/Advent-of-Code-CLI/pkg/cred"
	"github.com/spf13/cobra"
)

var twitterCmd = &cobra.Command{
	Use:     "twitter",
	Aliases: []string{"tw", "t"},
	Short:   "Login to Advent of Code via Twitter",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cred.StoreSessionId("twitterCmd")
	},
}
