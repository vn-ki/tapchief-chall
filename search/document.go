/* The functions in this file abstracts away storing and retreival of documents.
 * Right now these use maps, but later can be changed to use some kind of permanent storage.
 */
package search

import (
	"errors"
)

func storeDocument(dID string, data string) error {
	documents.lock.Lock()
	documents.documentMap[dID] = data
	documents.lock.Unlock()
	return nil
}

func getDocument(dID string) (string, error) {
	documents.lock.RLock()
	data, ok := documents.documentMap[dID]
	documents.lock.RUnlock()
	if !ok {
		return "", errors.New("dID does not exist")
	}

	return data, nil
}
