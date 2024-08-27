package main

// For Google OAuth :-
// Link :- https://console.cloud.google.com/home/

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
		return
	}
}

func cspMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header(
			"Content-Security-Policy", "default-src 'self'; script-src 'self' https://apis.google.com; style-src 'self' https://fonts.googleapis.com; img-src 'self' data:; connect-src 'self' https://accounts.google.com; font-src 'self' https://fonts.gstatic.com; frame-src 'self' https://accounts.google.com",
		)
		c.Next()
	}
}

func main() {
	ClientID := os.Getenv("CLIENT_ID")
	ClientSecret := os.Getenv("CLIENT_SECRET")

	key := "Secret-session-key"
	maxAge := 86400 * 30 // 24 * 60 * 60
	isProd := false

	// Create a cookie based store for the session
	store := cookie.NewStore([]byte(key))
	store.Options(sessions.Options{
		MaxAge:   maxAge,
		Path:     "/",
		HttpOnly: isProd,
	})

	gothic.Store = store

	callBackURL := "http://localhost:5000/auth/google/callback"
	goth.UseProviders(
		google.New(ClientID, ClientSecret, callBackURL, "email", "profile"), // In scope mean we are telling give us the email and profile
	)

	r := gin.Default() // default returns a gin engine instance

	r.Use(cspMiddleware())
	// This path also registered on to the google developer console to get back the response :- /auth/{provider}/callback
	r.GET("/auth/google/callback", signUp)
	r.GET("/auth/google", func(c *gin.Context) {
		gothic.CompleteUserAuth(c.Writer, c.Request)
	})

	r.GET("/", renderHTML)

	log.Println("Server running on port : 5000")
	log.Fatal(http.ListenAndServe(":5000", r)) // alternative way := r.Run()
}

func signUp(c *gin.Context) {
	// c.Writer is used as the http.ResponseWriter, which is the standard way in Go to write HTTP responses. Similarly, c.Request provides access to the incoming HTTP request.
	OUser, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		log.Printf("Error completing user auth: %v", err)
		c.HTML(http.StatusInternalServerError, "index.html", gin.H{
			"error": "Authentication failed",
		})
		return
	}

	// --------------------------------------------------------------------------------------------------------------
	// The template package in Go is part of the standard library and is used for rendering HTML templates. In your signUp function, you're using it to render an HTML page (index.html) and pass data (OUser) to it.

	t, err := template.ParseFiles("templates/index.html") // This function reads the HTML file from the specified path and parses it into a *template.Template object. The ParseFiles function can accept multiple file paths, and it returns a pointer to the template and an error.
	if err != nil {
		log.Fatalf("Error parsing template %v", err)
	}

	err = t.Execute(c.Writer, OUser) // we have bind the data in HTML template or HTML pages
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	// c.HTML(http.StatusOK, "./templates/success.html", gin.H{
	// 	"UserID": OUser.UserID,
	// 	"Email":  OUser.Email,
	// 	"Name":   OUser.Name,
	// })
}

func renderHTML(c *gin.Context) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatalf("Error parsing template %v", err)
	}

	err = t.Execute(c.Writer, false)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	// c.HTML(http.StatusOK, "./templates/index.html", gin.H{
	// 	"user": nil,
	// })
}
