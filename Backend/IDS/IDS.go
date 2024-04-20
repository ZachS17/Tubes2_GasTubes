package IDS

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetWikipediaLinks(pageName string) ([]string, error) {
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

	// Check if "parse" key exists in the response
	parseData, ok := data["parse"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("parse data not found in JSON response")
	}

	// Extract links from the response
	linksData, ok := parseData["links"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("links data not found in parse section of JSON response")
	}

	linkTitles := make([]string, len(linksData))
	for i, link := range linksData {
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
// func dls(currentPageName string, desiredPageName string, visited []string, solution []string, currentDepth int, desiredDepth int) []string {
// 	// // pembuatan array visited dan solution
// 	// if visited == nil {
// 	// visited = make([]string, 0)
// 	// solution = make([]string, 0)
// 	// }
// 	if currentDepth == desiredDepth { // sudah capai depth
// 		fmt.Println("Link yang diperiksa:" + currentPageName)
// 		fmt.Println("Yang dibandingkan: " + desiredPageName)
// 		return solution
// 	} else {
// 		if currentPageName == desiredPageName { // ketemu
// 			return solution
// 		} else { // blm ketemu
// 			if !isInArray(visited, currentPageName) { // blm dikunjungi
// 				links, err := GetWikipediaLinks(currentPageName)
// 				if err != nil {
// 					return nil
// 				} else {
// 					for _, link := range links {
// 						visited = append(visited, link)
// 						solution = append(solution, link)
// 						solution = dls(link, desiredPageName, visited, solution, currentDepth+1, desiredDepth)
// 					}
// 				}
// 			}
// 		}
// 	}

// 	// Add a return statement at the end of the function
// 	return solution
// }

func dls(currentPageName string, desiredPageName string, visited []string, solution []string, currentDepth int, desiredDepth int) []string {
	// fmt.Println("Link yang ditelusuri: " + currentPageName)
	// fmt.Println("Link yang mau dicari: " + desiredPageName)

	// utamain ketemu
	if currentPageName == desiredPageName {
		return solution
	} else {
		// utamain kedalaman
		if currentDepth == desiredDepth {
			// fmt.Println("Reached desired depth")
			return solution
		} else {
			// utamain sudah visited
			if isInArray(visited, currentPageName) {
				// fmt.Println("Already visited:", currentPageName)
				return solution
			} else {
				visited = append(visited, currentPageName)
				links, err := GetWikipediaLinks(currentPageName)
				// menangani error dapat link
				if err != nil {
					fmt.Println("Error getting links for", currentPageName, ":", err)
					return solution
				}
				for _, link := range links {
					// Recursively search for the desired page
					newSolution := append([]string{}, solution...)
					newSolution = append(newSolution, link)
					newSolution = dls(link, desiredPageName, visited, newSolution, currentDepth+1, desiredDepth)
					if len(newSolution) > len(solution) && newSolution[len(newSolution)-1] == desiredPageName {
						return newSolution // Return immediately if the desired page is found
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
	hasil := make([]string, 0)
	copyInitialPageName := initialPageName // Initialize copyInitialPageName outside the loop
	for !found && currentDepth <= desiredDepth {
		fmt.Println(copyInitialPageName, desiredPageName, currentDepth)
		visited := make([]string, 0)
		solution := []string{copyInitialPageName}
		fmt.Println(visited, solution)
		hasil = dls(copyInitialPageName, desiredPageName, visited, solution, 0, currentDepth) // Pass 0 as current depth to dls
		if isInArray(hasil, desiredPageName) {
			found = true
		}
		fmt.Println(hasil)
		currentDepth = currentDepth + 1
		copyInitialPageName = initialPageName // Reset copyInitialPageName to initialPageName for the next iteration
	}
	return hasil, found
}
