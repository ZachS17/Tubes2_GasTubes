package main

import (
	"Backend/BFS"
	"fmt"
	"time"
)

func main() {
	startTitle := "https://en.wikipedia.org/wiki/Iowa"      // contoh start page
	targetTitle := "https://en.wikipedia.org/wiki/Beyblade" // contoh target page

	waktuMulai := time.Now()
	temp, numArticlesVisited, numArticlesChecked := BFS.CallBFS(startTitle, targetTitle)
	waktuAkhir := time.Now()
	waktuEksekusi := waktuAkhir.Sub(waktuMulai)
	fmt.Println(temp)
	fmt.Println(numArticlesVisited)
	fmt.Println(numArticlesChecked)
	fmt.Println(waktuEksekusi)
}
