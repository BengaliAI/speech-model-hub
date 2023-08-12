package utils

import "github.com/go-playground/validator"

func ValidateStruct(model interface{}) error {
	validate := validator.New()
	err := validate.Struct(model)
	if err != nil {
		return err
	}
	return nil
}
