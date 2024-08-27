package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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

type jwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

// jwt adds email as a claim to the token
type JwtClaim struct {
	Email string
	jwt.RegisteredClaims
}

func (j *jwtWrapper) generateToken(email string) (signedToken string, err error) { // j is a pointer to the jwtWrapper struct
	// Creates a pointer to a JwtClaim which further holds the reference
	claims := &JwtClaim{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(j.ExpirationHours) * time.Hour)),
			Issuer:    j.Issuer,
		},
	}

	// Now we are using SHA-256 method
	// The claims created earlier are attached to this token.
	// NewWithClaims creates a new Token with the specified signing method and claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// SignedString creates and returns a complete, signed JWT. The token is signed using the SigningMethod specified in the token.
	signedToken, err = token.SignedString([]byte(j.SecretKey))
	if err != nil {
		fmt.Println("Error in signing token :- ", err)
		return
	}
	return
}

// Validate Token validate the JWT token
func (j *jwtWrapper) validateToken(signedToken string) (claims *JwtClaim, err error) {
	token, err := jwt.ParseWithClaims( //Parses the signed JWT (signedToken) with the expectation that it contains JwtClaim claims.
		signedToken,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) { // The function passed as a third argument returns the secret key used to sign the token, allowing the parsing function to verify the signature.
			return []byte(j.SecretKey), nil
		},
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	// we are just checking the token time in this function
	claims, ok := token.Claims.(*JwtClaim)
	if !ok {
		err = errors.New("Couldn't parse the claims")
		return
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		err = errors.New("JWT is expired")
		return
	}
	return
}

// Auhz authenicate token and authorize user
func Authz() gin.HandlerFunc {
	// Closure function
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			c.JSON(403, "No authorization header provided")
			c.Abort()
			return
		}

		extractedToken := strings.Split(clientToken, "Bearer ")

		if len(extractedToken) == 2 {
			clientToken = strings.TrimSpace(extractedToken[1])
		} else {
			c.JSON(400, "Incorrect format of Authorization Token")
			c.Abort()
			return
		}

		jwtWrapper1 := jwtWrapper{
			SecretKey: "esfsdfkpskodkf234234243243",
			Issuer:    "admin",
		}

		// jwtWrapper can invoked the method directly becuase it matches the reciever_type of method
		claims, err := jwtWrapper1.validateToken(clientToken)
		if err != nil {
			c.JSON(401, err.Error())
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Next() // Next should be used only inside middleware. It executes the pending handlers in the chain inside the calling handler

	}
}

func main() {
	r := gin.Default() // Default returns the gin engine instance

	r.GET("/get-data", getAll)
	r.GET("/token", getToken)
	r.Use(Authz()) // Pass the middleware in Use method
	// Below are the three api's will authorize first from the middleware
	r.POST("/post-data", insertData)
	r.PUT("/put-data", updateData)
	r.DELETE("/del-data", delData)

	r.Run(":5000")
}

func getToken(c *gin.Context) {
	// jwtWrapper1 can invoked the method directly becuase it matches the reciever_type of method
	jwtWrapper1 := jwtWrapper{
		SecretKey:       "esfsdfkpskodkf234234243243",
		Issuer:          "admin",
		ExpirationHours: 48,
	}

	signedToken, err := jwtWrapper1.generateToken("test@gmail.com")
	if err != nil {
		errors.New(err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"token":   signedToken,
	})

}

// Passing *gin.Context as a pointer is efficient because it avoids copying the entire context structure, which could be large. It also ensures that any modifications to the context (e.g., setting a response header) are reflected across all parts of the code handling the request.
func insertData(c *gin.Context) { // c is a pointer to a gin Context struct
	var d Data

	if c.Request.Method != "POST" {
		c.JSON(405, gin.H{
			"success": false,
			"message": "Request method is invalid",
		})
		return
	}

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
	if email, ok := c.Get("email"); ok {
		fmt.Println("The email get from response :- ", email)
	}

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

	fmt.Println("The response is :- ", *res)
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
