package router

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouterHTTPSetup(t *testing.T) {
	HTTPSetup()

	req := httptest.NewRequest("GET", "/quotes", nil)
	w := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(w, req)
	if w.Code == http.StatusNotFound {
		t.Error("/quotes handler not registered")
	}

	req = httptest.NewRequest("GET", "/quotes/random", nil)
	w = httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(w, req)
	if w.Code == http.StatusNotFound {
		t.Error("/quotes/random handler not registered")
	}
}
