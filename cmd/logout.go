package cmd

import (
	"github.com/dvojak-cz/Advent-of-Code-CLI/pkg/cred"
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout from Advent of Code",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get current user
		return cred.DeleteSessionId()
	},
}
