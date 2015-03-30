package main

import "net/http"

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

func main() {

	mux := multiplexer{
		"/":        foo{},
		"/about":   page{"about"},
		"/contact": page{"contact"},
	}

	http.ListenAndServe(":8000", mux)
}
