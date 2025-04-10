package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/tidwall/gjson"
)

const baseUrl = "https://api-search.anroll.net/"

var commonHeaders = map[string]string{
	"Content-Type": "application/json",
	"User-Agent":   "MyApp/1.0",
	"Accept":       "application/json",
}

func (s *scraper) search(term string) {
	url := fmt.Sprintf("%s/data?q=%s", baseUrl, term)
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	for key, value := range commonHeaders {
		req.Header.Set(key, value)
	}
	res, err := s.cli.Do(req)
	if err != nil {
		panic(err)
	}
	resBody, _ := io.ReadAll(res.Body)
	data := gjson.GetBytes(resBody, "data").Array()
	for _, item := range data {
		m := item.Map()
		for key, value := range m {
			fmt.Printf("%s: %s\n", key, value.String())

		}
	}

}
