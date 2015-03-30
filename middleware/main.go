package main

import (
	"net/http"
	"strings"
)

type foo struct{}

func (f foo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("foo"))
}

type page struct {
	body string
}

func (p page) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(p.body))
}

// string is the URL path and http.Handler is any type that has a ServeHTTP method.
type multiplexer map[string]http.Handler

func (m multiplexer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler, ok := m[r.RequestURI]; ok {
		handler.ServeHTTP(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func appendTrailingSlash(h http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if !strings.HasSuffix(r.RequestURI, "/") {
			http.Redirect(w, r, r.RequestURI+"/", http.StatusFound)
		} else {
			h.ServeHTTP(w, r)
		}
	})
}

func main() {

	mux := multiplexer{
		"/":         foo{},
		"/about/":   page{"about"},
		"/contact/": page{"contact"},
	}

	http.ListenAndServe(":8000", appendTrailingSlash(mux))
}
