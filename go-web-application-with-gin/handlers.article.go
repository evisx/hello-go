package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	render(c, gin.H{
		"title": "Home Page",
		"payload": getAllArticles(),
	}, "index.html")
}

func showArticle(c *gin.Context) {
	article := article{
		ID: 1,
		Title: "title test only",
		Content: "content test only",
	}

	render(c, gin.H{
		"title": article.Title,
		"payload": article,
	}, "article.html")
}

func render(c *gin.Context, data gin.H, templateName string) {
	c.HTML(http.StatusOK, templateName, data)
}