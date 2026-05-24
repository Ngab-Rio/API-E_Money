package utils

import "github.com/go-playground/validator/v10"

func ValidateHandler(err error) map[string]string {
	errors := make(map[string]string)

	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, v := range ve {
			switch v.Tag() {
			case "required":
				errors[v.Field()] = "Field is required"
			case "email":
				errors[v.Field()] = "Field must be an email"
			case "min":
				errors[v.Field()] = "Field must be at least " + v.Param() + " characters long"
			default:
				errors[v.Field()] = "Invalid field"
			}
		}
	}
	return errors
}
