package client

import (
	"net/http"
	"io/ioutil"
)

func FetchExtensionMetadata(extensionId string) string {
	client := &http.Client{}

	request, _ := http.NewRequest("GET", "https://extensions.gnome.org/extension-info", nil)

	query := request.URL.Query()
	query.Add("pk", extensionId)

	request.URL.RawQuery = query.Encode()

	response, _ := client.Do(request)

	payload, _ := ioutil.ReadAll(response.Body)
	body := string(payload)

	return body
}
