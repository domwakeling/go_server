/*
Package classification Product API

Documentation for coffee shop's Product API. This is a trial for
completing a microservice using GoLang.

Schemes: http
BasePath: /
Version: 0.1.0
License: MIT http://opensource.org/licenses/MIT
Contact: Dom Wakeling<info@domwakeling.com> http://domwakeling.com

Consumes:
- application/json
Produces:
- application/json
swagger:meta
*/
package handlers

import "github.com/domwakeling/go_server/data"

// A list of products
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in the system.
	// in:body
	Body data.Products	
}

// No content is returned by this API endpoiint
// swagger:response noContent
type productsNoContentWrapper struct {
}

// Not found (id not present in database)
// tempcode:response notFound
type productsNotFoundWrapper struct {
}

// Validation errors defined as an array of string
// tempcode:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

// Internal server error
// tempcode:response internalServerError
type internalErrorWrapper struct {
}

// tempcode:parameters deleteProduct
type productIDParameterWrapper struct {
	// The id of the product to which the operation relates
	// in: path
	// required: true
	ID int `json:"id"`
}