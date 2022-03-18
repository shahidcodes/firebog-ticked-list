package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func GetJson(url string, target interface{}) {
	r, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	jsonBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(jsonBytes, &target)
}

func GetText(url string) string {
	r, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()

	bodyBytes, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	return string(bodyBytes)
}
