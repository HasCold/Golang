package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"
	proto "unary_operat/Streaming_BigFiles/protoc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client proto.StreamUploadClient

func main() {
	// 	Connection to internal grpc server
	//  making the internal tcp connection
	conn, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	// Prompt for file path
	fmt.Println("Enter the file path : ")
	var path string
	fmt.Scanln(&path)

	// Checks if the file exist and accessible
	if _, e := os.Stat(path); os.IsNotExist(e) {
		log.Printf("The file %s does not exist.\n", path)
		return
	} else if e != nil {
		log.Printf("Error accessing file %s: %v\n", path, err)
		return
	}

	client = proto.NewStreamUploadClient(conn)

	mb := 1024 * 1024 * 2 // 2 MB file
	uploadStramFile(path, mb)
}

func uploadStramFile(path string, batchSize int) {
	t := time.Now() // To identify the time taken by current function
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error in Accessing File :-", err)
		return
	}

	buff := make([]byte, batchSize) // reading data in fixed-size chunks (like in file or network operations), ensuring the buffer is of a fixed size for each read operation.
	batchNumber := 1

	stream, err := client.Upload(context.TODO())
	if err != nil {
		panic(err)
	}

	for {
		num, err := file.Read(buff)
		if err == io.EOF {
			fmt.Println(err)
			return
		}

		chunk := buff[:num] // buff[:num]: create a new slice from buff starting from index 0 and up to index num (but not including num).
		if e := stream.Send(&proto.UploadRequest{FilePath: path, Chunks: chunk}); e != nil {
			fmt.Println(e)
			return
		}
		log.Printf("Sent - batch #%v - size - %v \n", batchNumber, len(chunk))
		batchNumber += 1
	}

	res, er := stream.CloseAndRecv()
	if er != nil {
		fmt.Println(er)
		return
	}

	fmt.Println("Received response")
	fmt.Println(time.Since(t))
	log.Printf("FileSize-%v, Bytes-%v", res.GetFileSize(), res.GetMessage())
}
