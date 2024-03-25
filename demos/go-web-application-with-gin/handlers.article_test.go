package main

import (
	"io/ioutil"
	"strings"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// Test that a GET request to the home page returns the home page with
// the HTTP code 200 for an unauthenticated user
func TestShowIndexPageUnauthenticated(t *testing.T) {
	simpleTestHTTPResponse(t, "GET", "/",
		// set server
		func (r *gin.Engine) {
			r.GET("/", showIndexPage)
		},
		// test response logic
		func (w *httptest.ResponseRecorder) bool {
			// check result
			statusOK := w.Code == http.StatusOK

			// Test that the page title is "Home Page"
			// You can carry out a lot more detailed tests using libraries that can
			// parse and process HTML pages
			p, err := ioutil.ReadAll(w.Body)
			pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

			return statusOK && pageOK
		})
}

func TestShowArticleUnauthenticated(t *testing.T) {
	simpleTestHTTPResponse(t, "GET", "/article/view/1",
		func (r *gin.Engine) {
			r.GET("/article/view/:article_id", showArticle)
		},
		func (w *httptest.ResponseRecorder) bool {
			return w.Code == http.StatusOK
		})

	simpleTestHTTPResponse(t, "GET", "/article/view/11",
		func (r *gin.Engine) {
			r.GET("/article/view/:article_id", showArticle)
		},
		func (w *httptest.ResponseRecorder) bool {
			return w.Code == http.StatusNotFound
		})
}
