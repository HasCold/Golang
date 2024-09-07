package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"
	proto "unary_operat/Bi-Directional_Streaming/protoc" // Alias Keyword mention

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// we have used the UnimplementedExampleServer struct because they are further used in a method
type server struct {
	proto.UnimplementedExampleServer
}

func main() {
	listener, tcpError := net.Listen("tcp", ":5000")
	if tcpError != nil {
		panic(tcpError)
	}

	// Initialize the gRPC server
	srv := grpc.NewServer()
	proto.RegisterExampleServer(srv, &server{})
	reflection.Register(srv)

	if err := srv.Serve(listener); err != nil {
		panic(err)
	}
}

// You have to make all the functions and methods as according to the protobuf generated files :- Client_grpc.pb.go
func (s *server) ServerReply(stream proto.Example_ServerReplyServer) error {
	for i := 0; i < 10; i++ {
		err := stream.Send(&proto.HelloResponse{Reply: "Message" + strconv.Itoa(i) + "from server"})
		if err != nil {
			return errors.New("unable to send data from server")
		}
	}

	for {
		req, err := stream.Recv()
		if err == io.EOF { // err == End Of File Operation
			break
		}
		fmt.Println("Client messages :-", req.SomeString)
	}

	return nil
}

// In _grpc.pb.go file :-

// type ExampleServer interface {
// 	// rpc is a keyword which is used to define the services of gRPC
// 	// stream --> Means our data is transmit in the small packets or chunks
// 	ServerReply(grpc.BidiStreamingServer[HelloRequest, HelloResponse]) error
// 	mustEmbedUnimplementedExampleServer()
// }
