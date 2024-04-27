package IDS

func isInArray(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

func DLS(currentPageName string, desiredPageName string, visited []string, solution []string, currentDepth int, desiredDepth int, numArticlesVisited int) ([]string, int) {
	numArticlesVisited++
	// utamain ketemu
	if currentPageName == desiredPageName {
		return solution, numArticlesVisited
	} else {
		// utamain kedalaman
		if currentDepth == desiredDepth {
			return solution, numArticlesVisited
		} else {
			// utamain sudah dikunjungi
			if isInArray(visited, currentPageName) {
				return solution, numArticlesVisited
			} else {
				// tambah daftar yang sudah dikunjungi
				visited = append(visited, currentPageName)
				// dapat link
				links := GetWikipediaLinks(currentPageName)
				for _, link := range links {
					// tambahin visited
					numArticlesVisited++
					// dibuat array baru agar terpisah dari parameter solution
					newSolution := append([]string{}, solution...)
					newSolution = append(newSolution, link)
					newSolution, _ = DLS(link, desiredPageName, visited, newSolution, currentDepth+1, desiredDepth, numArticlesVisited)
					if len(newSolution) > len(solution) && newSolution[len(newSolution)-1] == desiredPageName {
						return newSolution, numArticlesVisited
					}
				}
			}
		}
	}

	return solution, numArticlesVisited
}

func IDS(initialPageName string, desiredPageName string) ([]string, int, int) {
	var found bool = false
	var currentDepth int = 0
	var result []string
	var numArticlesVisited int
	var visited []string
	// Inisialisasi dan dibuat copy karena akan diubah pada DLS
	copyInitialPageName := initialPageName
	// selama belum ketemu dan kedalamannya belum maksimal
	for !found {
		// // termasuk yang awal
		// numArticlesVisited = 1
		// solution diubah setiap DLS
		solution := []string{copyInitialPageName}
		// dapat hasil
		result, numArticlesVisited = DLSMT(copyInitialPageName, desiredPageName, visited, solution, 0, currentDepth, numArticlesVisited)
		// cek solusi
		if len(result) != 1 {
			found = true
		}
		// tambah kedalaman maksimal
		currentDepth = currentDepth + 1
		// reset untuk iterasi berikutnya karena diubah dalam DLS
		copyInitialPageName = initialPageName
	}
	return result, numArticlesVisited, len(result) - 1
}
