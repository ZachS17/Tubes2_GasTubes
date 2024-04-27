package BFS

import (
	"math"
	"sync"
)

func BFSMT(possibleSolutions [][]string, visited []string, desiredPageName string, numArticlesVisited int) ([]string, int) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	found := make(chan struct {
		path []string
		num  int
	}, len(possibleSolutions))

	for _, stringArray := range possibleSolutions {
		wg.Add(1)
		go func(stringArray []string, visited []string, desiredPageName string, numArticlesVisited int) {
			defer wg.Done()

			mu.Lock()
			numArticlesVisited++
			lastURL := lastElement(stringArray)
			if isInArray(visited, lastURL) {
				mu.Unlock()
				return
			}
			visited = append(visited, lastURL)
			mu.Unlock()

			if lastURL == desiredPageName {
				found <- struct {
					path []string
					num  int
				}{stringArray, numArticlesVisited}
				return
			}

			linkArray := GetWikipediaLinks(lastURL)
			for _, link := range linkArray {
				copyStringArray := append([]string{}, stringArray...)
				copyStringArray = append(copyStringArray, link)

				wg.Add(1)
				go func(copyStringArray []string, visited []string, desiredPageName string, numArticlesVisited int) {
					defer wg.Done()
					BFSMT([][]string{copyStringArray}, visited, desiredPageName, numArticlesVisited)
				}(copyStringArray, visited, desiredPageName, numArticlesVisited)
			}
		}(stringArray, visited, desiredPageName, numArticlesVisited)
	}

	go func() {
		wg.Wait()
		close(found)
	}()

	var shortestPath []string
	shortestNum := math.MaxInt64
	for result := range found {
		if len(result.path) > 0 && result.num < shortestNum {
			shortestPath = result.path
			shortestNum = result.num
		}
	}

	return shortestPath, shortestNum
}
