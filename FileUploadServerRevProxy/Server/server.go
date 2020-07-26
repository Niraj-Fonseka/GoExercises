package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var (
	REDIRECT_URL = "http://localhost:8080"
)

func HandleUploadRequest(w http.ResponseWriter, r *http.Request) {
	url, err := url.Parse(REDIRECT_URL)

	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Parsed URL	:", url)
	proxy := httputil.NewSingleHostReverseProxy(url)

	r.URL.Host = url.Host
	r.URL.Scheme = url.Scheme
	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
	r.Host = url.Host

	proxy.ServeHTTP(w, r)
}

func main() {
	log.Println("Starting server ..")

	http.HandleFunc("/upload", HandleUploadRequest)

	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal(err)
	}
}
