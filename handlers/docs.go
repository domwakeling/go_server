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

// Following are only used for documentating the API in Swagger

// Error message returned as a string
// swagger:response genericErrorResponse
type errorGenericResponseWrapper struct {
	// Description of the error
	// in:body
	Body GenericError
}

// Validation errors returned as an array of strings
// swagger:response validationErrorResponse
type errorValidationResponseWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

// No content is returned by this API endpoiint on a successful call
// swagger:response noContentResponse
type responseNoContentWrapper struct {
}

// A list of products
// swagger:response productsResponse
type responseProductsWrapper struct {
	// All products in the system
	// in:body
	Body data.Products	
}

// swagger:parameters deleteProduct updateProduct
type parameterIDWrapper struct {
	// The id of the product to which the operation relates
	// in: path
	// required: true
	ID int `json:"id"`
}

// swagger:parameters addProduct updateProduct
type parameterProductWrapper struct {
	// Product data structure to Add or Update a product.
	//
	// Note: the {id} field is not required and will be ignored. New product
	// will have an {id} generated, updated product will use the {id} passed
	// in the path.
	// in:body
	// required: true
	Body data.Product
}