package main

import (
	"fmt"
	"io"
	"net"
	"strconv"
	proto "unary_operat/Client_Streaming/protoc" // Alias Keyword mention

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
	total := 0 // Total message count from the client side or user
	// while-loop
	for {
		request, err := stream.Recv()
		// When the client messages are completed
		if err == io.EOF { // EOF  --> End Of File
			return stream.SendAndClose(&proto.HelloResponse{
				Reply: strconv.Itoa(total),
			})
		}
		if err != nil {
			return err
		}

		total++
		fmt.Println(request)
	}

}
