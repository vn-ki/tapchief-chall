package search

import (
	"sync"
)

type indexEntryType struct {
	DocumentID string
	Position   int
}

type indexMapType map[string][]indexEntryType

type indexType struct {
	indexMap indexMapType
	lock     sync.RWMutex
}

type documentMapType map[string]string

type documentType struct {
	documentMap documentMapType
	lock        sync.RWMutex
}

// SearchResult contains the document which contains the word and the position of the word search
type SearchResult struct {
	Document string `json:"document"`
	Position []int  `json:"position"`
}

var indices = indexType{indexMapType{}, sync.RWMutex{}}
var documents = documentType{documentMapType{}, sync.RWMutex{}}
