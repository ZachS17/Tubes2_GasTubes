package main

import (
	"Backend/BFS"
	"fmt"
)

func main() {
	var tujuan string = "Democratic Party"
	solution := BFS.BFS("Barack Obama", tujuan, nil, nil)
	var num int
	num = 0
	for _, link := range solution {
		fmt.Println(link)
		num = num + 1
		if link == tujuan {
			break
		}
	}
	fmt.Println("Banyak link yang ditelusuri:")
	fmt.Println(len(solution))
	fmt.Println("Banyak link untuk sampai:")
	fmt.Println(num)
	// fmt.Println(ketemu)
	// fmt.Println(solution)
}
