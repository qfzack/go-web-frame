package validator

import "github.com/go-playground/validator"

type Validator struct {
	validate *validator.Validate
}

func (v *Validator) ValidateStruct(s interface{}) error {
	// TODO: input verify, XSS, SQL injection
	return nil
}
