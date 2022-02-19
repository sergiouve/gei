package cmd

import (
	"github.com/spf13/cobra"
	"github.com/buger/jsonparser"
	"fmt"
	"os"
	"gitlab.com/yugarinn/gei/client"
)

type PrintableSearchResult struct {
	Name string
	Description string
	Pk string
}

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search available extensions in https://extensions.gnome.org",
	Long: ` A longer description that spans multiple lines and likely contains examples
			and usage of using your command. For example:`,
	Run: func(cmd *cobra.Command, args []string) {
		search(args)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}

func search(args []string) {
	results := fetchSearch(args)
	extensions, _, _, _ := jsonparser.Get([]byte(results), "extensions")
	printableResults := []PrintableSearchResult{}

	jsonparser.ArrayEach(extensions, func(searchResult []byte, dataType jsonparser.ValueType, offset int, err error) {
		searchName, _, _, _ := jsonparser.Get(searchResult, "name")
		searchDescription, _, _, _ := jsonparser.Get(searchResult, "description")
		searchPk, _, _, _ := jsonparser.Get(searchResult, "pk")

		printableSearchResult := PrintableSearchResult{
			Name: string(searchName),
			Description: string(searchDescription),
			Pk: string(searchPk),
		}

		printableResults = append(printableResults, printableSearchResult)
	})

	for _, result := range printableResults {
		fmt.Println(result.Pk)
		fmt.Println(result.Name)
		fmt.Println(result.Description)
		fmt.Println("---")
	}
}

func fetchSearch(args []string) string {
	if (len(args) < 1) {
		fmt.Println("A search term is required")
		os.Exit(0)
	}

	searchResult := client.FetchSearch(args[0])

	return searchResult
}
