package BFS

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type BFSRes struct {
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

func lastElement(array []string) string {
	return array[len(array)-1]
}

// rekursif
// semua elemen berupa arrayString yang terdiri dari semua urutan link
// perbandingan langsung dilakukan pada elemen terakhir setiap stringarray pada possible solution
func BFS(possibleSolutions [][]string, visited []string, desiredPageName string, numArticlesVisited int, numArticlesChecked int) ([]string, int, int) {
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
			// tambah checked (tidak ada di visited -> pasti harus dicheck)
			numArticlesChecked++
			// ketemu
			if lastElement(stringArray) == desiredPageName {
				return stringArray, numArticlesVisited, numArticlesChecked
			} else {
				// daftar link dari elemen terakhir (yang akan diexpand)
				linkArray := GetWikipediaLinks(lastElement(stringArray))
				// copy dari yang mau diexpand
				copyStringArray := stringArray
				// untuk setiap link ditambahin
				for _, link := range linkArray {
					// tambahin link baru
					copyStringArray = append(copyStringArray, link)
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
	return BFS(tempPossibleSolutions, visited, desiredPageName, numArticlesVisited, numArticlesChecked)
}

func CallBFS(initialPageName string, desiredPageName string) ([]string, int, int) {
	possibleSolutions := [][]string{{initialPageName}}
	var visited []string
	var articlesVisited int = 1
	var articlesChecked int
	solution, articlesVisited, articlesChecked := BFS(possibleSolutions, visited, desiredPageName, articlesVisited, articlesChecked)
	return solution, articlesVisited, articlesChecked
}

func BFSHandler(w http.ResponseWriter, r *http.Request) {
	initialPage := r.URL.Query().Get("initial")
	destinationPage := r.URL.Query().Get("destination")

	startTime := time.Now()
	path, articlesVisited, articlesChecked := CallBFS(initialPage, destinationPage)
	finishedTime := time.Now()
	executionTime := finishedTime.Sub(startTime)

	formatExecutionTime := time.Duration(executionTime.Nanoseconds() / int64(time.Millisecond))

	solution := BFSRes{
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
