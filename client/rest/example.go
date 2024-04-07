package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func Get() error {
	var (
		payload = make(map[string]any)
		method  = "GET"
		params  = url.Values{}
	)

	uri, err := url.Parse("https://webhook.site/f31a468e-731b-467d-bd83-5ec96ef2a189")
	if err != nil {
		return err
	}

	params.Add("foo", "bar")
	params.Add("age", "30")

	uri.RawQuery = params.Encode()

	request, err := http.NewRequest(method, uri.String(), nil)
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer 123456")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(responseData, &payload)
	if err != nil {
		return err
	}

	fmt.Println(payload)

	return nil
}

func Post() error {
	var (
		payload = make(map[string]any)
		method  = "POST"
		uri     = "https://355ea40a7eb949ee8cbff740950d3359.api.mockbin.io/"
	)

	request, err := http.NewRequest(method, uri, nil)
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer 123456")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(responseData, &payload)
	if err != nil {
		return err
	}

	fmt.Println(payload)

	return nil
}
