package main

// Framework :- an essential supporting structure of a building applications
// Types of frameworks in golang
// Gin / Gin-Gonic   --->> Top most used framework
// Beego
// Echo

// Gin-Gonic Github :- https://github.com/gin-gonic/gin

// Command to use the github package :-
// go mod init goGin
// go mod tidy  // Jo package mera is file me use hua ha usko ye internally bana ke rakh lega

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()   // Default returns an engine instance which is used to build a middleware , logger and routing purposes
	r.GET("/ping", test) // ("/path", Callback function)
	r.GET("/test2", test2)
	r.Run(":5000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{ // Pre-defined status code in golang net/http package
		"message": "pong",
	})
}

func test2(c *gin.Context) {
	c.JSON(201, gin.H{ // Pre-defined status code in golang net/http package
		"message": "Created successfully",
	})
}
