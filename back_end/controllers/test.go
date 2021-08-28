package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestHandler(c *gin.Context) {
	id := c.Query("id")
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": id,
	})
}
