package BFS

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
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

func lastElement(array []string) string {
	return array[len(array)-1]
}

// rekursif
// semua elemen berupa arrayString yang terdiri dari semua urutan link
// perbandingan langsung dilakukan pada elemen terakhir setiap stringarray pada possible solution
func BFS(possibleSolutions [][]string, visited []string, desiredPageName string) []string {
	// penyimpanan possibleSolution dengan level yang baru
	var tempPossibleSolutions [][]string
	// penelusuran untuk setiap array of string pada possible solutions
	for _, stringArray := range possibleSolutions {
		// sudah dikunjungi (elemen terakhir dari setiap array of string) -> iterasi selanjutnya
		if isInArray(visited, lastElement(stringArray)) {
			continue
		} else {
			// ketemu
			if lastElement(stringArray) == desiredPageName {
				return stringArray
			} else {
				// daftar link dari elemen terakhir (yang akan diexpand)
				linkArray := GetLinks(lastElement(stringArray))
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
	return BFS(tempPossibleSolutions, visited, desiredPageName)
}

func CallBFS(initialPageName string, desiredPageName string) ([]string, bool) {
	found := true
	possibleSolutions := [][]string{{initialPageName}}
	var visited []string
	solution := BFS(possibleSolutions, visited, desiredPageName)
	if !isInArray(solution, desiredPageName) {
		found = false
	}
	return solution, found
}
