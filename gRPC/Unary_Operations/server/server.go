package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	proto "unary_operat/protoc" // Alias Keyword mention

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// we have used the UnimplementedExampleServer struct because they are further used in a method
type server struct {
	proto.UnimplementedExampleServer
}

func main() {
	listener, tcpErr := net.Listen("tcp", ":5000")
	if tcpErr != nil {
		panic(tcpErr)
	}

	// Initialize the gRPC server
	srv := grpc.NewServer() // engine
	proto.RegisterExampleServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}

// implement the rpc call service
func (s *server) ServerReply(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) { // Pointer Reciever
	fmt.Println("Recieve request from client", req.SomeString)
	fmt.Println("Hello from server")
	return &proto.HelloResponse{}, errors.New("")
}
