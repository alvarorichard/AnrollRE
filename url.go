package main

import "fmt"

const cdnBaseUrl = `https://cdn-zenitsu-2-gamabunta.b-cdn.net/cf/hls/animes`

func (a *Anime) getUrl(numEp string) string {
	return fmt.Sprintf("%s/%s/%s.mp4/media-1/stream.m3u8",
		cdnBaseUrl, a.slug, numEp)
}
