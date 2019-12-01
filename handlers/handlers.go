package handlers

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/labstack/gommon/log"
	"github.com/vn-ki/tapchief-chall/search"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/index", http.StatusFound)
}

func ClearAPIHandler() http.HandlerFunc {
	return isAPI("GET", func(w http.ResponseWriter, r *http.Request) apiResponse {
		search.ClearIndex()
		return apiResponse{"sucess", "index cleared"}
	})
}

func IndexAPIHandler() http.HandlerFunc {
	return isAPI("POST", func(w http.ResponseWriter, r *http.Request) apiResponse {
		r.ParseForm()
		para, ok := r.Form["para"]
		if !ok {
			return apiResponse{"error", "parameter 'para' missing"}
		}
		log.Debugf("searching for %s", para[0])
		search.Index(para[0])
		return apiResponse{"success", "paragraph indexed"}
	})
}

func SearchAPIHandler() http.HandlerFunc {
	return isAPI("GET", func(w http.ResponseWriter, r *http.Request) apiResponse {
		params := r.URL.Query()
		query, ok := params["query"]
		if !ok {
			return apiResponse{"error", "parameter query not found"}
		}
		results, err := search.Search(query[0])
		if err != nil {
			return apiResponse{"error", err.Error()}
		}
		out, err := json.Marshal(results)
		if err != nil {
			return apiResponse{"error", err.Error()}
		}
		return apiResponse{"success", string(out)}
	})
}

func ClearHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		search.ClearIndex()
		tmpl.Execute(w, "Index cleared")
	}
}

func IndexHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("asdf")
		if r.Method == "POST" {
			r.ParseForm()
			para, ok := r.Form["para"]
			if !ok {
				log.Info("not found para")
				return
			}
			log.Infof("indexing %s", para[0])
			search.Index(para[0])
			tmpl.Execute(w, "Paragraph indexed")
		} else {
			err := tmpl.Execute(w, nil)
			if err != nil {
				log.Error(err)
			}
			return
		}
	}
}

func SearchHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()
			query, ok := r.Form["query"]
			if !ok {
				log.Error("not found")
				return
			}
			results, err := search.Search(query[0])
			if err != nil {
				log.Error(err)
				return
			}
			log.Info("executing")
			log.Info(results)
			tmpl.Execute(w, results)
			return
		} else {
			err := tmpl.Execute(w, "asdf")
			if err != nil {
				log.Error(err)
			}
			return
		}
	}
}
