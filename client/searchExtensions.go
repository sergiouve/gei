package client

import (
	"net/http"
	"io/ioutil"
)

func FetchSearch(extensionName string) string {
	params := QueryParameter{name: "search", value: extensionName}
	searchResult := Fetch("GET", "https://extensions.gnome.org/extension-query", [params])

	return searchResult
}
