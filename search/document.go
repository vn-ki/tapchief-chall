package search

import (
	"errors"
)

func storeDocument(dID string, data string) error {
	//return ioutil.WriteFile(file, []byte(data), 0644)
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
