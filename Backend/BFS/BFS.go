package BFS

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// masalah -> DFS bukan BFS
func getWikipediaLinks(pageName string) ([]string, error) {
	// URL of the Wikimedia API endpoint
	apiUrl := "https://en.wikipedia.org/w/api.php"

	// Parameters for the API request
	params := map[string]string{
		"action": "parse",
		"page":   pageName,
		"format": "json",
		"prop":   "links",
	}

	// Construct the URL with query parameters
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %v", err)
	}
	q := req.URL.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	// Make the API request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making API request: %v", err)
	}
	defer resp.Body.Close()

	// Decode the JSON response
	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("error decoding JSON response: %v", err)
	}

	// Extract links from the response
	links := data["parse"].(map[string]interface{})["links"].([]interface{})
	linkTitles := make([]string, len(links))
	for i, link := range links {
		title := link.(map[string]interface{})["*"].(string)
		linkTitles[i] = title
	}

	return linkTitles, nil
}

func isInArray(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

// rekursif
// arrayToDiscover berisi semua link dari setiap level
func BFS(arrayToDiscover []string, visited []string, desiredPageName string, solution []string) []string {
	// menyimpan semua yang diexpand untuk langsung dipassing tanpa perubahan pada arrayToDiscover
	tempArrayToDiscover := make([]string, 0)
	for _, item := range arrayToDiscover {
		if isInArray(visited, item) { // sudah dikunjungi -> iterasi selanjutnya
			continue
		} else {
			if item == desiredPageName { // ketemu
				return solution
			} else {
				temp, err := getWikipediaLinks(item)
				if err == nil { // kalau ada hasilnya
					tempArrayToDiscover = append(tempArrayToDiscover, temp...)
				}
				// masalah -> harus sambil simpan induknya
				// solusi (mungkin) -> array of string elemennya, dipisah terakhir untuk pengecekan sudah visit dan ditambah di akhir untuk pencarian berikutnya
			}
		}
	}
	return solution
}
