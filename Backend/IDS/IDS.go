package IDS

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

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

	var awalan = []string{
		"/wiki/Draft:",
		"/wiki/Module:",
		"/wiki/MediaWiki:",
		"/wiki/Index:",
		"/wiki/Education_Program:",
		"/wiki/TimedText:",
		"/wiki/Gadget:",
		"/wiki/Gadget_Definition:",
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
	// cari link
	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		// wikipedia page
		if strings.HasPrefix(link, "/wiki/") {
			skip := false
			// awalan nggak kepakai (diskip)
			for _, prefix := range awalan {
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
