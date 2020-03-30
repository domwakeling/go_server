package handlers

import (
	"net/http"

	"github.com/domwakeling/go_server/data"
)

/*
swagger:route GET /products products listProducts
List all products

Returns a list of all products in the database.

Responses:
200: productsResponse
*/

// GetProducts is a Handler that writes out the current db of Products
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	prods := data.GetProducts()
	err := data.ToJSON(prods, rw)
	if err != nil {
		// we should never get here, but just in case
		p.l.Println("[ERROR} serializing product", err)
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}