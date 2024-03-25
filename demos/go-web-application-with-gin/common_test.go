package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

// This function is used for setup before executing the test functions
func TestMain(m *testing.M) {
  //Set Gin to Test Mode
  gin.SetMode(gin.TestMode)

  // Run the other tests
  os.Exit(m.Run())
}

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

// Helper function for simple HTTP testing
func simpleTestHTTPResponse(t *testing.T, method string, path string, config func(r *gin.Engine), f func (w *httptest.ResponseRecorder) bool) {
	r := initializeRouter(config)
	req, _ := http.NewRequest(method, path, nil)
	testHTTPResponse(t, r, req, f)
}
