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

	var ret []SearchResult
	for _, value := range retmap {
		ret = append(ret, value)
	}

	return ret, nil
}
