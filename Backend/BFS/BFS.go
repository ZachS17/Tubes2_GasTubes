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
// terima parameter visitedArray, target
func BFS(currentPageName string, pageName string, visited []string, solution []string) []string {
	if visited == nil {
		// If the slice is nil, initialize it with an empty slice
		visited = make([]string, 0)
	}
	if solution == nil {
		// If the slice is nil, initialize it with an empty slice
		solution = make([]string, 0)
	}
	if currentPageName == pageName { // ketemu
		return solution
	} else { // blm ketemu
		if !isInArray(visited, currentPageName) { // blm dikunjungi
			links, err := getWikipediaLinks(currentPageName)
			if err != nil {
				fmt.Println("Error!")
				return solution
			} else {
				for _, link := range links {
					visited = append(visited, link)
					// Recursively call BFS and store the result in a local variable
					// result := BFS(link, pageName, visited, solution)
					solution = append(solution, link)
					// Concatenate the local result with the current solution
					// solution = append(solution, result...)
					solution = BFS(link, pageName, visited, solution)
				}
			}
		}
	}

	// Add a return statement at the end of the function
	return solution
}
