package data

import (
	"encoding/json"
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

// ToJSON method of Products
func (p *Products) ToJSON(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) FromJSON(r *http.Request) error {
	e := json.NewDecoder(r.Body)
	return e.Decode(p)
}

func AddProduct(p *Product) {
	p.Id = getNextID()
	productList = append(productList, p)
}

func getNextID() int {

	// grab the last index of the product
	lp := productList[len(productList)-1]

	// add +1 for uniqueness
	return lp.Id + 1
}
