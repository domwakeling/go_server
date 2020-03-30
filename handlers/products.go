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

import (
	"log"

	"github.com/domwakeling/go_server/data"
)

// Products is a Handler
type Products struct {
	l *log.Logger
	v *data.Validation
}

// NewProducts is a function that returns a new Products (which is a Handler)
// with the log.Logger that you pass in
func NewProducts(l *log.Logger, v *data.Validation) *Products {
	return &Products{l, v}
}

// ValidationError is a collection of validation error messsages
type ValidationError struct {
	Messages []string	`json:"messages"`
}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string		`json:"message"`
}

// KeyProduct is a key that will be used to add/extract product from context
type KeyProduct struct {}