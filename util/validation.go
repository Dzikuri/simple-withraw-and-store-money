package util

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

var Validator = validator.New()

type Validation struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationResponse struct {
	Success     bool         `json:"success"`
	Validations []Validation `json:"message"`
	Data        interface{}  `json:"data"`
}

func GenerateValidationMessage(field string, rule string) (message string) {
	switch rule {
	case "required":
		return fmt.Sprintf("Field '%s' is '%s'.", field, rule)
	default:
		return fmt.Sprintf("Field '%s' is not valid.", field)
	}
}

func GenerateValidationResponse(err error) (response ValidationResponse) {
	var a struct{}
	response.Success = false
	response.Data = a
	var validations []Validation

	// Set validation error
	fieldErr := err.(validator.ValidationErrors)

	for _, value := range fieldErr {

		// Get validation field & rule
		field, rule := value.Field(), value.Tag()

		// Create validation object
		validation := Validation{Field: field, Message: GenerateValidationMessage(field, rule)}

		// Add validation object to validations
		validations = append(validations, validation)
	}

	// Set validation response
	response.Validations = validations

	return response

}

func ValidateEmailOrUsername(input string) bool {
	// Regular expression for validating an email
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	emailMatch, _ := regexp.MatchString(emailRegex, input)

	// Regular expression for validating a username (letters, numbers, and underscores, 3-20 chars)
	usernameRegex := `^[a-zA-Z0-9_]{3,20}$`
	usernameMatch, _ := regexp.MatchString(usernameRegex, input)

	return emailMatch || usernameMatch
}
