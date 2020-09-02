package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// As for now this wont work
// 'TMDB' refuses connection from proxy?

func main() {
	log.Println("Running reverse proxy to https://api.themoviedb.org")
	remoteURL, _ := url.Parse("https://api.themoviedb.org")
	proxy := httputil.NewSingleHostReverseProxy(remoteURL)
	http.ListenAndServe(":8899", proxy)
}
