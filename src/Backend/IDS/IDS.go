package IDS

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

// kalau g boleh cache

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

// 	var awalan = []string{
// 		"/wiki/Draft:",
// 		"/wiki/Module:",
// 		"/wiki/MediaWiki:",
// 		"/wiki/Index:",
// 		"/wiki/Education_Program:",
// 		"/wiki/TimedText:",
// 		"/wiki/Gadget:",
// 		"/wiki/Gadget_Definition:",
// 		"/wiki/Main_Page",
// 		"/wiki/Main_Page:",
// 		"/wiki/Special:",
// 		"/wiki/Talk:",
// 		"/wiki/User:",
// 		"/wiki/Portal:",
// 		"/wiki/Wikipedia:",
// 		"/wiki/File:",
// 		"/wiki/Category:",
// 		"/wiki/Help:",
// 		"/wiki/Template:",
// 	}

// 	links := []string{}
// 	// cari link
// 	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
// 		link, _ := s.Attr("href")
// 		// wikipedia page
// 		if strings.HasPrefix(link, "/wiki/") {
// 			skip := false
// 			// awalan nggak kepakai (diskip)
// 			for _, prefix := range awalan {
// 				if strings.HasPrefix(link, prefix) {
// 					skip = true
// 					break
// 				}
// 			}
// 			if !skip {
// 				links = append(links, "https://en.wikipedia.org"+link)
// 			}
// 		}
// 	})
// 	return links
// }

var cache map[string][]string
var cacheMutex sync.Mutex

func init() {
	cache = make(map[string][]string)
	loadCache()
}

func loadCache() {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()
	file, err := os.Open("/app/Backend/cache.gob")
	if err != nil {
		if os.IsNotExist(err) {
			// Cache file doesn't exist, create it
			file, err = os.Create("/app/Backend/cache.gob")
			if err != nil {
				log.Fatal("Error creating cache file:", err)
			}
			defer file.Close()
			return
		}
		log.Fatal("Error opening cache file:", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal("Error getting file information:", err)
	}

	if fileInfo.Size() == 0 {
		// Cache file is empty, no need to load anything
		return
	}

	decoder := gob.NewDecoder(file)
	if err := decoder.Decode(&cache); err != nil {
		log.Fatal("Error decoding cache:", err)
	}
}

func saveCache() {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()
	file, err := os.OpenFile("/app/Backend/cache.gob", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Error opening cache file:", err)
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	if err := encoder.Encode(cache); err != nil {
		log.Fatal("Error encoding cache:", err)
	}
}

func GetWikipediaLinks(URL string) []string {
	// cek yang mau diekspan ada di cache atau tidak
	if links, ok := cache[URL]; ok {
		return links
	}

	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	links := []string{}

	if resp.Request.URL.String() != URL {
		URL = resp.Request.URL.String()
		links = append(links, URL)
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

	// cari link
	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		// wikipedia page
		if strings.HasPrefix(link, "/wiki/") {
			hasPrefix := false
			// awalan nggak kepakai (diskip)
			for _, prefix := range awalan {
				if strings.HasPrefix(link, prefix) {
					hasPrefix = true
					break
				}
			}
			if !hasPrefix {
				links = append(links, "https://en.wikipedia.org"+link)
			}
		}
	})

	// cache
	cache[URL] = links
	// simpan dalam data
	saveCache()

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
		result, numArticlesVisited = DLS(copyInitialPageName, desiredPageName, visited, solution, 0, currentDepth, numArticlesVisited)
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
