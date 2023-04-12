package cmd

import (
	"github.com/dvojak-cz/Advent-of-Code-CLI/cmd/get"
	"github.com/dvojak-cz/Advent-of-Code-CLI/cmd/login"
	_ "github.com/dvojak-cz/Advent-of-Code-CLI/cmd/login"
	"github.com/dvojak-cz/Advent-of-Code-CLI/pkg/defs"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	RootCmd.AddCommand(login.Cmd)
	RootCmd.AddCommand(logoutCmd)
	RootCmd.AddCommand(get.Cmd)

}

var RootCmd = &cobra.Command{
	Use:   defs.AppShortName,
	Short: "Aof is a simple CLI for Advent of Code",
	Long: `Aof is a simple CLI for Advent of Code
    It allows you to:
        to download puzzles
        to download puzzle input
        to submit puzzle answers and see the results`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
