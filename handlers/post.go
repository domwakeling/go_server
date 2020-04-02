package handlers

import (
	"net/http"

	"github.com/domwakeling/go_server/data"
)

/*
swagger:route POST /products Products addProduct
Add a product

Adds a product to the database.

Responses:
204: noContentResponse
400: genericErrorResponse
422: validationErrorResponse
500: genericErrorResponse
*/

// AddProduct is a Handler that receives data (in the Request.Body) and stores in db
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("[INFO] Handle POST Product")

	// the product to add has been added to request context by middleware; get it out and cast (type assertion)
	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	p.l.Printf("[DEBUG] Inesrting product: %#v\n", prod)
	data.AddProduct(prod)

	rw.WriteHeader(http.StatusNoContent)
}
