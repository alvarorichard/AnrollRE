package main

import (
	"fmt"
	"net/http"
)

func main() {
	s := newScraper()
	animes := s.search("black")
	for i, anime := range animes {
		fmt.Printf("%d - Title: %s\n", i, anime.Title)
	}
	var choice int
	fmt.Println("Enter the number of the anime you want to watch:")
	fmt.Scan(&choice)
	if choice < 0 || choice >= len(animes) {
		fmt.Println("Invalid choice")
		return
	}
	anime := animes[choice]

	a := Anime{
		slug: anime.Slug,
		id:   anime.ID,
	}
	episodes := a.GetEpisodes()
	for _, ep := range episodes {
		fmt.Printf("Episode %s - %s\n", ep.NEpisodio, ep.TituloEpisodio)
	}

	var epChoice string
	fmt.Println("Enter the number of the episode you want to watch:")
	fmt.Scan(&epChoice)

	if epChoice == "" {
		fmt.Println("Invalid choice")
		return
	}

	url := a.getUrl(epChoice)
	mux := http.NewServeMux()
	mux.Handle("GET /play", a.Play(url))
	mux.HandleFunc("/embed", embed)
	se := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}
	fmt.Println("Server started at http://localhost:8000")

	if err := se.ListenAndServe(); err != nil {
		panic(err)
	}
}
