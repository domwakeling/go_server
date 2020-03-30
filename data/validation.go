package data

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator"
)

// ValidationError is a wrapper to validator package's FieldError
type ValidationError struct {
	validator.FieldError
}

// Error method on a ValidationError retrives the FieldError details as a string
func (v ValidationError) Error() string {
	return fmt.Sprintf(
		"Key: '%s' Error: Field validation for '%s' failed on the '%s' tag",
		v.Namespace(),
		v.Field(),
		v.Tag(),
	)
}

// ValidationErrors is a collection of ValidationError
type ValidationErrors []ValidationError

// Errors method on a ValidationErrors collection returns a collection of error strings
func (v ValidationErrors) Errors() []string {
	errs := []string{}
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

// Validation is a struct wrapping a pointer to a validator.Validate struct
type Validation struct {
	validate *validator.Validate
}

// NewValidation creates a concrete Validation struct with our required sku validation
func NewValidation() *Validation {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return &Validation{validate}
}

// Validate uses the validator Struct method and retuns a (data) ValidationErrors collection
func (v *Validation) Validate(i interface{}) ValidationErrors {
	err := v.validate.Struct(i)

	// if there are no errors, return nil
	if err == nil {
		return nil
	}

	// otherwise convert the err to validation.ValidationErrors and then to a slice of (data) ValidationErrors
	errs := err.(validator.ValidationErrors)
	var returnErrs []ValidationError
	for _, err := range errs {
		// cast the FieldError into a (data) ValidationError and append
		ve := ValidationError{err.(validator.FieldError)}
		returnErrs = append(returnErrs, ve)
	}
	return returnErrs
}

func validateSKU(fl validator.FieldLevel) bool {
	// sku format three groups of letters separated by dash eg abc-defg-hijk
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1) // slice of string
	// should only get one matched group ...
	if len(matches) != 1 {
		return false
	}
	return true
}