package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type response struct {
	Data string `json:"data"`
}

func base(w http.ResponseWriter, req *http.Request) {
	log.Println("Processing request...")
	resp := response{Data: "simple api app"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", base)

	port := ":8080"
	if p, ok := os.LookupEnv("PORT"); ok {
		port = ":" + p
	}

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalln(err)
	}
}
