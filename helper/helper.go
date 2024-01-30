package helper

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}
	return jsonResponse
}

func FormatValidationError(err error) []string {
	var errors []string

	// Check if err is of type validator.ValidationErrors
	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		// Handle the error as a validator.ValidationErrors
		for _, e := range validationErrs {
			errors = append(errors, e.Error())
		}
		return errors
	}

	// If not a validator.ValidationErrors, handle other error types
	if jsonSyntaxErr, ok := err.(*json.SyntaxError); ok {
		errors = append(errors, jsonSyntaxErr.Error())
		return errors
	}

	// Handle other types of errors
	errors = append(errors, fmt.Sprintf("Unexpected error: %v", err))
	return errors
}
