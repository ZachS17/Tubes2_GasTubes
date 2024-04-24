package main

import (
	"Backend/BFS"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
)

// func main() {
// 	// startTitle := "https://en.wikipedia.org/wiki/Barack_Obama"        // contoh start page
// 	// targetTitle := "https://en.wikipedia.org/wiki/Wikipedia:Contents" // contoh target page
// 	// waktuMulai := time.Now()
// 	// temp, numArticlesVisited, numArticlesChecked := IDS.IDS(startTitle, targetTitle)
// 	// waktuAkhir := time.Now()
// 	// waktuEksekusi := waktuAkhir.Sub(waktuMulai)
// 	// fmt.Println(temp)
// 	// fmt.Println(numArticlesVisited)
// 	// fmt.Println(numArticlesChecked)
// 	// fmt.Println(waktuEksekusi)
// 	// waktuMulai = time.Now()
// 	// temp, numArticlesVisited, numArticlesChecked = BFS.CallBFS(startTitle, targetTitle)
// 	// waktuAkhir = time.Now()
// 	// waktuEksekusi = waktuAkhir.Sub(waktuMulai)
// 	// fmt.Println(temp)
// 	// fmt.Println(numArticlesVisited)
// 	// fmt.Println(numArticlesChecked)
// 	// fmt.Println(waktuEksekusi)
// 	// link := BFS.GetWikipediaLinks(startTitle)
// 	// fmt.Println(link[0])
// 	// i := 0
// 	// for {
// 	// 	fmt.Println(link[i])
// 	// 	if i == 10 {
// 	// 		break
// 	// 	}
// 	// 	i++
// 	// }
// 	// fmt.Println(link)

// }

func main() {
	// Create a new CORS handler allowing requests from localhost:3000
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})

	// Wrap the default ServeMux with the CORS handler
	handler := corsHandler.Handler(http.DefaultServeMux)

	http.HandleFunc("/shortestpath", BFS.BFSHandler)
	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
