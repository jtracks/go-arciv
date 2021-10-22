package arciv

import (
	"log"
	"testing"
)

// Test searching the arXiv database
func TestSearch(t *testing.T) {
	var result *SearchResult
	var err error

	result, err = Search(
		SimpleQuery{
			Search:     "electron",
			MaxResults: 5,
		},
	)

	if err != nil {
		log.Printf("error occurred during Search(): %v", err)
		t.FailNow()
	}

	if result == nil {
		log.Printf("empty SearchResult: %v", result)
		t.FailNow()
	}

	if len(result.Entries) != 5 {
		log.Printf("invalid length for SearchResult.entries: %v", result.Entries)
		t.FailNow()
	}

}
