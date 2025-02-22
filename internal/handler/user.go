package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary DeleteUser
// @Tags user
// @Description delete user
// @ID delete-user
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/{id} [delete]
func (h *Handler) deleteUser(c *gin.Context) {
	userId := c.Param("id")

	if err := h.service.UserService.DeleteUser(userId); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
