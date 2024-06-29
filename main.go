package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ridwanulhoquejr/go-microservice/handlers"
)

func main() {

	l := log.New(os.Stdout, "product-api ", log.LstdFlags)

	// creating an instance of our Hello struct, which implements the http.Handler interface
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	// then creating our own ServeMux
	// This ServeMux actually binding our route-path with the handler func we defined for each specific route
	sm := http.NewServeMux()

	// then, we are saying bind (handle) this specific route-path ("/") with HelloHandler (hh)
	sm.Handle("/hello", hh)
	sm.Handle("/goodbye", gh)

	//! alternative way to declare and register a func with route-path
	// sm.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println("Goodbye registering ServeMux with handleFunc method")
	// })

	// then passing our own serveMux to the listenAndServe to handle which handleFunc would be executed
	err := http.ListenAndServe(":8080", sm)

	if err != nil {
		log.Printf("Error while connecting to the TCP connection; %s", err)
	}
}
