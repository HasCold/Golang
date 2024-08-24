package main

// install Gin-Gonic package :-
// 1. go get -u github.com/gin-gonic/gin

// Add these gin-swagger packages
// 2. go get -u github.com/swaggo/swag/cmd/swag
// 3. go get -u github.com/swaggo/gin-swagger
// 4. go get -u github.com/swaggo/files

// -------------------------------------- Docs ----------------------------------------
// Swagger Integration in Gin :-
// GitHub - swaggo/gin-swagger  Link :- https://github.com/swaggo/gin-swagger
// GitHub - swaggo/swag   Link :- https://github.com/swaggo/swag 		|		<Important>
// These repositories provide documentation and examples on how to use Swagger with the Gin framework.

import (
	"log"
	"net/http"

	_ "swagger/docs" // Correct import path based on your module path

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// These are the main swagger parameters which are placed in the main file

// @title           Swagger Gin Config
// @version         1.0
// @description     This is sample
// @termsOfService

// @contact.name   Hasan
// @contact.url    http://www.swagger.io/support
// @contact.email  hasantest@swagger.io

// @license.name
// @license.url

// @host      localhost:5000
// @BasePath  /api/v1

func main() {

	router := gin.Default() // Default returns an engine instance which is used to build a middleware, logger and routing purposes

	router.GET("/greeting", Greeting)

	url := ginSwagger.URL("http://localhost:5000/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	err := router.Run(":5000")
	if err != nil {
		log.Fatal(err) // when critical errors encounter in the program which stops the continuation of the program so we have to log the error messages and then immediately terminates the program with a non-zero exit status code
	}
}

// --------------------------------------------------------------------------------------------------------
// When you mention all the specified parameters before the main function and also specified parameter before every function just like below then  you can run the command
// 1. swag init .
// 2. go run .
// 3. open http://localhost:5000/swagger/index.html in your browser to see the documentation
// --------------------------------------------------------------------------------------------------------

// These specific parameters placed before every functions

// Greeting godoc
// @Summary greeting service
// @Description  this is greeting service
// @Tags         Greeting
// @Accept       */*
// @Produce      json
// @Success      200  {object}   json
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /greeting [get]

func Greeting(c *gin.Context) {
	result := map[string]interface{}{ // Initialize the map
		"data": "This is response",
	}

	// c is a pointer to the gin context struct ;
	// Passing *gin.Context as a pointer is efficient because it avoids copying the entire context structure, which could be large. It also ensures that any modifications to the context (e.g., setting a response header) are reflected across all parts of the code handling the request.

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    result,
	})
}

// --------------------------------------------------------------------------------------------------------
// For Global Installation (go install):
// If you want to use swag as a tool on your machine, independent of any particular Go project, go install is better. This is more appropriate for tools that you’ll use across multiple projects.
// -->> go install github.com/swaggo/swag/cmd/swag@latest

// For Project-Specific Dependency (go get):
// If you need swag as part of your project’s dependencies, and you want to ensure that everyone working on the project uses the same version, go get -u is better. This method also tracks the package in your go.mod file, making it easier to manage dependencies.
// -->> go get -u github.com/swaggo/swag/cmd/swag

// Run the Swag at your Go project root path
// -->> swag init .
