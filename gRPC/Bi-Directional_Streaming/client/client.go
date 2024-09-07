package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	proto "unary_operat/Bi-Directional_Streaming/protoc" // Alias Keyword mention

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
	stream, err := client.ServerReply(context.TODO())
	if err != nil {
		fmt.Println("Something error :- ", err)
		return
	}

	send, receive := 0, 0
	for i := 0; i < 6; i++ {
		err := stream.Send(&proto.HelloRequest{SomeString: "The incoming request " + strconv.Itoa(i) + " from client"})
		if err != nil {
			fmt.Println("Something error :- ", err)
			return
		}
		send++
	}

	if err := stream.CloseSend(); err != nil {
		log.Fatal(err)
	}

	// Now we run the infinite loop because we don't know how much request we get from the server

	for {
		message, err := stream.Recv()
		if err == io.EOF { // io.End Of File Operation
			break
		}

		fmt.Println("Server messages :-", message.Reply)
		receive++
	}

	c.JSON(http.StatusOK, gin.H{
		"success":         true,
		"message_receive": receive,
		"message_sent":    send,
	})
}

// In _grpc.pb.go file :-

// type ExampleClient interface {
// 	// rpc is a keyword which is used to define the services of gRPC
// 	// stream --> Means our data is transmit in the small packets or chunks
// 	ServerReply(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[HelloRequest, HelloResponse], error)
// }
