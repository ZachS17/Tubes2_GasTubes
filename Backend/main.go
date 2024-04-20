package main

import (
	"Backend/IDS"
	"fmt"
	"time"
)

func main() {
	var awal string = "Barack Obama"
	var tujuan string = "Michelle Obama"
	mulai := time.Now()
	solution, found := IDS.IDS(awal, tujuan, 1)
	akhir := time.Now()
	var num int = 0
	num = 1
	if !found { // tidak ketemu
		fmt.Println("Tidak ditemui")
	}
	// visited, _ := IDS.GetWikipediaLinks(awal)
	fmt.Println("Solusi: ")
	for _, link := range solution {
		fmt.Println(num, link)
		num = num + 1
	}
	fmt.Println("Banyak link yang ditelusuri:")
	fmt.Println(len(solution))
	fmt.Println("Banyak link untuk sampai:")
	fmt.Println(num)
	waktuEksekusi := akhir.Sub(mulai)
	fmt.Println("Waktu eksekusi", waktuEksekusi)
}
