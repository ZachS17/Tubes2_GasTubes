package BFS

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
	solution, articlesVisited := BFS(possibleSolutions, visited, desiredPageName, articlesVisited)
	return solution, articlesVisited, len(solution)
}
