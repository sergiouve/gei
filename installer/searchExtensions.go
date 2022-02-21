package installer

import (
	"github.com/buger/jsonparser"
	"gitlab.com/yugarinn/gei/client"
)

type PrintableSearchResult struct {
	Name string
	Description string
	Pk string
}

func FetchSearch(searchTerm string) []PrintableSearchResult {
	searchResults := client.FetchSearch(searchTerm)

	extensions, _, _, _ := jsonparser.Get([]byte(searchResults), "extensions")
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


	return printableResults
}
