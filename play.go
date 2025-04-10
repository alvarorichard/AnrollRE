package main

import (
	"io"
	"net/http"
)

func (a *Anime) Play(url string) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req, _ := http.NewRequest(http.MethodGet, url, nil)

		req.Header.Set("pragma", "no-cache")
		req.Header.Set("cache-control", "no-cache")
		req.Header.Set("sec-ch-ua-platform", "\"macOS\"")
		req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36")
		req.Header.Set("sec-ch-ua", "\"Google Chrome\";v=\"135\", \"Not-A.Brand\";v=\"8\", \"Chromium\";v=\"135\"")
		req.Header.Set("sec-ch-ua-mobile", "?0")
		req.Header.Set("accept", "*/*")
		req.Header.Set("sec-fetch-site", "same-origin")
		req.Header.Set("sec-fetch-mode", "no-cors")
		req.Header.Set("sec-fetch-dest", "script")
		req.Header.Set("referer", "https://www.anroll.net/e/fWW8QtN82n")

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		defer res.Body.Close()
		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		newBody := replaceBody(string(resBody))
		w.Header().Add("Content-Type", res.Header.Get("Content-Type"))
		w.Write([]byte(newBody))
	})
}
