package client

import (
	"os"
	"io"
	"fmt"
)

const TEMPORAL_DOWNLOADS_FOLDER = "extension.zip"

func DownloadExtension(extensionId string) {
	// TODO: first call to https://extensions.gnome.org/extension-info/?pk=79&shell_version=40 and then extract the metadata to call
	// download-extension/hide-dash@zacbarton.com.shell-extension.zip?version_tag=1993
	extensionMetadata := fetchExtensionMetadata(extensionId)
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

	extensionZip, err := os.Create(TEMPORAL_DOWNLOADS_FOLDER)

	if err != nil {
		fmt.Println(err)
	}

	io.Copy(extensionZip, response.Body)

	defer extensionZip.Close()
}

func fetchExtensionMetadata(extensionId string) string {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", "https://extensions.gnome.org/extension-info", nil)
	query := request.URL.Query()

	query.Add("pk", extensionId)
	query.Add("shell_version", getSystemShellVersion())
	request.URL.RawQuery = query.Encode()
}

func getSystemShellVersion() string {
	return "40"
}
