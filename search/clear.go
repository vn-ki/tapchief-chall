package search

// ClearIndex clears indices and documents by making new maps for each.
func ClearIndex() {
	indices.lock.Lock()
	documents.lock.Lock()
	indices.indexMap = indexMapType{}
	documents.documentMap = documentMapType{}
	documents.lock.Unlock()
	indices.lock.Unlock()
}
