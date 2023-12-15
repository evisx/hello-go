package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	render(c, gin.H{
		"title": "Home Page",
	}, "index.html")
}

func render(c *gin.Context, data gin.H, templateName string) {
	c.HTML(http.StatusOK, templateName, data)
}