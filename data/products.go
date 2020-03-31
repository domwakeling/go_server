package data

import (
	"time"
	"fmt"
)

// Product defines the structure for a product under this API.
// swagger:model
type Product struct {
	// the id for the product
	//
	// required: false
	// min: 1
	ID			int		`json:"id"`

	// the name for this product
	//
	// required: true
	// max length: 255
	Name		string	`json:"name" validate:"required"`

	// the description for this product
	//
	// required: true
	// max length: 1000
	Description string	`json:"description"`

	// the price for the producct
	//
	// required: true
	// min: 0.01
	Price		float32 `json:"price" validate:"gt=0"`

	// the SKU for the product
	//
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	SKU			string	`json:"sku" validate:"required,sku"`

	CreatedOn	string	`json:"-"`
	UpdatedOn	string	`json:"-"`
	DeletedOn	string	`json:"-"`
}

// Products is a slice of Product
type Products []*Product

// GetProducts exports the product data as a 'Products' type - declaring the return
// signature appears to automatically MAKE it a Products with the attendant method
func GetProducts() Products {
	return productList
}

// AddProduct aads a new Product to the store, generating new ID along the way
func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

// UpdateProduct updates an existing product in the store
func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)

	if err != nil {
		return err
	}

	p.ID = id

	productList[pos] = p
	return nil
}

// DeleteProduct deletes an existing product in the store
func DeleteProduct(id int) error {
	_, pos, err := findProduct(id)

	if err != nil {
		return err
	}

	productList = append(productList[:pos], productList[pos+1:]...)
	return nil
}

func findProduct(id int) (*Product, int, error) {
	for i,p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

// ErrProductNotFound is an error message when the project isn't found
var ErrProductNotFound = fmt.Errorf("Product not found")

func getNextID() int {
	lastProduct := productList[len(productList) - 1]
	return lastProduct.ID + 1
}

var productList = []*Product{
	&Product{
		ID:			 1,
		Name:		 "Latte",
		Description: "Frothy milky coffee",
		Price:		 2.45,
		SKU:		 "aa-bb-cc",
		CreatedOn:	 time.Now().UTC().String(),
		UpdatedOn:	 time.Now().UTC().String(),
	},
	&Product{
		ID:			 2,
		Name:		 "Espresso",
		Description: "Short and strong coffee without milk",
		Price:		 1.95,
		SKU:		 "dd-ee-ff",
		CreatedOn:	 time.Now().UTC().String(),
		UpdatedOn:	 time.Now().UTC().String(),
	},
}