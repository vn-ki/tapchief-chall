package search

import "strings"

// Search searches for the word in the dictionary and return all the documents
func Search(query string) ([]SearchResult, error) {
	query = strings.TrimSpace(strings.ToLower(query))

	indices.lock.RLock()
	searched, ok := indices.indexMap[query]
	indices.lock.RUnlock()

	if !ok {
		return []SearchResult{}, nil
	}

	var retmap = map[string]SearchResult{}

	// When we index, we store each position as a different indexed value.
	// In this loop, we convert the positions from a single document into a
	// single SearchResult object
	for _, s := range searched {
		data, err := getDocument(s.DocumentID)
		if err != nil {
			return []SearchResult{}, err
		}

		if res, ok := retmap[s.DocumentID]; !ok {
			retmap[s.DocumentID] = SearchResult{data, []int{s.Position}}
		} else {
			res.Position = append(res.Position, s.Position)
			retmap[s.DocumentID] = res
		}
	}

	ret := mapToSlice(retmap)

	return ret, nil
}

func mapToSlice(m map[string]SearchResult) []SearchResult {
	var ret []SearchResult
	for _, value := range m {
		ret = append(ret, value)
	}
	return ret
}
