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

func (p *Product) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		p.GetProduct(w, r)
		return
	}
	if r.Method == http.MethodPost {
		p.AddProduct(w, r)
		return
	}

	// catch all
	WriteError(
		w,
		http.StatusMethodNotAllowed,
		"Method not allowed",
	)
}

// Get Products
func (p *Product) GetProduct(w http.ResponseWriter, r *http.Request) {

	lp := data.GetProducts()
	err := lp.ToJSON(w)

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

	prod := &data.Product{}
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
	data.AddProduct(prod)
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
