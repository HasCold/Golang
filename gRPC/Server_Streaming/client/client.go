package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
	proto "unary_operat/Server_Streaming/protoc" // Alias Keyword mention

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

	// Prepaing a gRPC request
	stream, err := client.ServerReply(context.TODO(), &proto.HelloRequest{SomeString: "Usman! kaise ho"})
	if err != nil {
		fmt.Println("Something error")
		return
	}

	// Now we run the infinite loop because we don't know how much request we get from the server

	count := 0
	for {
		messages, err := stream.Recv()
		if err == io.EOF { // error == End Of File operation
			break
		}
		fmt.Println("Friend Messages :- ", messages)
		time.Sleep(1 * time.Second)
		count++
	}

	err = stream.CloseSend()
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"success":       true,
		"message_count": count,
	})
}

// You can follow this to make the client request for gRPC from the  generated files :-

// type ExampleClient interface {
// 	// rpc is a keyword which is used to define the services of gRPC
// 	// stream --> Means our data is transmit in the small packets or chunks
// 	ServerReply(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[HelloResponse], error)
// }
