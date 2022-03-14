package client

import (
	"io"
	"os"
	"fmt"
	"net/http"
	"io/ioutil"
)

type QueryParameter struct {
	name  string
	value string
}

func fetch(method string, url string, params []QueryParameter) []byte {
	client := &http.Client{}
	request, _ := http.NewRequest(method, url, nil)

	if (len(params) > 0 && method == "GET") {
		query := request.URL.Query()

		for _, param := range params {
			query.Add(param.name, param.value)
		}

		request.URL.RawQuery = query.Encode()
	}

	response, _ := client.Do(request)
	payload, _ := ioutil.ReadAll(response.Body)

	return payload
}

func FetchSearch(extensionName string) []byte {
	params := []QueryParameter{(QueryParameter{name: "search", value: extensionName})}
	searchResult := fetch("GET", "https://extensions.gnome.org/extension-query", params)

	return searchResult
}

func DownloadExtension(extensionId string) {
	// TODO: first call to https://extensions.gnome.org/extension-info/?pk=79&shell_version=40 and then extract the metadata to call
	// download-extension/hide-dash@zacbarton.com.shell-extension.zip?version_tag=1993
	extensionMetadata := ""
	fmt.Println(extensionMetadata)

	client := &http.Client{}

	request, _ := http.NewRequest("GET", "https://extensions.gnome.org/download-extension/pomodoro@arun.codito.in.shell-extension.zip", nil)
	query := request.URL.Query()

	query.Add("version_tag", "41.0")
	request.URL.RawQuery = query.Encode()

	response, err := client.Do(request)

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	extensionZip, err := os.Create("")

	if err != nil {
		fmt.Println(err)
	}

	io.Copy(extensionZip, response.Body)

	defer extensionZip.Close()
}
