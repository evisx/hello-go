package main

import (
  "io/ioutil"
  "net/http"
  "net/http/httptest"
  "strings"
  "testing"
)

// Test that a GET request to the home page returns the home page with
// the HTTP code 200 for an unauthenticated user
func TestShowIndexPageUnauthenticated(t *testing.T) {
	// set server
	r := getRouterWithInitialized(true)
	r.GET("/", showIndexPage)
	// create a request
	req, _ := http.NewRequest("GET", "/", nil)
	// create a response recorder
	w := httptest.NewRecorder()
	// simulate
	r.ServeHTTP(w, req)
	// check result
	statusOK := w.Code == http.StatusOK

	// Test that the page title is "Home Page"
    // You can carry out a lot more detailed tests using libraries that can
    // parse and process HTML pages
    p, err := ioutil.ReadAll(w.Body)
    pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

	if !(statusOK && pageOK) {
		t.Fail()
	}
}