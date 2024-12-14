package main

import (
	"net/http"

	"github.com/Jhone-Paul/UrlShortener/urlshortener"
)

func main() {
	// These are the URLS to be mapped
	pathsToUrls := map[string]string{
		"/git": "https://github.com/Jhone-Paul",
		"/go":  "https://golang.org",
	}

	fallback := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte("404 - Not Found"))
	})

	handler := urlshortener.MapHandler(pathsToUrls, fallback)
	//you can select your own port :0
	port := ":8080"
	println("Starting the server on", port)
	err := http.ListenAndServe(port, handler)
	if err != nil {
		panic(err)
	}
}
