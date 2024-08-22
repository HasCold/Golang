package main

// Docs :- https://github.com/mongodb/mongo-go-driver

// go mod init <module-name>
// go mod tidy  //  Install all the packages dependency related to our code

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// The manager struct holds the MongoDB connection and context information. The Connection is a pointer to the MongoDB client, Ctx is the context used for operations, and Cancel is the cancel function to clean up the context.
type manager struct {
	Connection *mongo.Client // Connection is a pointer to a mongo.Client Struct we are referencing the actual memory address location becuase it is more efficient to use like this instead of copying the whole struct
	Ctx        context.Context
	Cancel     context.CancelFunc
}

// The context package is a powerful tool to manage operations like timeouts, cancelation, deadlines etc. Among these operations , context with timeout is mainly used when we want to make external request, such as a network request or a database request.

var Mgr manager

func connectDB() {
	// MongoDB by-default port :- 27018, 27019 on that will run the port
	uri := "localhost:27017" // 127.0.0.1:27017

	// Golang in-built package context
	// In Context has some information that may be required to our mongoDB or functions or handlers or routers
	// Below we are doing the cancellation operations of the parent context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // This ensures that the cancel function is called to release the resources associated with the context, even if an early return occurs due to an error.

	// The client.Connect(ctx) call is redundant because mongo.Connect already establishes the connection. Additionally, client.Connect is deprecated.
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("%s%s", "mongodb://", uri))) // The correct URI string should look like mongodb://localhost:27017, which is what the fmt.Sprintf line will produce.
	if err != nil {
		fmt.Println(err)
		return
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("Error while taking response from MongoDB :-", err)
		return
	}

	Mgr = manager{Connection: client, Ctx: ctx, Cancel: cancel}
	fmt.Println("MongoDB successfully connected !!!")
}

func Close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	// CancelFunc to cancel the context
	defer cancel() // defer will execute this at the end of all nearby function execution or simply make some delay

	// Client provides a method to Close a mongoDB Connection
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func main() {
	connectDB()
}

// ------------------------------------- Explanation Of connectDB function ------------------------------------

// 1. Context Creation: You create a context with a timeout of 10 seconds. This context will automatically cancel if the operation takes too long, preventing hanging operations.

// 2. MongoDB Connection: mongo.Connect establishes a connection to MongoDB using the provided URI.

// 3. Ping MongoDB: You use client.Ping to verify that the connection to MongoDB is successful.

// 4. Assign to Global Manager: You store the client and context in the Mgr global variable.

//
// ------------------------------------- Explanation Of Close function ------------------------------------

// - This function is responsible for properly closing the MongoDB connection and canceling the context when you are done with the database. It ensures that resources are cleaned up.

// Error Handling:
// The if err := client.Disconnect(ctx); err != nil { ... } part checks if an error occurred during the disconnection.
// If an error does occur (e.g., if the disconnection fails), the panic(err) line is executed. panic will stop the normal execution of the program and print the error, which is useful for debugging critical issues that should not be ignored.

// - Resource Cleanup: This code ensures that the MongoDB client is always disconnected when the surrounding function (Close) finishes, regardless of whether the function exits normally or due to an error. This is essential for freeing up resources and preventing memory leaks or open connections that could exhaust the database connection pool.

// - Panic on Error: The panic(err) part indicates that if something goes wrong during disconnection, it’s considered a serious problem, and the program should not continue running. This is often used during development to catch unexpected issues.
