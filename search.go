package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/tidwall/gjson"
)

const baseUrl = "https://api-search.anroll.net"

var commonHeaders = map[string]string{
	"Content-Type": "application/json",
	"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36",
	"Accept":       "application/json",
}

type SearchTermBody struct {
	Type         string `json:"type"`
	ID           int64  `json:"id"`
	Title        string `json:"title"`
	Slug         string `json:"slug"`
	Year         any    `json:"year"`
	Censorship   int64  `json:"censorship"`
	Synopsis     string `json:"synopsis"`
	TotalEps     int64  `json:"total_eps"`
	GenID        string `json:"gen_id"`
	FriendlyPath string `json:"friendly_path"`
	GenericPath  string `json:"generic_path"`
}

func (s *scraper) search(term string) []SearchTermBody {
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
	var items []SearchTermBody
	for _, item := range data {
		var i SearchTermBody
		err := json.Unmarshal([]byte(item.Raw), &i)
		if err != nil {
			fmt.Println("Error unmarshalling item:", err)
			continue
		}
		items = append(items, i)
	}
	return items
}
