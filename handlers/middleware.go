package handlers

import (
	"context"
	"net/http"

	"example.com/dwake/serve1/data"
)

// MiddlewareProductValidation is a middleware to validate the product (as JSON) and add to Request context
func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		
		prod := &data.Product{}						// get an empty Product
		
		err := data.FromJSON(prod, r.Body)			// read in; FromJSON method on a Product returns an err
		if err != nil {
			p.l.Println("[ERROR] deserialzing product", err)
			rw.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, rw)
			return
		}

		// try to validate
		errs := p.v.Validate(prod)
		if len(errs) != 0 {
			p.l.Println("[ERROR] validating product", errs)
			rw.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJSON(&ValidationError{Messages: errs.Errors()}, rw)
			return
		}

		// get the context from the original request and add a new key/value pair to get to the de-serialized product
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		// make a new request with the revised context
		req := r.WithContext(ctx)
		next.ServeHTTP(rw, req)
	})
}