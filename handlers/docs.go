package handlers

import "example.com/dwake/serve1/data"

// Following wrappers exist only for purposes of generating swagger documentation

// A list of products
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in the system.
	// In: body
	Body []data.Product
}

// No content is returned by this API endpoiint
// swagger:response noContent
type productsNoContentWrapper struct {
}

// Not found (id not present in database)
// swagger:response notFound
type productsNotFoundWrapper struct {
}

// Validation errors defined as an array of string
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

// Internal server error
// swagger:response internalServerError
type internalErrorWrapper struct {
}

// swagger:parameters deleteProduct
type productIDParameterWrapper struct {
	// The id of the product to which the operation relates
	// in: path
	// required: true
	ID int `json:"id"`
}