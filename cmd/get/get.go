package get

import (
	"net/http"
	"net/http/cookiejar"
	"time"
	"github.com/dvojak-cz/Advent-of-Code-CLI/pkg/cred"
	"github.com/dvojak-cz/Advent-of-Code-CLI/pkg/defs"
	"github.com/spf13/cobra"
)

var (
	FlagStdOut bool
	FlagYear   int
	FlagDay    int
	client     *http.Client
	loggedIn   bool
)

func init() {
	now := time.Now()
	defaultDay := now.Day()
	if defaultDay > 25 {
		defaultDay = 25
	}
	defaultYear := now.Year()
	if now.Month() != 12 {
		defaultYear -= 1
	}
	Cmd.PersistentFlags().IntVarP(&FlagYear, "year", "y", defaultYear, "Year")
	Cmd.PersistentFlags().IntVarP(&FlagDay, "day", "d", defaultDay, "Day")
	Cmd.PersistentFlags().BoolVarP(&FlagStdOut, "std-out", "s", false, "")
	Cmd.AddCommand(puzzleCmd)
	Cmd.AddCommand(dataCmd)

	logIn()
}

func logIn() {
	jar, _ := cookiejar.New(nil)
	id, err := cred.GetSessionId()
	loggedIn = false

	if err == nil {
		loggedIn = true
		cookie := &http.Cookie{
			Name:  "session",
			Value: id,
		}
		jar.SetCookies(defs.BaseUrl, []*http.Cookie{cookie})
	}

	// Create HTTP client with cookie jar
	client = &http.Client{
		Jar: jar,
	}
}

var Cmd = &cobra.Command{
	Use:   "get",
	Short: "Get input for a day",
}
