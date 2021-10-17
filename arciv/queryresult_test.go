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
// into a QueryResult
func TestQueryResultUnmarshal(t *testing.T) {
	var query QueryResult

	xmlFile, err := os.Open(filepath.Join("..", "testdata", "queryresult.xml"))

	if err != nil {
		t.Error(err)
	}

	defer xmlFile.Close()

	body, err := io.ReadAll(xmlFile)

	if err != nil {
		t.Error(err)
	}

	err = xml.Unmarshal(body, &query)

	if err != nil {
		t.Error(err)
	}

	if reflect.DeepEqual(query, QueryResult{}) {
		t.Error("Failed to unmarshal QueryResult from xml")
	}
}
