package arciv

import (
	"fmt"
)

const baseQuery = "query?"

type Query interface {

	// Convert the query into a query string for the url
	QueryString() string
}

type SimpleQuery struct {
	Search     string
	MaxResults int
}

func (s SimpleQuery) QueryString() string {
	return fmt.Sprintf(
		"%vsearch_query=all:%v&start=0&max_results=%v",
		baseQuery, s.Search, s.MaxResults)
}

type AdvancedQuery struct {
	Search     string
	IdList     []string
	StartIndex int
	MaxResults int
	SortBy     SortStrategy
	SortOrder  SortRule
}

func (s AdvancedQuery) QueryString() string {
	return fmt.Sprintf(
		"%vsearch_query=all:%v&start=%v&max_results=%v",
		baseQuery, s.Search, s.StartIndex, s.MaxResults)
}

type SortStrategy string

const (
	Relevance      SortStrategy = "relevance"
	LastUpdateDate SortStrategy = "lastUpdateDate"
	SubmittedDate  SortStrategy = "submittedDate"
)

type SortRule string

const (
	Ascending  SortRule = "ascending"
	Descending SortRule = "descending"
)
