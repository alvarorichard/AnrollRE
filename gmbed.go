package main

import (
	"encoding/base64"
	"io"
	"net/http"
	"strings"
)

func embed(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "missing url parameter", http.StatusBadRequest)
		return
	}
	urlDecoded, err := base64.StdEncoding.DecodeString(url)
	if err != nil {
		http.Error(w, "invalid url parameter", http.StatusBadRequest)
		return
	}
	newUrl := strings.TrimSpace(string(urlDecoded))
	req, _ := http.NewRequest(http.MethodGet, newUrl, nil)
	req.Header.Add("connection", "keep-alive")
	req.Header.Add("pragma", "no-cache")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36")
	req.Header.Add("sec-ch-ua", "\"Google Chrome\";v=\"135\", \"Not-A.Brand\";v=\"8\", \"Chromium\";v=\"135\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("accept", "*/*")
	req.Header.Add("origin", "https://www.anroll.net")
	req.Header.Add("sec-fetch-site", "cross-site")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("referer", "https://www.anroll.net/")
	req.Header.Add("accept-language", "en-US,en;q=0.9")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "failed to fetch url", http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()
	_, err = io.Copy(w, res.Body)
	if err != nil {
		http.Error(w, "failed to copy response", http.StatusInternalServerError)
		return
	}
}
