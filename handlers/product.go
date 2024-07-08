package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

// update product
func (p *Product) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// take the id from the request-payload
	vars := mux.Vars(r)["id"]
	id, err := strconv.Atoi(vars)

	if err != nil {
		WriteError(
			w,
			http.StatusBadRequest,
			"Invalid id given",
		)
		return
	}

	prod := &data.Product{}
	err = prod.FromJSON(r)

	if err != nil {
		WriteError(
			w,
			http.StatusBadRequest,
			"Unable to unmarshal the JSON object",
		)
		return
	}
	err = data.UpdateProduct(id, prod)
	if err != nil {
		WriteError(
			w,
			http.StatusNotFound,
			"The product not found for the given id",
		)
		return
	}

	// get all the products
	// ps := data.GetProducts()
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
	//? no longeer to use w.Header(), we use JSONMiddleware for this
	// w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(
		Error{
			Code:    code,
			Details: details,
		},
	)
}
