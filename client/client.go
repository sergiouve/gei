package client

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"gitlab.com/yugarinn/gei/idos"
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
	params := []QueryParameter{{name: "search", value: extensionName}}
	searchResult := fetch("GET", "https://extensions.gnome.org/extension-query", params)

	return searchResult
}

func FetchExtensionMetadata(extensionId string, systemShellVersion string) []byte {
	params := []QueryParameter{{name: "pk", value: extensionId}, {name: "shell_version", value: systemShellVersion}}
	extensionMetadata := fetch("GET", "https://extensions.gnome.org/extension-info/", params)

	return extensionMetadata
}

func DownloadExtension(extensionMetadata idos.ExtensionMetadata) {
	client := &http.Client{}
	homeDir, _ := os.UserHomeDir()
	fileName := fmt.Sprintf("%s.zip", extensionMetadata.Uuid)

	downloadUrl := fmt.Sprintf("https://extensions.gnome.org%s", extensionMetadata.DownloadUrl)
	request, _ := http.NewRequest("GET", downloadUrl, nil)

	response, err := client.Do(request)

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	extensionZip, err := os.Create(filepath.Join(fmt.Sprintf("%s/.local/share/gnome-shell/extensions", homeDir), filepath.Base(fileName)))

	if err != nil {
		fmt.Println(err)
	}

	io.Copy(extensionZip, response.Body)

	defer extensionZip.Close()
}
