package client

import (
	"net/http"
	"io/ioutil"
)

func FetchSearch(extensionName string) string {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", "https://extensions.gnome.org/extension-query", nil)

	query := request.URL.Query()
	query.Add("search", extensionName)

	request.URL.RawQuery = query.Encode()

	response, _ := client.Do(request)

	payload, _ := ioutil.ReadAll(response.Body)
	body := string(payload)

	return body
}
