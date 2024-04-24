package IDS

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery" // Import the HTML parser package
)

type IDSRes struct {
	Path               []string      `json:"path"`
	NumArticlesVisited int           `json:"numArticlesVisited"`
	NumArticlesChecked int           `json:"numArticlesChecked"`
	ExecutionTime      time.Duration `json:"executionTime"`
}

// func GetWikipediaLinks(URL string) []string {
// 	resp, err := http.Get(URL)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != 200 {
// 		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
// 	}

// 	doc, err := goquery.NewDocumentFromReader(resp.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	links := []string{}
// 	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
// 		link, _ := s.Attr("href")
// 		if strings.HasPrefix(link, "/wiki/") {
// 			links = append(links, "https://en.wikipedia.org"+link)
// 		}
// 	})
// 	return links
// }

func GetWikipediaLinks(URL string) []string {
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var prefixes = []string{
		"/wiki/Main_Page",
		"/wiki/Main_Page:",
		"/wiki/Special:",
		"/wiki/Talk:",
		"/wiki/User:",
		"/wiki/Portal:",
		"/wiki/Wikipedia:",
		"/wiki/File:",
		"/wiki/Category:",
		"/wiki/Help:",
		"/wiki/Template:",
	}

	links := []string{}
	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		if strings.HasPrefix(link, "/wiki/") {
			skip := false
			for _, prefix := range prefixes {
				if strings.HasPrefix(link, prefix) {
					skip = true
					break
				}
			}
			if !skip {
				links = append(links, "https://en.wikipedia.org"+link)
			}
		}
	})
	return links
}

func isInArray(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

func DLS(currentPageName string, desiredPageName string, visited []string, solution []string, currentDepth int, desiredDepth int, numArticlesVisited int, numArticlesChecked int) ([]string, int, int) {
	// utamain ketemu
	if currentPageName == desiredPageName {
		return solution, numArticlesVisited, numArticlesChecked
	} else {
		// utamain kedalaman
		if currentDepth == desiredDepth {
			return solution, numArticlesVisited, numArticlesChecked
		} else {
			// utamain sudah dikunjungi
			if isInArray(visited, currentPageName) {
				return solution, numArticlesVisited, numArticlesChecked
			} else {
				// tambah daftar yang sudah dikunjungi
				visited = append(visited, currentPageName)
				// dapat link
				links := GetWikipediaLinks(currentPageName)
				// tambah visited dan check untuk yang mau diexpand
				numArticlesVisited++
				numArticlesChecked++
				for _, link := range links {
					// tambahin visited
					numArticlesVisited++
					// tambahin checked
					if !isInArray(visited, link) {
						numArticlesChecked++
					}
					// dibuat array baru agar terpisah dari parameter solution
					newSolution := append([]string{}, solution...)
					newSolution = append(newSolution, link)
					newSolution, _, _ = DLS(link, desiredPageName, visited, newSolution, currentDepth+1, desiredDepth, numArticlesVisited, numArticlesChecked)
					if len(newSolution) > len(solution) && newSolution[len(newSolution)-1] == desiredPageName {
						return newSolution, numArticlesVisited, numArticlesChecked
					}
				}
			}
		}
	}

	return solution, numArticlesVisited, numArticlesChecked
}

func IDS(initialPageName string, desiredPageName string) ([]string, int, int) {
	var found bool = false
	var currentDepth int = 0
	var result []string
	var numArticlesVisited int
	var numArticlesChecked int
	var visited []string
	// Inisialisasi dan dibuat copy karena akan diubah pada DLS
	copyInitialPageName := initialPageName
	// selama belum ketemu dan kedalamannya belum maksimal
	for !found {
		// termasuk yang awal
		numArticlesVisited = 1
		numArticlesChecked = 0
		// solution diubah setiap DLS
		solution := []string{copyInitialPageName}
		// dapat hasil
		result, numArticlesVisited, numArticlesChecked = DLS(copyInitialPageName, desiredPageName, visited, solution, 0, currentDepth, numArticlesVisited, numArticlesChecked)
		// cek solusi
		if len(result) != 1 {
			found = true
		}
		// tambah kedalaman maksimal
		currentDepth = currentDepth + 1
		// reset untuk iterasi berikutnya karena diubah dalam DLS
		copyInitialPageName = initialPageName
	}
	return result, numArticlesVisited, numArticlesChecked
}

func IDSHandler(w http.ResponseWriter, r *http.Request) {
	initialPage := r.URL.Query().Get("start")
	destinationPage := r.URL.Query().Get("target")

	startTime := time.Now()
	path, articlesVisited, articlesChecked := IDS(initialPage, destinationPage)
	finishedTime := time.Now()
	executionTime := finishedTime.Sub(startTime)

	formatExecutionTime := time.Duration(executionTime.Nanoseconds() / int64(time.Millisecond))

	solution := IDSRes{
		Path:               path,
		NumArticlesVisited: articlesVisited,
		NumArticlesChecked: articlesChecked,
		ExecutionTime:      formatExecutionTime,
	}

	jsonResponse, err := json.Marshal(solution)
	if err != nil {
		http.Error(w, "Unable to marshal JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
