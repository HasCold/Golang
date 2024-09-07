package main

import (
	"io"
	"log"
	"net"
	"os"
	proto "unary_operat/Streaming_BigFiles/protoc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedStreamUploadServer
}

func main() {
	listener, tcpErr := net.Listen("tcp", ":9000")
	if tcpErr != nil {
		panic(tcpErr)
	}

	// Initialize the gRPC server engine
	srv := grpc.NewServer()
	proto.RegisterStreamUploadServer(srv, &server{})
	reflection.Register(srv)

	if er := srv.Serve(listener); er != nil {
		panic(er)
	}
}

func (s *server) Upload(stream proto.StreamUpload_UploadServer) error {
	var fileBytes []byte
	var fileSize int64 = 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}

		chunks := req.GetChunks() // this is the short-hand method to declare a variable and compiler will judges their data-type at a compile time
		fileBytes = append(fileBytes, chunks...)
		fileSize += int64(len(chunks))
	}

	file, err := os.Create("./test.json")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	_, fileError := file.Write(fileBytes)
	if fileError != nil {
		log.Fatal(fileError)
	}

	return stream.SendAndClose(&proto.UploadResponse{FileSize: fileSize, Message: "File Written Successfully"})
}
