package main

import (
	"net/http"
	"net/http/cookiejar"
)

type scraper struct {
	cli *http.Client
	jar http.CookieJar
}

func newScraper() *scraper {
	jar, _ := cookiejar.New(nil)
	cli := &http.Client{
		Jar: jar,
	}
	return &scraper{
		cli: cli,
		jar: jar,
	}
}
