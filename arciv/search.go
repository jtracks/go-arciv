package arciv

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
)

const BaseURL = "http://export.arxiv.org/api/"

// Call the arxiv search api using a SimpleQuery or AdvancedQuery
// return *SearchResult, error
func Search(query Query) (*SearchResult, error) {

	return CustomSearch(query.QueryString())

}

// Call the arxiv search api using a string as query
// return *SearchResult, error
func CustomSearch(query string) (*SearchResult, error) {

	var result SearchResult
	var err error

	url := BaseURL + query

	log.Printf("Fetching url: %v\n", url)
	resp, err := http.Get(url)

	if err != nil {
		return &result, err
	}

	log.Printf("Response code: %v\n", resp.StatusCode)
	if resp.StatusCode != 200 {
		return &result, fmt.Errorf(
			"request failed with code: %v",
			resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return &result, err
	}

	err = xml.Unmarshal(body, &result)

	if err != nil {
		return &result, err
	}

	return &result, nil
}
