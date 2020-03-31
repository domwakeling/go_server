package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"github.com/domwakeling/go_server/data"
)

/*
swagger:route PUT /products/{id} products updateProduct
Update a product

Updates product {id} in the database. Will fail if {id} does not exist.

Responses:
204: noContentResponse
400: genericErrorResponse
422: validationErrorResponse
500: genericErrorResponse
*/

// UpdateProduct is a Handler that receives an id (through mux) and data (in the Request.Body) and updated that id in the db
func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	// get the id from mux and check that it's valid
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
	}

	p.l.Println("Handle PUT Product on id", id)

	// the product to update has been added to request context by middleware; get it out and cast (type assertion)
	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	// update the db
	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}