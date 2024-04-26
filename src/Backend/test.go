package main

// import (
// 	"Backend/BFS"
// 	"fmt"
// 	"time"
// )

// func main() {
// 	startTitle := "https://en.wikipedia.org/wiki/Iowa"     // contoh start page
// 	targetTitle := "https://en.wikipedia.org/wiki/Bandung" // contoh target page

// 	waktuMulai := time.Now()
// 	temp, numArticlesVisited, numArticlesChecked := BFS.CallBFS(startTitle, targetTitle)
// 	waktuAkhir := time.Now()
// 	waktuEksekusi := waktuAkhir.Sub(waktuMulai)
// 	fmt.Println("BFS:")
// 	fmt.Println(temp)
// 	fmt.Println(numArticlesVisited)
// 	fmt.Println(numArticlesChecked)
// 	fmt.Println(waktuEksekusi)

// 	waktuMulai = time.Now()
// 	temp, numArticlesVisited, numArticlesChecked = IDS.IDS(startTitle, targetTitle)
// 	waktuAkhir = time.Now()
// 	waktuEksekusi = waktuAkhir.Sub(waktuMulai)
// 	fmt.Println("IDS:")
// 	fmt.Println(temp)
// 	fmt.Println(numArticlesVisited)
// 	fmt.Println(numArticlesChecked)
// 	fmt.Println(waktuEksekusi)

// 	waktuMulai = time.Now()
// 	temp, numArticlesVisited, numArticlesChecked = BFS.BFS1(startTitle, targetTitle)
// 	waktuAkhir = time.Now()
// 	waktuEksekusi = waktuAkhir.Sub(waktuMulai)
// 	fmt.Println("BFS1:")
// 	fmt.Println(temp)
// 	fmt.Println(numArticlesVisited)
// 	fmt.Println(numArticlesChecked)
// 	fmt.Println(waktuEksekusi)
// }
