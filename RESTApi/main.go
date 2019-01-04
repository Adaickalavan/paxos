package main

import (
	"database"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// Hooks that may be overridden for testing
var inputReader io.Reader = os.Stdin
var outputWriter io.Writer = os.Stdout

// Initialize database
var product = database.Product{}

func main() {

	// Connect to database
	product.Server = os.Getenv("MONGO_PORT")
	product.DatabaseName = os.Getenv("DATABASE_NAME")
	product.CollectionName = os.Getenv("COLLECTION_NAME")
	product.UserName = os.Getenv("MONGO_USERNAME")
	product.Password = os.Getenv("MONGO_PASSWORD")
	product.Session = product.Connect()
	defer product.Session.Close()

	// Ensure database index is unique
	product.EnsureIndex([]string{"digest"})

	if err := run(); err != nil {
		log.Fatal(err.Error())
	}
}

// run runs the server
func run() error {
	httpAddr := os.Getenv("LISTENING_ADDR")
	mux := makeMuxRouter()
	s := &http.Server{
		Addr:           ":" + httpAddr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Listening on ", httpAddr)
	if err := s.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
