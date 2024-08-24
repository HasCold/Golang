package main

// go mod init <module-name>
// go mod tidy  //  Install all the packages dependency related to our code

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive" // Import the primitive package
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type manager struct {
	Connection *mongo.Client // Connection is a pointer to a mongo.Client Struct we are referencing the actual memory address location becuase it is more efficient to use like this instead of copying the whole struct
	Ctx        context.Context
	Cancel     context.CancelFunc
}

var Mgr Manager

// Interface are just like your templates in which you declared function without the implementation of their actual logic.
// Interface is the collection of methods and signature
type Manager interface {
	Insert(interface{}) error
	GetAll() ([]User, error)
	DeleteData(primitive.ObjectID) error
	UpdateData(User) error
}

func connectDB() {
	uri := "localhost:27017"

	// Context is a built-in golang package which has information that required to our mongoDB , functions, handlers and routers
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 1. Context Creation: You create a context with a timeout of 10 seconds. This context will automatically cancel if the operation takes too long, preventing hanging operations.
	// 2. The client.Connect(ctx) call is redundant because mongo.Connect already establishes the connection. Additionally, client.Connect is deprecated.

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("%s%s", "mongodb://", uri))) // MongoDB Connection: mongo.Connect establishes a connection to MongoDB using the provided URI.
	if err != nil {
		fmt.Println("Error in connecting MongoDB :- ", err)
		return
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("Error in taking response from MongoDB :- ", err)
		return
	}

	fmt.Println("MongoDB successfully connected !!!")
	// Interface can hold any value regardless of their data-type e.g. :- string, array, slice, bool, integer
	Mgr = &manager{Connection: client, Ctx: ctx, Cancel: cancel} // The expression &manager{...} creates a pointer to a manager struct.
}

func Close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	defer cancel()

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(ctx)
		}
	}()
}

func init() { // Init function will mainly run before the main goroutine function becuase we are establishing the connectivity from our mongoDB
	connectDB()
}

// unique ID :- primitive.ObjectID
// MongoDB stores data in BSON (Binary JSON) format which is binary-encoded serialization format used to store and transfer data in MongoDB same like when we communicate with the client so we do mostly in JSON
type User struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"` // omitempty means if this field is avaible so OK otherwise there will be no problem
	Name  string             `bson:"name"`
	Email string             `bson:"email"`
}

func main() {
	// Insert record to mongoDB
	// Now we have matched the User struct data-type with method recievers data-type so we can call the function directly
	u := User{Name: "go Hasan", Email: "ha029292@gmail.com"}
	err := Mgr.Insert(u)
	if err != nil {
		fmt.Println("Error occur while inserting data into the database", err)
	}

	// Get all records from DB
	data, err := Mgr.GetAll()
	fmt.Println("The data we get from the GET ALL Operation :-", data, err)

	// Delete record from DB
	id := "66c982c1aee1294df3544819"
	objectID, err := primitive.ObjectIDFromHex(id) // We are converting the id string into ObjectID primitive form
	if err != nil {
		fmt.Println(err)
		return
	}

	err = Mgr.DeleteData(objectID)
	if err != nil {
		log.Fatal(err) // when critical errors encounter in the program which stops the continuation of the program so we have to log the error messages and then immediately terminates the program with a non-zero exit status code
	}

	//
	// Update the Record

	ObjectID, err := primitive.ObjectIDFromHex(id) // we are converting the id string into primitive ObjectID form
	if err != nil {
		fmt.Println(err)
	}
	u.ID = ObjectID
	u.Name = "test"
	u.Email = "test@gmail.com"
	err = Mgr.UpdateData(u)
	if err != nil {
		log.Fatal(err)
	}
}

// If want to understand more about methods in golang then look at the  Mutex_Lock_Unlock section
// Receiver types
// There are two types of receivers that are available in Go. The value receivers and the pointer receivers.

