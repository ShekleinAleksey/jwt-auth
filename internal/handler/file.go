package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary UploadFile
// @Tags files
// @Description uploadFile
// @ID uploadFile
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Файл для загрузки"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /files/upload [post]
func (h *Handler) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "ошибка при получении файла")
		return
	}

	// Сохраняем файл
	if err := c.SaveUploadedFile(file, "uploads/"+file.Filename); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "ошибка при сохранении файла")
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "походу загрузилось",
	})
}
