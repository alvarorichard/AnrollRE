package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/tidwall/gjson"
)

type Episode struct {
	IDSeriesEpisodios int64     `json:"id_series_episodios"`
	SEPgad            int64     `json:"se_pgad"`
	IDSerie           int64     `json:"id_serie"`
	PremiereLastEp    int64     `json:"premiere_last_ep"`
	NEpisodio         string    `json:"n_episodio"`
	TituloEpisodio    string    `json:"titulo_episodio"`
	SinopseEpisodio   string    `json:"sinopse_episodio"`
	Link              string    `json:"link"`
	VStream           any       `json:"v_stream"`
	Aviso             string    `json:"aviso"`
	GenerateID        string    `json:"generate_id"`
	DataRegistro      time.Time `json:"data_registro"`
	Anime             Anime     `json:"anime"`
}

const baseEpisodes = `https://apiv3-prd.anroll.net`

func (a *Anime) GetEpisodes() []Episode {
	u := fmt.Sprintf("%s/animes/%d/episodes?page=1&order=desc", baseEpisodes, a.id)
	req, _ := http.NewRequest(http.MethodGet, u, nil)
	for key, value := range commonHeaders {
		req.Header.Set(key, value)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	resBody, _ := io.ReadAll(res.Body)
	println(string(resBody))
	data := gjson.GetBytes(resBody, "data").Array()
	var episodes []Episode
	for _, item := range data {
		var ep Episode
		err := json.Unmarshal([]byte(item.Raw), &ep)
		if err != nil {
			fmt.Println("Error unmarshalling episode:", err)
			continue
		}
		episodes = append(episodes, ep)
	}
	return episodes
}
