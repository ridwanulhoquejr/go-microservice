package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/ridwanulhoquejr/go-microservice/handlers"
)

// const bindAddr = env.String("BIND_ADDRESS", false, ":8080", "Bind address for the server")

func main() {

	l := log.New(os.Stdout, "product-api ", log.LstdFlags)

	// creating an instance of our Hello struct, which implements the http.Handler interface
	ph := handlers.NewProduct(l)

	// then creating our own ServeMux
	// This ServeMux actually binding our route-path with the handler func we defined for each specific route
	sm := mux.NewRouter()

	// then, we are saying bind (handle) this specific route-path ("/") with HelloHandler (hh)
	sm.HandleFunc("/product/get", ph.GetProduct).Methods("GET")
	sm.HandleFunc("/product/create", ph.AddProduct).Methods("POST")
	sm.HandleFunc("/product/{id}", ph.UpdateProduct).Methods("PUT")

	//! alternative way to declare and register a func with route-path
	// sm.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println("Goodbye registering ServeMux with handleFunc method")
	// })

	// add some configurations in http.Server
	s := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatalf("Error while connecting to the TCP connection; %s", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	sig := <-c

	l.Println("Recieved terminate, graceful shutdown", sig)

	tc, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	s.Shutdown(tc)
	log.Println("shut down gracefully")
}
