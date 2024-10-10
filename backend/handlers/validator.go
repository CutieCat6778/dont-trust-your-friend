package handlers

import "github.com/go-playground/validator/v10"

type (
	ValidatorHandler struct {
		*validator.Validate
	}

	ErrorResponse struct {
		Error       bool
		FailedField string
		Tag         string
		Value       interface{}
	}
)

var (
	Validator *ValidatorHandler
)

func NewValidator() *ValidatorHandler {
	v := validator.New()
	Validator = &ValidatorHandler{v}
	return Validator
}
