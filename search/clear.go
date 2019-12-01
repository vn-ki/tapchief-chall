package search

import "sync"

// ClearIndex clears indices and documents by making new maps for each.
func ClearIndex() {
	indices = indexType{indexMapType{}, sync.RWMutex{}}
	documents = documentType{documentMapType{}, sync.RWMutex{}}
}
