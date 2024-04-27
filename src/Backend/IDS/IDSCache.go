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

func GetWikipediaLinksCache(URL string) []string {
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
