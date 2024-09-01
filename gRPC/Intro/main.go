package main

// -------------------------- RPC ---------------------------------
// A RPC is a method to form of Client-Server Communication that uses function call rather than a usual HTTP call.
// It uses IDL (Interface Definition Language) as a form of contract on functions to be called and on the data-type. Basically you can achieve the data-abstraction.

// --------------------- gRPC --------------------------------
// What is gRPC :- A gRPC is a powerful framework for working with Remote Procedure Calls
// So gRPC is technically not a new concept. Rather it was adopted from this old technique and improved upon, making it very popular in just the span of 6-7 years. gRPC is the advanced version of rpc protocol.
// "g" in the gRPC represent the advanced version of gRPC protocol naming.

//
// --------------------- gRPC Popular --------------------------------
// 1. Abstraction is easy (it's a function call)
// 2. It is very performant
// 3. HTTP calls are often confusing, so this make it easier
// 4. Microservices will often be running several services in different programing lang :- Golang, Java etc.

// ------------------------- gRPC Architecture --------------------------------
// In gRPC, the communication will be establish in the buffer form

// gRPC FILE :-
// 1. All gRPC file have this .proto extension.
// 2. Stub represents the proto file.
// 3. Request/Response cycle will be held in the buffer medium.

// message User {  // User/Client will follow this structure
// 	Id string
// 	Name string
// }

// Services will be exposed like this way :-
// services Rpc{
// 	service RegisterUser(User) return(User);
// }
// --------------------------------------------------------------------------------------------

//
// ------------------------- Data Transfer Medium --------------------------------
// 1. xml, 1991 - 1999 HTTP 1.0    --->>  Mainly Used in the SOAP, Tag form <tag></tag> ; Are heavy-weight
// 2. json, 1999 - 2015 HTTP 1.1   --->>  Mainly Used in the HTTP, Object form {"user":"Hasan"} ; Are lighter than xml
// 3. buffer, 2015... HTTP 2.0     --->>  Mainly Used in the gRPC, [bytes] / Array of bytes ; Are more lighter than json and it consumes low memory to transport and the code works very efficiently
// --------------------------------------------------------------------------------------------

//
// ------------------------- gRPC offer --------------------------------
// What else does gRPC offer ?
// 1. MetaData (header)  --->>> The information of the data
// 2. Streaming
// --------------------------------------------------------------------------------------------

//
// ------------------------- Types of gRPC --------------------------------
// 1. Unary operation  == Client Request => <= Server Response
// 2. Server streaming == Client <= Data comes from server in the form of junks, bytes, buffer or streaming
// 3. Client streaming == Server <= Client
// 4. Full duplex bidirectional communication or streaming / In the sense of Web-Sockets
// --------------------------------------------------------------------------------------------

//
// Load Balancing :- Nginx

// ---------------------- Command --------------------------
// Command to run the protocol buffer file :- protoc --go-grpc_out=. --go_out=. *.proto
// This command will generate two files one for modal and second for services

func main() {

}
