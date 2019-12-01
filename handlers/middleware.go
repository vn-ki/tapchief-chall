package handlers

import (
	"encoding/json"
	"net/http"
)

type apiResponse struct {
	Status   string `json:"status"`
	Response string `json:"response"`
}

// TODO: Move into other file
func isAPI(method string, handlerFunc func(http.ResponseWriter, *http.Request) apiResponse) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == method {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
		resp := handlerFunc(w, r)
		json.NewEncoder(w).Encode(resp)
	}
}
