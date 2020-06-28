package handlers

import "main.go/data"

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}


// A list of products
// swagger:response productsResponse
type productsResponse struct {
	// All current products
	// in: body
	Body []data.Product
}


// swagger:response noContent
type productNoContent struct {

}

// swagger:parameters deleteProduct
type productIDParameterWrapper struct {
	// The Id of the product to delete from the database
	// in:path
	// required: true
	ID int	 `json:"id"`
}


// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}