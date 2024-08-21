package main

// Command to use the github package :-
// go mod init goGin
// go mod tidy  // Jo package mera is file me use hua ha usko ye internally bana ke rakh lega

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() // Initialize the Default gin engine
	r.GET("/get", getValues)
	r.Run(":5000")
}

var url = "http://date.jsontest.com/"

func getValues(c *gin.Context) { // c is a pointer to the gin.Context struct
	resp, err := http.Get(url)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
	}

	// Close the response body
	// Here we use the defer keyword bc when this function execution will done so then we are closing the response body to avoid breaching of data
	// Basically hame response object ma ak field Body milti ha jo ke ak io.ReadClose() operation so after perform some processing or read the resposne we have to close the response Body after the function execution has been done and if we don't close the body then the system might be exhausted or there could be some memory or resource leak issue.
	defer resp.Body.Close()

	// Read all response body
	data, err := io.ReadAll(resp.Body) // return the data in the byte array and also return error if present
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
	}

	var target map[string]interface{} // key is string and value is interface type

	// json.Unmarshal(The data which you want to bind in the byte-array, The data which you want to bind with,just passed the reference of memory address)
	er := json.Unmarshal(data, &target)
	if er != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": target["milliseconds_since_epoch"],
	})

}

//
// Why Use defer?
// Resource Management:
// When you make an HTTP request using http.Get(url), it returns an http.Response object that includes a Body field representing the response's content. This Body field is an io.ReadCloser, meaning it needs to be closed after you are done with it to free up resources.
// If you forget to close the Body, it can lead to a resource leak, which might exhaust system resources like file descriptors or memory over time.

// Guaranteed Execution:
// By placing the resp.Body.Close() inside a defer, you're ensuring that it will be executed no matter how the function exits, whether it's a normal return or due to an error.

// Readability and Maintainability:
// The defer statement is often placed right after the resource is acquired (e.g., after a successful http.Get call). This makes the code easier to read and maintain, as the resource cleanup code is located near the resource acquisition, reducing the chances of forgetting to close it.
