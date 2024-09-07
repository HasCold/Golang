package main

import (
	"fmt"
	"net"
	"time"
	proto "unary_operat/Server_Streaming/protoc" // Alias Keyword mention

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
func (s *server) ServerReply(req *proto.HelloRequest, stream proto.Example_ServerReplyServer) error {
	fmt.Println(req.SomeString)
	time.Sleep(5 * time.Second)

	friendReply := []*proto.HelloResponse{
		{Reply: "Hello"},
		{Reply: "Hasan !"},
		{Reply: "Ha ma thek hu"},
		{Reply: "Baki tm batao"},
	}

	for _, msg := range friendReply {
		err := stream.Send(msg)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}

// you will follow this to take an argument from the generate files :-

// type ExampleServer interface {
// 	// rpc is a keyword which is used to define the services of gRPC
// 	// stream --> Means our data is transmit in the small packets or chunks
// 	ServerReply(*HelloRequest, grpc.ServerStreamingServer[HelloResponse]) error
// 	mustEmbedUnimplementedExampleServer()
// }
