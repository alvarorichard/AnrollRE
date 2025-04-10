package main

import "net/http"

func main() {
	// s := newScraper()
	// s.search("Black Clover")
	a := Anime{
		slug: "black-clover",
	}
	url := a.getUrl("170")
	mux := http.NewServeMux()
	mux.Handle("GET /play", a.Play(url))
	mux.HandleFunc("/embed", embed)
	s := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}

}
