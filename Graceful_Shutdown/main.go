package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() // Default returns a gin engine instance
	router.GET("/", func(c *gin.Context) {
		time.Sleep(10 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server") // after 10 seconds this will return a response in string format
	})

	srv := &http.Server{
		Addr:    ":8080", // Port Number
		Handler: router,
	}

	// Goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %v \n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds
	quit := make(chan os.Signal) // This channel is specifically to grab the signals

	// syscall.SIGINT  -->> when presses Ctrl + C
	// syscall.SIGTERM -->> Other system termination listen
	// How to capture signals
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// when the server listen these signals (syscall.SIGINT, syscall.SIGTERM) so it will shutdown after 10 seconds
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown", err)
	}

	// Catching ctx.Done(). timeout of 5 seconds
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds")
	}

	log.Println("Server exiting ...")

}
