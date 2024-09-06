package main

import (
	"context"
	"fmt"
	"net/http"
	proto "unary_operat/Client_Streaming/protoc" // Alias Keyword mention

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
	r.GET("/sent-stream", clientConnectionServer)
	r.Run(":8000") // By-Default run at :8080

	// Preparing the gRPC request from REST API
	// req := &proto.HelloRequest{SomeString: "Hello Golang, from Hasan !"}

	// client.ServerReply(context.TODO(), req)
}

func clientConnectionServer(c *gin.Context) { // c is a pointer to gin.Context struct
	// Static Request
	req := []*proto.HelloRequest{
		{SomeString: "Request 1"},
		{SomeString: "Request 2"},
		{SomeString: "Request 3"},
		{SomeString: "Request 4"},
		{SomeString: "Request 5"},
		{SomeString: "Request 6"},
	}

	// Prepaing a gRPC request
	stream, err := client.ServerReply(context.TODO())
	if err != nil {
		fmt.Println("something error")
		return
	}

	for _, msg := range req {
		// we are sending the message to the server
		err = stream.Send(msg)
		if err != nil {
			fmt.Println("request not fulfill")
			return
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Println("There is some error occur", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":       true,
		"message_count": res,
	})
}
