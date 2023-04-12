package get

import (
	"fmt"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/anaskhan96/soup"
	"github.com/dvojak-cz/Advent-of-Code-CLI/pkg/defs"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
	"strings"
)

var (
	FlagArticleOutputFile string
)

func init() {
	puzzleCmd.PersistentFlags().StringVarP(&FlagArticleOutputFile, "file", "f", "README.md", "")
}

var puzzleCmd = &cobra.Command{
	Use:   "puzzle",
	Short: "Get puzzle for a day",
	RunE: func(cmd *cobra.Command, args []string) error {
		articles, err := getArticles(FlagYear, FlagDay)
		if err != nil {
			return err
		}

		var stream io.Writer
		if FlagStdOut {
			stream = io.Writer(os.Stdout)
		} else {
			f, err := os.OpenFile(FlagArticleOutputFile, os.O_RDWR|os.O_CREATE, 0644)
			if err != nil {
				return err
			}
			defer f.Close()
			stream = io.Writer(f)
		}

		converter := md.NewConverter("", true, nil)

		var mdArticles []string
		for _, art := range articles {
			markdown, err := converter.ConvertString(art.HTML())
			if err != nil {
				return err
			}
			mdArticles = append(mdArticles, markdown)
		}

		_, err = fmt.Fprintln(stream, strings.Join(mdArticles, "\n\n---\n\n"))
		if err != nil {
			return err
		}
		return nil
	},
}

func createArticleRequest(year, day int) (*http.Request, error) {
	path := fmt.Sprintf("/%d/day/%d", year, day)
	return http.NewRequest("GET", defs.BaseUrl.String()+path, nil)
}

func getArticles(year, day int) ([]soup.Root, error) {
	req, err := createArticleRequest(year, day)
	if err != nil {
		return []soup.Root{}, err
	}
	res, err := client.Do(req)
	if err != nil {
		return []soup.Root{}, err
	}

	if res.StatusCode != http.StatusOK {
		return []soup.Root{}, fmt.Errorf("invalid status code: %d", res.StatusCode)
	}

	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return []soup.Root{}, err
	}
	doc := soup.HTMLParse(string(b))
	article := doc.FindAll("article")
	return article, nil
}
