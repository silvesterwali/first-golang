package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) gin.H {
	errors := make(map[string][]string)

	if validationError, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationError {
			field := fieldError.Field()
			tag := fieldError.Tag()

			var message string

			switch tag {
			case "required":
				message = "The field " + field + " is required"

			case "email":
				message = "The field " + field + " must be a valid email"

			case "min":
				message = "The field " + field + " must be at least " + fieldError.Param()

			case "max":
				message = "The field " + field + " must be less than " + fieldError.Param()

			case "numeric":
				message = "The field " + field + " must be numeric"

			case "alpha":
				message = "The field " + field + " must be alpha"

			case "alphanum":
				message = "The field " + field + " must be alpha numeric"

			case "email_unique":
				message = "The " + field + " is already in use"

			default:
				message = "The field " + field + " is invalid"
			}

			errors[field] = append(errors[field], message)
		}
	}

	return gin.H{
		"message": "Validation Error",
		"errors":  errors,
	}
}

func FormatDefaultError(err error, message ...string) gin.H {
	var msg string

	if len(message) > 0 {
		msg = message[0]
	} else {
		msg = "Something went wrong"
	}

	return gin.H{
		"message": msg,
		"errors":  err.Error(),
	}
}
