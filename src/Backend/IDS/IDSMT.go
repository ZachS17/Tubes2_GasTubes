package IDS

import "sync"

func DLSMT(currentPageName string, desiredPageName string, visited []string, solution []string, currentDepth int, desiredDepth int, numArticlesVisited int) ([]string, int) {
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
				var wg sync.WaitGroup
				results := make(chan struct {
					solution           []string
					numArticlesVisited int
				}, len(links))

				for _, link := range links {
					wg.Add(1)
					go func(link string) {
						defer wg.Done()
						// tambahin visited
						numArticlesVisited++
						// dibuat array baru agar terpisah dari parameter solution
						newSolution := append([]string{}, solution...)
						newSolution = append(newSolution, link)
						newSolution, _ = DLSMT(link, desiredPageName, visited, newSolution, currentDepth+1, desiredDepth, numArticlesVisited)
						results <- struct {
							solution           []string
							numArticlesVisited int
						}{newSolution, numArticlesVisited}
					}(link)
				}

				go func() {
					wg.Wait()
					close(results)
				}()

				for result := range results {
					if len(result.solution) > len(solution) && result.solution[len(result.solution)-1] == desiredPageName {
						return result.solution, result.numArticlesVisited
					}
				}
			}
		}
	}

	return solution, numArticlesVisited
}
