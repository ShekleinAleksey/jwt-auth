package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Refresh(c *gin.Context) {
	var requestBody struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
		Ip           string
	}

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userID := "extracted_user_id"
	ip := "123"
	tokens, err := createToken(userID, ip)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, tokens)
}
