package get

import (
	"fmt"
	"github.com/dvojak-cz/Advent-of-Code-CLI/pkg/defs"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
)

var (
	FlagDataOutputFile string
)

func init() {
	dataCmd.PersistentFlags().StringVarP(&FlagDataOutputFile, "file", "f", "data.txt", "")
}

var dataCmd = &cobra.Command{
	Use:   "data",
	Short: "Get input for a day",
	RunE: func(cmd *cobra.Command, args []string) error {
		if !loggedIn {
			return fmt.Errorf("please log in first")
		}

		req, err := createInputDataRequest(FlagYear, FlagDay)
		if err != nil {
			return err
		}

		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("could not get input data for day %d: %s", FlagDay, resp.Status)
		}
		var stream io.Writer
		if FlagStdOut {
			stream = io.Writer(os.Stdout)
		} else {
			f, err := os.OpenFile(FlagDataOutputFile, os.O_RDWR|os.O_CREATE, 0644)
			if err != nil {
				return err
			}
			defer f.Close()
			stream = io.Writer(f)
		}

		_, err = io.Copy(stream, resp.Body)
		if err != nil {
			return err
		}
		return nil
	},
}

func createInputDataRequest(year, day int) (*http.Request, error) {
	path := fmt.Sprintf("/%d/day/%d/input", year, day)
	return http.NewRequest("GET", defs.BaseUrl.String()+path, nil)
}
