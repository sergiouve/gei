package installer

import (
	"github.com/buger/jsonparser"
	"gitlab.com/yugarinn/gei/client"
)

type PrintableSearchResult struct {
	Name string
	Pk string
	Description string
}

func FetchSearch(searchTerm string) []PrintableSearchResult {
	searchResults := client.FetchSearch(searchTerm)

	extensions, _, _, _ := jsonparser.Get([]byte(searchResults), "extensions")
	printableResults := []PrintableSearchResult{}

	jsonparser.ArrayEach(extensions, func(searchResult []byte, dataType jsonparser.ValueType, offset int, err error) {
		searchName, _, _, _ := jsonparser.Get(searchResult, "name")
		searchPk, _, _, _ := jsonparser.Get(searchResult, "pk")
		searchDescription, _, _, _ := jsonparser.Get(searchResult, "description")

		printableSearchResult := PrintableSearchResult{
			Name: string(searchName),
			Pk: string(searchPk),
			Description: string(searchDescription),
		}

		printableResults = append(printableResults, printableSearchResult)
	})

	return printableResults
}
