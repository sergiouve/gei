package client

import (
	"errors"
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

func fetch(method string, url string, params []QueryParameter) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if len(params) > 0 && method == "GET" {
		query := req.URL.Query()

		for _, param := range params {
			query.Add(param.name, param.value)
		}

		req.URL.RawQuery = query.Encode()
	}

	resp, _ := client.Do(req)

	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		err = errors.New(http.StatusText(resp.StatusCode))
	}

	if err != nil {
		return nil, err
	} else {
		payload, _ := ioutil.ReadAll(resp.Body)

		return payload, nil
	}
}

func FetchSearch(extensionName string) []byte {
	params := []QueryParameter{{name: "search", value: extensionName}}
	searchResult, _ := fetch("GET", "https://extensions.gnome.org/extension-query", params)

	return searchResult
}

func FetchExtensionMetadata(extensionId string, systemShellVersion string) ([]byte, error) {
	params := []QueryParameter{{name: "pk", value: extensionId}, {name: "shell_version", value: systemShellVersion}}
	extensionMetadata, err := fetch("GET", "https://extensions.gnome.org/extension-info/", params)

	if err != nil {
		return nil, err
	}

	return extensionMetadata, nil
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
