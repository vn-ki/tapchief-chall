package search

import (
	"regexp"
	"strings"

	"github.com/rs/xid"
)

// Index splits the given para into documents and constructs the dictionary
func Index(para string) error {
	// This has to be done because sometimes \r is used instead of \n.
	// Example provided uses both \r and \n(?)
	re := regexp.MustCompile("\\r")
	para = re.ReplaceAllString(para, "\n")

	documents := strings.Split(para, "\n\n")

	for _, document := range documents {
		dID := xid.New().String()

		for idx, word := range strings.Fields(document) {
			word = strings.ToLower(word)

			indices.lock.Lock()
			if _, ok := indices.indexMap[word]; !ok {
				indices.indexMap[word] = make([]indexEntryType, 0)
			}
			indices.indexMap[word] = append(indices.indexMap[word], indexEntryType{dID, idx})
			indices.lock.Unlock()
		}
		err := storeDocument(dID, document)
		if err != nil {
			return err
		}
	}
	return nil
}
