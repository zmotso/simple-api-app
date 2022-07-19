package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type response struct {
	Data string `json:"data"`
}

func base(w http.ResponseWriter, req *http.Request) {
	resp := response{Data: "simple api app"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		fmt.Println(err)
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
		fmt.Println(err)
		os.Exit(1)
	}
}
