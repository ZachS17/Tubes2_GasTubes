package BFS

func isInArray(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

func lastElement(array []string) string {
	return array[len(array)-1]
}

// rekursif
// semua elemen berupa arrayString yang terdiri dari semua urutan link
// perbandingan langsung dilakukan pada elemen terakhir setiap stringarray pada possible solution
func BFS(possibleSolutions [][]string, visited []string, desiredPageName string, numArticlesVisited int) ([]string, int) {
	// penyimpanan possibleSolution dengan level yang baru
	var tempPossibleSolutions [][]string
	// penelusuran untuk setiap array of string pada possible solutions
	for _, stringArray := range possibleSolutions {
		// tambah visited
		numArticlesVisited++
		// sudah dikunjungi (elemen terakhir dari setiap array of string) -> iterasi selanjutnya
		if isInArray(visited, lastElement(stringArray)) {
			continue
		} else {
			// ketemu
			if lastElement(stringArray) == desiredPageName {
				return stringArray, numArticlesVisited
			} else {
				// daftar link dari elemen terakhir (yang akan diexpand)
				linkArray := GetWikipediaLinks(lastElement(stringArray))
				// copy dari yang mau diexpand
				copyStringArray := stringArray
				// untuk setiap link ditambahin
				for _, link := range linkArray {
					// tambah visited
					numArticlesVisited++
					// tambahin link baru
					copyStringArray = append(copyStringArray, link)
					// pengecekan juga
					if lastElement(copyStringArray) == desiredPageName {
						return copyStringArray, numArticlesVisited
					}
					// dan dimasukkan dalam semua array sesuai kebutuhan
					tempPossibleSolutions = append(tempPossibleSolutions, copyStringArray)
					// pengembalian nilai jadi awal
					copyStringArray = stringArray
				}
				visited = append(visited, lastElement(stringArray))
			}
		}
	}
	// pemanggilan rekursi
	return BFS(tempPossibleSolutions, visited, desiredPageName, numArticlesVisited)
}

func CallBFS(initialPageName string, desiredPageName string) ([]string, int, int) {
	possibleSolutions := [][]string{{initialPageName}}
	var visited []string
	var articlesVisited int
	solution, articlesVisited := BFSMT(possibleSolutions, visited, desiredPageName, articlesVisited)
	return solution, articlesVisited, len(solution) - 1
}
