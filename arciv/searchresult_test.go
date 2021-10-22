package arciv

import (
	"encoding/xml"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

// Test decoding an xml file from arXiv query response
// into a SearchResult
func TestSearchResultUnmarshal(t *testing.T) {
	var result SearchResult

	xmlFile, err := os.Open(filepath.Join("..", "testdata", "queryresult.xml"))

	if err != nil {
		t.Error(err)
	}

	defer xmlFile.Close()

	body, err := io.ReadAll(xmlFile)

	if err != nil {
		t.Error(err)
	}

	err = xml.Unmarshal(body, &result)

	if err != nil {
		t.Error(err)
	}

	if reflect.DeepEqual(result, SearchResult{}) {
		t.Error("Failed to unmarshal SearchResult from xml")
	}
}
