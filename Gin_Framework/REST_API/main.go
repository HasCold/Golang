package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive" // Import the primitive package
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// BSON (Binary JSON) format is a binary encoded serialization format used to store and trasnfer data in mongoDB. BSON is the extended version of JSON which is very Rich data-type that supports in mongoDB

type Data struct {
	// json deals with the frontend api data and bson is for the mongoDB
	ID    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"` // omitempty means if this field is coming from frontend so OK we will consume it otherwise there will be no problem mongoDB will generate the ObjectID for us
	Name  string             `json:"name" bson:"name"`
	Email string             `json:"email" bson:"email"`
}

type ApiData struct {
	ID    string `json:"id,omitempty" bson:"id,omitempty"` // omitempty means if this field is coming from frontend so OK we will consume it otherwise there will be no problem mongoDB will generate the ObjectID for us
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
}

type manager struct {
	Connection *mongo.Client
	Ctx        context.Context
	Cancel     context.CancelFunc
}

// golang basically provide us the powerful and flexible dynamic type interfacing
// Interface can hold any value regardless of the data-type
var Mgr Manager

type Manager interface {
	Insert(interface{}) error
	GetAll() ([]Data, error)
	DeleteData(primitive.ObjectID) error
	UpdateData(Data) error
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
	// In Go, an interface can hold any type that implements the interface's methods.
	// Interface can hold any value regardless of their data-type e.g. :- string, array, slice, bool, integer
	Mgr = &manager{Connection: client, Ctx: ctx, Cancel: cancel} // The expression &manager{...} creates a pointer to a manager struct.
}

func Close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	defer cancel()

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func init() { // Init function will mainly run before the main goroutine function becuase we are establishing the connectivity from our mongoDB
	connectDB()
}

func main() {
	r := gin.Default() // Default returns the gin engine instance

	r.POST("/post-data", insertData)
	r.GET("/get-data", getAll)
	r.PUT("/put-data", updateData)
	r.DELETE("/del-data", delData)

	r.Run(":5000")
}

// Passing *gin.Context as a pointer is efficient because it avoids copying the entire context structure, which could be large. It also ensures that any modifications to the context (e.g., setting a response header) are reflected across all parts of the code handling the request.
func insertData(c *gin.Context) { // c is a pointer to a gin Context struct
	var d Data

	fmt.Println("The coming JSON obj :-", d)
	err := c.BindJSON(&d) // passing the reference of json data
	if err != nil {
		fmt.Println(err)
		return
	}

	err = Mgr.Insert(d)
	if err != nil {
		log.Fatal(err) // when critical errors encounter in the program which stops the continuation of the program so we have to log the error messages and then immediately terminates the program with a non-zero exit status code
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": d,
	})
}

func getAll(c *gin.Context) {
	respData, err := Mgr.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    respData,
	})
}

func delData(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "id parameter is required",
		})
	}

	objectID, err := primitive.ObjectIDFromHex(id) // We are converting the id string into ObjectID primitive form
	if err != nil {
		fmt.Println(err)
		return
	}

	err = Mgr.DeleteData(objectID)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Deleted Successfully !",
	})
}

func updateData(c *gin.Context) {
	var reqData ApiData
	var finalData Data

	fmt.Println("The coming JSON obj :-", reqData)
	err := c.BindJSON(&reqData) // passing the reference of JSON data
	if err != nil {
		fmt.Println("Error something incoming request")
		return
	}

	objectID, err := primitive.ObjectIDFromHex(reqData.ID) // Convert the string ID into primitive Object ID
	if err != nil {
		fmt.Println("Error something incoming request")
		return
	}

	finalData.ID = objectID
	finalData.Name = reqData.Name
	finalData.Email = reqData.Email

	err = Mgr.UpdateData(finalData)
	if err != nil {
		log.Fatal("Error in Updating Data :- ", err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Data Updated Successfully",
	})
}

// Method Creation
// (reciever_name reciever_type)
func (mgr *manager) Insert(data interface{}) error {
	orgCollection := mgr.Connection.Database("REST_GO").Collection("golang") // create an instance of DB after making the connection
	res, err := orgCollection.InsertOne(context.TODO(), data)                // when we don't know what context should we passed so then context.TODO()

	fmt.Println("The response is :- ", res)
	return err
}

func (mgr *manager) GetAll() (data []Data, err error) {
	orgCollection := mgr.Connection.Database("REST_GO").Collection("golang")
	// Pass these options to the find method
	findOption := options.Find()

	cur, err := orgCollection.Find(context.TODO(), bson.M{}, findOption)

	// For loop runs until there is a data inside the cursor
	for cur.Next(context.TODO()) {
		var d Data
		err := cur.Decode(&d) // Bind the data with the d struct
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, d) // Nil Slice, This means that data is automatically initialized to its zero value when the function is called. In Go, the zero value of a slice (like []User) is nil, but this doesn't cause issues when using append.
	}

	if err = cur.Err(); err != nil {
		return nil, err
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return data, nil
}

func (mgr *manager) DeleteData(id primitive.ObjectID) error {
	orgCollection := mgr.Connection.Database("REST_GO").Collection("golang")

	filter := bson.D{{"_id", id}} // D is an ordered representation of a BSON document. This type should be used when the order of the elements matters, such as MongoDB command documents.
	_, err := orgCollection.DeleteOne(context.TODO(), filter)
	return err
}

func (mgr *manager) UpdateData(data Data) error {
	orgCollection := mgr.Connection.Database("REST_GO").Collection("golang")

	filter := bson.D{{"_id", data.ID}}
	update := bson.D{{"$set", data}}

	res, err := orgCollection.UpdateOne(context.TODO(), filter, update)
	fmt.Println("The updated result is :- ", res)

	return err
}
