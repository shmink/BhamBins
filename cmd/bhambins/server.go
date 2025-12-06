package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/shmink/BhamBins/internal/bins"
)

func runHTTPServer(port string) {
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	mux.HandleFunc("/collections", func(w http.ResponseWriter, r *http.Request) {
		postcode := r.URL.Query().Get("postcode")
		uprn := r.URL.Query().Get("uprn")

		if postcode == "" || uprn == "" {
			http.Error(w, "missing postcode or uprn", http.StatusBadRequest)
			return
		}

		collections, err := bins.Fetch(postcode, uprn, false)
		if err != nil {
			log.Printf("fetch error: %v", err)
			http.Error(w, "failed to fetch collections", http.StatusBadGateway)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(bins.Collections{Bins: collections}); err != nil {
			log.Printf("encode error: %v", err)
		}
	})

	addr := ":" + port
	log.Printf("starting HTTP server on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
