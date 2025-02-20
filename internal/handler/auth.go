package handler

import (
	"net/http"

	"github.com/ShekleinAleksey/jwt-auth/internal/entity"
	"github.com/gin-gonic/gin"
)

type signUpInput struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body entity.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input signUpInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.service.AuthService.CreateUser(entity.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  input.Password,
	})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary SignIn
// @Tags auth
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body signInInput true "credentials"
// @Success 200 {string} string "token"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.AuthService.GenerateToken(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

type TokenDetails struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at"`
}

// @Summary GetUsers
// @Tags auth
// @Description get users
// @ID get-users
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.User
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/get-users [get]
func (h *Handler) GetUsers(c *gin.Context) {
	users, err := h.service.AuthService.GetUsers()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}

// func (h *Handler) Token(c *gin.Context) {
// 	guid := c.Query("guid")
// 	ip := c.ClientIP()
// 	fmt.Println("ip: ", ip)
// 	tokens, err := createToken(guid, ip)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, err.Error())
// 	}
// 	// _, err := uuid.Parse(guid)
// 	// if err != nil {
// 	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid GUID format"})
// 	// 	return
// 	// }
// 	c.JSON(http.StatusOK, tokens)
// }
