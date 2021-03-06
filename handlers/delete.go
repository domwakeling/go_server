package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/domwakeling/go_server/data"
)

/*
swagger:route DELETE /products/{id} Products deleteProduct
Delete product for a given {id}

Deletes a product from the database based on {id}.

Responses:
204: noContentResponse
404: genericErrorResponse
500: genericErrorResponse
*/

// DeleteProduct is a Handler that receives an id and deletes from the db
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	// extract id and ensure it is in valid format
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to parse id", http.StatusBadRequest)
	}

	// attempt to handle request ...
	p.l.Println("[INFO} Handle DELETE Product on id", id)
	err = data.DeleteProduct(id)
	// ... writing a StatusNotFound if appropriate error (succesful check but no data) ...
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	// ... or a StatusInternalServerError if other error (something has gone wrong) ...
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
	// ... otherwise succesful, write a StatusNoContent (204)
	rw.WriteHeader(http.StatusNoContent)
}
