package services

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetImage(ctx *gin.Context) {
	filename := ctx.Param("filename")

	file, err := os.ReadFile("uploads/image/" + filename)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid filename",
		})
		return
	}

	ctx.Header("Content-Type", "image/png")

	ctx.Data(http.StatusOK, "application/octet-stream", file)
}