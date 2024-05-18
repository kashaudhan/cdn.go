package services

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)



func UploadImage(ctx *gin.Context) {
	// var file File

	fileHeader, err := ctx.FormFile("image")
	fileName := ctx.PostForm("name")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Please provide a valid image",
		})
		return
	}

	file, err := fileHeader.Open()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Please provide a valid image",
		})
		return
	}

	defer file.Close()

	if fileName != "" {
		fileName = fileName + filepath.Ext(fileHeader.Filename)
	} else {
		fileName = fileHeader.Filename
	}

	f, err := os.Create("uploads/image/" + fileName);

	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create file",
		})
		return
	}

	_, err = io.Copy(f, file);

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error while copying image",
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Image saved",
	})

}