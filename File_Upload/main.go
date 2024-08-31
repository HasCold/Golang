package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() // default returns a gin engine instance
	r.POST("/upload", handleUpload)

	srv := &http.Server{
		Addr:    ":5000",
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("listen: %v \n", err)
	}
}

// When we upload any file to API that will send as a Multipart form
func handleUpload(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	files := form.File["files"]
	for key, file := range files {
		err := saveFile(file, key)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Files Uploaded Successfully"})
}

func saveFile(fileHeader *multipart.FileHeader, key int) error {
	fmt.Println("File Header File Name :-", fileHeader.Filename)

	// If want to extract the file Name and extension
	ext := strings.Split(fileHeader.Filename, ".") // -->> Separate the file name and their extension
	if ext[1] != "png" {
		return errors.New("Pass a valid png file")
	}

	src, err := fileHeader.Open()
	if err != nil {
		return err
	}

	defer src.Close()

	// Create a new file in the desired destination folder
	dstPath := filepath.Join("./destination", ext[0]+"-"+strconv.Itoa(key)+"."+ext[1]) // Itoa is equivalent to FormatInt(int64(i), 10)
	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}

	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	return nil
}
