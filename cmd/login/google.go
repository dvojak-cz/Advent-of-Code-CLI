package login

import (
	"github.com/dvojak-cz/Advent-of-Code-CLI/pkg/cred"
	"github.com/spf13/cobra"
)

var googleCmd = &cobra.Command{
	Use:     "google",
	Aliases: []string{"gg", "g"},
	Short:   "Login to Advent of Code via Google",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cred.StoreSessionId("googleCmd")
	},
}
