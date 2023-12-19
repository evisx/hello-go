package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
)

// Helper function to get router prepared for testing
func initializeRouter(config func(r *gin.Engine)) *gin.Engine {
	// create router and load templates
	r := getRouterWithInitialized(true)
	// add path and handler
	config(r)
	return r
}

// Helper function for HTTP testing
func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	// create recorder
	w := httptest.NewRecorder()
	// simulate
	r.ServeHTTP(w, req)
	if !f(w) {
		t.Fail()
	}
}
