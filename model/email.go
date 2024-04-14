package model

import "github.com/go-playground/validator/v10"

type Email struct {
	From        string `validate:"required,min=1"`
	Credentials string `validate:"required,min=1"`
	To          string `validate:"required,min=1"`
	Subject     string
	Message     string `validate:"required,min=1"`
}

// Validate validates email
func (e *Email) Validate() error {
	validate := validator.New()
	return validate.Struct(e)
}
