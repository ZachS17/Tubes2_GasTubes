package IDS

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery" // Import the HTML parser package
)

func GetLinks(URL string) []string {
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

	links := []string{}
	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		if strings.HasPrefix(link, "/wiki/") {
			links = append(links, "https://en.wikipedia.org"+link)
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

func DLS(currentPageName string, desiredPageName string, visited []string, solution []string, currentDepth int, desiredDepth int) []string {
	// utamain ketemu
	if currentPageName == desiredPageName {
		return solution
	} else {
		// utamain kedalaman
		if currentDepth == desiredDepth {
			return solution
		} else {
			// utamain sudah dikunjungi
			if isInArray(visited, currentPageName) {
				return solution
			} else {
				// tambah daftar yang sudah dikunjungi
				visited = append(visited, currentPageName)
				// dapat link
				links := GetLinks(currentPageName)
				for _, link := range links {
					// dibuat array baru agar terpisah dari parameter solution
					newSolution := append([]string{}, solution...)
					newSolution = append(newSolution, link)
					newSolution = DLS(link, desiredPageName, visited, newSolution, currentDepth+1, desiredDepth)
					if len(newSolution) > len(solution) && newSolution[len(newSolution)-1] == desiredPageName {
						return newSolution
					}
				}
			}
		}
	}

	return solution
}

func IDS(initialPageName string, desiredPageName string, desiredDepth int) ([]string, bool) {
	var found bool = false
	var currentDepth int = 0
	var hasil []string
	// Inisialisasi dan dibuat copy karena akan diubah pada DLS
	copyInitialPageName := initialPageName
	// selama belum ketemu dan kedalamannya belum maksimal
	for !found && currentDepth <= desiredDepth {
		// visited dari awal
		var visited []string
		// solution diubah setiap DLS
		solution := []string{copyInitialPageName}
		// dapat hasil
		hasil = DLS(copyInitialPageName, desiredPageName, visited, solution, 0, currentDepth)
		// cek solusi
		if isInArray(hasil, desiredPageName) {
			found = true
		}
		// tambah kedalaman maksimal
		currentDepth = currentDepth + 1
		// reset untuk iterasi berikutnya karena diubah dalam DLS
		copyInitialPageName = initialPageName
	}
	return hasil, found
}
