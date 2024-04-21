package main

import (
	"Backend/BFS"
	"Backend/IDS"
	"fmt"
	"time"
)

func main() {
	startTitle := "https://en.wikipedia.org/wiki/Barack_Obama"    // contoh start page
	targetTitle := "https://en.wikipedia.org/wiki/Michelle_Obama" // contoh target page
	waktuMulai := time.Now()
	temp, found := IDS.IDS(startTitle, targetTitle, 3)
	waktuAkhir := time.Now()
	waktuEksekusi := waktuAkhir.Sub(waktuMulai)
	fmt.Println(found)
	fmt.Println(temp)
	fmt.Println(waktuEksekusi)
	waktuMulai = time.Now()
	temp, found = BFS.CallBFS(startTitle, targetTitle)
	waktuAkhir = time.Now()
	waktuEksekusi = waktuAkhir.Sub(waktuMulai)
	fmt.Println(found)
	fmt.Println(temp)
	fmt.Println(waktuEksekusi)
}
