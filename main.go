package main

import (
	"encoding/json"
	"io"
	"net/http"

	"log"
	"os"

	"github.com/Jhone-Paul/UrlShortener/urlshortener"
)

func main() {
	pathsToUrls, err := loadMappings("mappings.json")
	if err != nil {
		log.Fatalf("Failed to load mappings: %v", err)
	}

	fallback := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte("404 - Not Found"))
	})

	handler := urlshortener.MapHandler(pathsToUrls, fallback)

	port := ":8080"
	println("Starting the server on", port)
	err = http.ListenAndServe(port, handler)
	if err != nil {
		panic(err)
	}
}

// loadMappings reads a JSON file and parses it into a map
func loadMappings(filename string) (map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var pathsToUrls map[string]string
	err = json.Unmarshal(byteValue, &pathsToUrls)
	if err != nil {
		return nil, err
	}

	return pathsToUrls, nil
}