// Methods :-
// (reciver_name type)  method_name(argument_name type) (return_type)
func (mgr *manager) Insert(data interface{}) error {
	orgCollection := mgr.Connection.Database("Go_Book").Collection("Golang_Collection") // create the instance of the collection
	result, err := orgCollection.InsertOne(context.TODO(), data)

	fmt.Println("The insert data result :- ", result.InsertedID)
	return err
}

// Purpose of context.TODO()
// Placeholder Context: context.TODO() is essentially a placeholder. It’s a way of saying, “I know this function requires a context, but I don’t have one to provide right now.” It’s useful when you’re still developing your code and haven’t yet decided how to handle the context.

// Intent for Later Review: The name TODO itself implies that the developer intends to replace it later with a more appropriate context, such as one created with context.WithTimeout, context.WithCancel, or context.Background.

func (mgr *manager) GetAll() (data []User, err error) { // mgr hold a reference to the manager struct or mgr is a pointer to the manager struct
	orgCollection := mgr.Connection.Database("Go_Book").Collection("Golang_Collection")

	// Pass these options to the find method
	findOptions := options.Find()
	// M is an unordered representation of a BSON document. This type should be used when the order of the elements does not matter. This type is handled as a regular map[string]interface{} when encoding and decoding.
	cur, err := orgCollection.Find(context.TODO(), bson.M{}, findOptions) // return the cursor and error

	// For loop runs until there is a data inside the cursor
	for cur.Next(context.TODO()) {
		var decodeUser User
		err := cur.Decode(&decodeUser) // bind the decodeUser variable
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, decodeUser) // Nil Slice, This means that data is automatically initialized to its zero value when the function is called. In Go, the zero value of a slice (like []User) is nil, but this doesn't cause issues when using append.
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	// Close the cursor once finished
	cur.Close(context.TODO())
	return data, nil
}

func (mgr *manager) DeleteData(id primitive.ObjectID) error {
	orgCollection := mgr.Connection.Database("Go_Book").Collection("Golang_Collection") // Create the instance of DB
	// bson.D{} -->> Order representation of Data
	filter := bson.D{{"_id", id}, {"name", "go Hasan"}}

	_, err := orgCollection.DeleteOne(context.TODO(), filter)
	return err
}

func (mgr *manager) UpdateData(data User) error {
	orgCollection := mgr.Connection.Database("Go_Book").Collection("Golang_Collection")

	// D is an ordered representation of a BSON document.
	filter := bson.D{{"_id", data.ID}}
	update := bson.D{{"$set", data}}

	updateResult, err := orgCollection.UpdateOne(context.TODO(), filter, update)
	fmt.Println("The updated result from mongoDB :- ", updateResult)

	return err
}

// ----------------------------------------------------------------------------------------------------------------------
// Mgr = &manager{Connection: client, Ctx: ctx, Cancel: cancel}

// 1. Pointer to Struct (&manager): The expression &manager{...} creates a pointer to a manager struct.
// 2. Interface Satisfaction: The *manager type (a pointer to the manager struct) implements the Manager interface because it has all the methods required by the Manager interface (Insert, GetAll, DeleteData, UpdateData).
// 3. Dynamic Typing in Interfaces: In Go, an interface can hold any type that implements the interface's methods. When you assign &manager{...} to Mgr, the Mgr variable now holds a reference to the manager struct, but it is treated as the Manager interface type.

//
// ------------------------------------------------------------------------------------------
// BSON (Binary JSON) is a binary-encoded serialization format used to store and transfer data in MongoDB. It stands for Binary JSON, and it extends the JSON format by adding support for more data types and enabling efficient data storage and retrieval.

// Key Features of BSON:
// Binary Encoding: Unlike JSON, which is a text-based format, BSON is binary. This makes it more compact and faster to encode and decode, which is crucial for performance in database operations.

// Rich Data Types: BSON supports additional data types that are not available in JSON, such as:
// ObjectID: A unique identifier for documents, which is a 12-byte binary value.
// Date: Stores date and time in a format that is easy to compare and sort.
// Binary Data: Allows storing binary data like images or files.
// Decimal128: Supports high-precision decimal values.
// Int32 and Int64: Support for both 32-bit and 64-bit integers.
// Double: Support for floating-point numbers.
