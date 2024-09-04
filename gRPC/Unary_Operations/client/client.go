package main

import (
	"context"
	"net/http"
	proto "unary_operat/protoc" // Alias Keyword mention

	"github.com/gin-gonic/gin"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Why we are using the client ?

// Server is not exposed outside becuase it is making a tcp connection so we have to implement the REST Api which takes the request from client/user and prepare that request for our grpc server

var client proto.ExampleClient

func main() {
	// 	Connection to internal grpc server
	//  making the internal tcp connection
	conn, err := grpc.NewClient("localhost:5000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client = proto.NewExampleClient(conn)
	//

	// Implement REST API
	r := gin.Default()
	r.GET("/sent-message-to-server/:message", clientConnection)
	r.Run(":8000") // By-Default run at :8080

	// Preparing the gRPC request from REST API
	// req := &proto.HelloRequest{SomeString: "Hello Golang, from Hasan !"}

	// client.ServerReply(context.TODO(), req)
}

func clientConnection(c *gin.Context) { // c is a pointer to gin.Context struct
	message := c.Param("message")

	// Prepaing a gRPC request
	req := &proto.HelloRequest{SomeString: message}
	client.ServerReply(context.TODO(), req)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Message sent successfully to the server" + message,
	})
}
