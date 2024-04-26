package main

import (
	"Backend/BFS"
	"Backend/IDS"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/rs/cors"
)

type Result struct {
	Path               []string      `json:"path"`
	NumArticlesVisited int           `json:"numArticlesVisited"`
	NumArticlesChecked int           `json:"numArticlesChecked"`
	ExecutionTime      time.Duration `json:"executionTime"`
}

func main() {
	// Create a new CORS handler allowing requests from localhost:3000
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})

	// Wrap the default ServeMux with the CORS handler
	handler := corsHandler.Handler(http.DefaultServeMux)

	http.HandleFunc("/wikirace", algorithmHandler)
	fmt.Println("Server listening on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", handler))
}

func algorithmHandler(w http.ResponseWriter, r *http.Request) {
	algorithmType := r.URL.Query().Get("algorithm")
	if algorithmType == "BFS" {
		BFSHandler(w, r)
	} else if algorithmType == "IDS" {
		IDSHandler(w, r)
	} else {
		http.Error(w, "Algoritma tidak valid!", http.StatusBadRequest)
	}
}

func BFSHandler(w http.ResponseWriter, r *http.Request) {
	initialPage := r.URL.Query().Get("initial")
	destinationPage := r.URL.Query().Get("destination")

	startTime := time.Now()
	path, articlesVisited, articlesChecked := BFS.CallBFS(initialPage, destinationPage)
	finishedTime := time.Now()
	executionTime := finishedTime.Sub(startTime)

	formatExecutionTime := time.Duration(executionTime.Nanoseconds() / int64(time.Millisecond))

	solution := Result{
		Path:               path,
		NumArticlesVisited: articlesVisited,
		NumArticlesChecked: articlesChecked,
		ExecutionTime:      formatExecutionTime,
	}

	jsonResponse, err := json.Marshal(solution)
	if err != nil {
		http.Error(w, "Unable to marshal JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func IDSHandler(w http.ResponseWriter, r *http.Request) {
	initialPage := r.URL.Query().Get("initial")
	destinationPage := r.URL.Query().Get("destination")

	startTime := time.Now()
	path, articlesVisited, articlesChecked := IDS.IDS(initialPage, destinationPage)
	finishedTime := time.Now()
	executionTime := finishedTime.Sub(startTime)

	formatExecutionTime := time.Duration(executionTime.Nanoseconds() / int64(time.Millisecond))

	solution := Result{
		Path:               path,
		NumArticlesVisited: articlesVisited,
		NumArticlesChecked: articlesChecked,
		ExecutionTime:      formatExecutionTime,
	}

	jsonResponse, err := json.Marshal(solution)
	if err != nil {
		http.Error(w, "Unable to marshal JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
