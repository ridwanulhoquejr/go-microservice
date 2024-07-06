package data

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Product struct
type Product struct {
	Id               int     `json:"id"`
	Name             string  `json:"name"`
	Description      string  `json:"description"`
	SKU              string  `json:"sku"`
	Price            float32 `json:"price"`
	CreationTime     string  `json:"-"`
	ModificationTime string  `json:"-"`
	DeletionTime     string  `json:"-"`
}

// declaring a custom type for having a method like TOJSON
type Products []*Product

// dummy data for product sturct
var productList = Products{
	{
		Id:               1,
		Name:             "product-1",
		Description:      "This is product-1 description",
		SKU:              "#121234",
		Price:            99.99,
		CreationTime:     time.Now().UTC().String(),
		ModificationTime: time.Now().UTC().String(),
	},
	{
		Id:               2,
		Name:             "product-2",
		Description:      "This is product-2 description",
		SKU:              "#121235",
		Price:            59.99,
		CreationTime:     time.Now().UTC().String(),
		ModificationTime: time.Now().UTC().String(),
	},
}

// for accessing the productList from other packages such as handlers
func GetProducts() Products {
	return productList
}

// ? interface for common ToJSON parsing
// idea is to parse the JSON from go-struct of a single or slice of products
// for that i wrote an interface, and implement this with (*Product) and (*Products)
// and a seperate function for core functionlity which is return by the implemented method
type JSON interface {
	ToJSON(w http.ResponseWriter) error
}

// ToJSON method for a single Product
func (p *Product) ToJSON(w http.ResponseWriter) error {
	return encodeJSON(w, p)
}

// ToJSON method of Products
func (ps *Products) ToJSON(w http.ResponseWriter) error {
	return encodeJSON(w, ps)
}

func encodeJSON(w http.ResponseWriter, v any) error {
	w.Header().Set("Content-Type", "application/json")
	e := json.NewEncoder(w)
	return e.Encode(v)
}

// this will be used in call side
func WriteJSON(w http.ResponseWriter, v JSON) error {
	return v.ToJSON(w)
}

func (p *Product) FromJSON(r *http.Request) error {
	e := json.NewDecoder(r.Body)
	return e.Decode(p)
}

func AddProduct(p *Product) {
	p.Id = getNextID()
	productList = append(productList, p)
}

// for generating the AUTO INCREMENTAL product-id
func getNextID() int {

	// grab the last index of the product
	lp := productList[len(productList)-1]

	// add +1 for uniqueness
	return lp.Id + 1
}

// type UpdateProductParams struct {
// 	Name        string  `json:"name"`
// 	Description string  `json:"description"`
// 	SKU         string  `json:"sku"`
// 	Price       float32 `json:"price"`
// }

func UpdateProduct(id int, p *Product) error {
	fp, i, err := findProduct(id)
	if err != nil {
		return err
	}

	fmt.Printf("product found at position %d %#v \n", i, fp)
	fmt.Printf("product in payload %#v", p)

	p.Id = id
	productList[i] = p

	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {

	for i, p := range productList {
		if p.Id == id {
			return p, i, nil

		}
	}
	return nil, 0, ErrProductNotFound
}
