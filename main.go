package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Kubesimplify struct {
	Website string `json:"website"`
	Twitter string `json:"twitter"`
	Sponsor string `json:"sponsor"`
	Founder string `json:"founder"`
}

func main() {
	addr := ":6400"
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", func(resWriter http.ResponseWriter, request *http.Request) {
		data := struct{ Status string }{Status: "OK"}
		resWriter.Header().Set("Content-Type", "application/json")
		resWriter.WriteHeader(http.StatusOK)
		json.NewEncoder(resWriter).Encode(data)
	})

	mux.HandleFunc("GET /health", func(resWriter http.ResponseWriter, request *http.Request) {
		kubesimplify := Kubesimplify{
			Website: "kubesimplify.com",
			Twitter: "@kubesimplify",
			Sponsor: "kubesimplify.com/sponsor",
			Founder: "Saiyam Pathak",
		}
		resWriter.Header().Set("Content-Type", "application/json")
		resWriter.WriteHeader(http.StatusOK)
		json.NewEncoder(resWriter).Encode(kubesimplify)
	})

	fmt.Printf("Starting server on %s", addr)
	err := http.ListenAndServe(addr, mux)
	log.Fatal(err)
}
