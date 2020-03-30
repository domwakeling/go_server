package data

import (
	"time"
	"fmt"
)

// Product is the struct definition for our API
type Product struct {
	ID			int		`json:"id"`
	Name		string	`json:"name" validate:"required"`
	Description string	`json:"description"`
	Price		float32 `json:"price" validate:"gt=0"`
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
		SKU:		 "abd123",
		CreatedOn:	 time.Now().UTC().String(),
		UpdatedOn:	 time.Now().UTC().String(),
	},
	&Product{
		ID:			 2,
		Name:		 "Espresso",
		Description: "Short and strong coffee without milk",
		Price:		 1.95,
		SKU:		 "def456",
		CreatedOn:	 time.Now().UTC().String(),
		UpdatedOn:	 time.Now().UTC().String(),
	},
}