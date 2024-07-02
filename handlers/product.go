package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ridwanulhoquejr/go-microservice/data"
)

type Product struct {
	l *log.Logger
}

type Error struct {
	Code    int    `json:"code"`
	Details string `json:"details"`
}

func NewProduct(l *log.Logger) *Product {
	return &Product{
		l: l,
	}
}

//! these below blocks are needed for built-in net/http
// func (p *Product) ServeHTTP(w http.ResponseWriter, r *http.Request) {

// 	if r.Method == http.MethodGet {
// 		p.GetProduct(w, r)
// 		return
// 	}
// 	if r.Method == http.MethodPost {
// 		p.AddProduct(w, r)
// 		return
// 	}

// 	// catch all
// 	WriteError(
// 		w,
// 		http.StatusMethodNotAllowed,
// 		"Method not allowed",
// 	)
// }

// Get Products
func (p *Product) GetProduct(w http.ResponseWriter, r *http.Request) {

	// get the dummy products
	lp := data.GetProducts()
	// conver Go-Struct into JSON
	err := data.WriteJSON(w, &lp)

	if err != nil {
		WriteError(
			w,
			http.StatusMethodNotAllowed,
			"Internal Server Error",
		)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// create product
func (p *Product) AddProduct(w http.ResponseWriter, r *http.Request) {

	p.l.Println("Handle POST method")

	// a ref of our Product struct for recieving the args
	prod := &data.Product{}
	// convert it to the JSON format
	err := prod.FromJSON(r)

	if err != nil {
		WriteError(
			w,
			http.StatusUnprocessableEntity,
			"Payload cannot be processed",
		)
		return
	}
	p.l.Printf("Prod: %#v", prod)

	// add the incoming req into our dummy Product
	data.AddProduct(prod)

	// sending back the product just created as a response
	err = data.WriteJSON(w, prod)

	if err != nil {
		WriteError(
			w,
			http.StatusMethodNotAllowed,
			"Internal Server Error",
		)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func WriteError(w http.ResponseWriter, code int, details string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(
		Error{
			Code:    code,
			Details: details,
		},
	)
}
