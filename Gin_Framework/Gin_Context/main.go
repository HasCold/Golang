package main

// What is Context
// This context contains all the information about the request that the handler might need to process it
// It is easy to pass request-scoped values, cancellation signals and deadline across API boundaries to all the goroutines involved in handling the request.

// :9090/api/:path?n=go&email=save
// like header , cookies etc

// json {}, [{}, {}]
// XML <tag></tag>

// --------------------- Usage of Context ---------------------------
// err := c.Bind(&obj)  // Pass a reference of object // Accepts a json or xml body
// err := c.BindQuery(&obj)  // Accepts a query parameter like in this  :- form:"n"
// err := c.BindXML(&obj)  // Accepts a xml data <tag></tag>
// err := c.BindYAML(&obj)
// err := c.BindHeader(&obj)  // Accept this thing  :- header:"header"
// err := c.BindJSON(&obj)  // Accept json data
// c.Header("user-id", "7237232398")  // Manually handle the header like we pass the key : user-id
// c.SetCookie(name, value: string, maxAge, int, path, domain: string, secure, httpOnly: bool)  // set the cookie
// val, err := c.Cookie("name")  // get the value from a cookie by their key name
// err := c.SaveUploadedFile(file.json, "/path/where/to/save")  // Save the files via Gin Context
// form := c.MultipartForm()  // We can save the form-data or files
// key := c.PostForm("key", "default-value")
// id := c.Query("id")	// /route/path?id=12
// id := c.Param("id")  // /route/2/:id  => 1, 2, 3
// name := c.DefaultQuery("name", "jack")  // suppose if you don't give anything in the query parameter name so it will select by-default// :9090/api/:path?name=&email=save

// c.GetFloat64("myKey")
// c.Set("myKey", "myValue")  // Set the key in req object
// c.Get("myKey")
// c.MustGet("myKey")  // Key must be obtained in any way
// c.GetString("myKey")  // Key data type will be convert to the string

// For Authorization (JWT) :-
// c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})  // send message from gin.H and  http.StatusInternalServerError if authorization failed
// c.IsAborted()  -->> Returns bool value if request will abort
// c.AbortWithStatus(http.StatusBadRequest) //  we can only abort just with the status
// c.Abort()  -->> Simply abort to the unauthorized person like hacker without sending to the message or status code

//
// When users are authorized so we have to give response in this way :-
// c.HTML(
// 	http.StatusOK, "index.html", gin.H{
// 		"message": "This is a message"
// 	}
// )

// c.JSON(
// 	http.StatusOK, gin.H{
// 		"message": "this is a message"
// 	}
// )

// c.XML(
// 	http.StatusOK, gin.H{
// 		"message": "this is a message"
// 	}
// )

// c.YAML(
// 	http.StatusOK, gin.H{
// 		"message": "this is a message",
//		"data": value
// 	}
// )

type api struct {
	Name  string `json:"name" form:"n" header:"header" binding:"required"`
	Email string `json:"email"`
}

func main() {

}
