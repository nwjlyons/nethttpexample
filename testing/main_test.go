package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAppendTrailingSlashMiddleware(t *testing.T) {
	server := httptest.NewServer(handlers)

	resp, _ := http.Get(server.URL + "/about")

	if resp.Request.URL.Path != "/about/" {
		t.Error("Didn't redirect to URL with trailing slash.")
	}
}
