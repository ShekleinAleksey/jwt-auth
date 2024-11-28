package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Request(c *gin.Context) {
	guid := c.Query("guid")
	fmt.Println(guid)

	_, err := uuid.Parse(guid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid GUID format"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"method":  c.Request.Method,
		"path":    c.Request.URL.Path,
		"headers": c.Request.Header,
		"query":   c.Request.URL.Query(),
		"body":    c.Request.Body})
}
