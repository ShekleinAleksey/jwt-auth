package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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

	c.JSON(http.StatusOK, gin.H{
		"message":  "файл успешно загружен",
		"filename": file.Filename,
	})
}
