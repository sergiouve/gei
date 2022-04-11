package installer

import (
	"encoding/json"

	"gitlab.com/yugarinn/gei/installer/client"
)

type FetchSearchResult struct {
	Extensions []ExtensionSearchResult
}

type ExtensionSearchResult struct {
	Uuid string
	Name string
	Creator string
	CreatorUrl string
	Pk int
	Description string
	Link string
	Icon string
	ShellVersionMap []ShellVersion
}

type ShellVersion struct {
	Pk int
	Version int
}

func FetchSearch(searchTerm string) FetchSearchResult {
	searchResponse := client.FetchSearch(searchTerm)

	var extensionSearchResult FetchSearchResult
	json.Unmarshal(searchResponse, &extensionSearchResult)

	return extensionSearchResult
}
