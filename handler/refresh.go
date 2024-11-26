package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func refresh(c *gin.Context) {
	fmt.Println(c)
	c.JSON(http.StatusOK, gin.H{"data": "your data here"})
}
