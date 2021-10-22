package main

import (
	"fmt"

	"github.com/jtracks/go-arciv/arciv"
)

func main() {

	var err error
	var result *arciv.SearchResult

	result, err = arciv.Search(
		arciv.SimpleQuery{
			Search:     "electron",
			MaxResults: 5,
		})

	if err != nil {
		panic(err)
	}

	for i, e := range result.Entries {
		fmt.Printf("Result %v: %v\n", i+1, e.Title)
	}
}
