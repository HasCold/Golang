package main

// Gin-Gonic Github :- https://github.com/gin-gonic/gin

// Command to use the github package :-
// go mod init goGin
// go mod tidy  // Jo package mera is file me use hua ha usko ye internally bana ke rakh lega

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type api struct {
	Name  string `json:"name"`  // this is corrsponding api return data-type that means response have must the key name
	Email string `json:"email"` // this is corrsponding api return data-type that means response have must the key email
}

var data api

func main() {
	r := gin.Default() // Initialize the default gin engine instance
	r.GET("/get", getValue)
	r.POST("/post", postValue)
	r.PUT("/put", putValue)
	r.DELETE("/delete", delValue)
	r.Run(":5000") // Gin-gonic server running on 5000 port // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// *gin.Context is a pointer to the gin.Context struct.
// Passing *gin.Context as a pointer is efficient because it avoids copying the entire context structure, which could be large. It also ensures that any modifications to the context (e.g., setting a response header) are reflected across all parts of the code handling the request.

func getValue(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": data,
		"success": true,
	})
}

func postValue(c *gin.Context) {
	err := c.BindJSON(&data) // In data we have api so we are binding them into the JSON object
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Something went wrong",
			"success": false,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": data,
		"success": true,
	})
}

func putValue(c *gin.Context) {
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Something went wrong",
			"success": false,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": data,
		"success": true,
	})
}

func delValue(c *gin.Context) {
	data = api{}
	c.JSON(http.StatusOK, gin.H{
		"message": data,
		"success": true,
	})
}
