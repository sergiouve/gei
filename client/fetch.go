package main

import (
	"net/http"
)

type QueryParameter struct {
	name  string
	value string
}

func Fetch(method string, url string, params [QueryParameter]string) string {
	client := &http.Client{}
	request, _ := http.NewRequest(method, url, nil)

	if (len(params) > 0) {
		query := request.URL.Query()

		for _, param := range params {
			query.Add(param.name, param.value)
		}

		request.URL.RawQuery = query.Encode()
	}

	response, _ := client.Do(request)
	payload, _ := ioutil.ReadAll(response.Body)
	body := string(payload)

	return body
}
