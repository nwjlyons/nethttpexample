package main

import "net/http"

type foo struct{}

func (f foo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("foo"))
}

func main() {
	http.ListenAndServe(":8000", foo{})
}
