package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

// Localization structure to store error messages and field names
type Localization struct {
	Messages   map[string]string
	FieldNames map[string]string
}

// Localizations for different languages
var localizations = map[string]Localization{
	"en": {
		Messages: map[string]string{
			"required": "The %s is required.",
			"min":      "The %s must be at least %s characters long.",
			"max":      "The %s must not exceed %s characters.",
			"email":    "The %s must be a valid email address.",
			"default":  "The %s is invalid.",
		},
		FieldNames: map[string]string{
			"FirstName": "first name",
			"LastName":  "last name",
			"Email":     "email",
			"Password":  "password",
		},
	},
	"bn": {
		Messages: map[string]string{
			"required": "%s আবশ্যক।",
			"min":      "%s কমপক্ষে %s অক্ষরের হতে হবে।",
			"max":      "%s %s অক্ষরের বেশি হতে পারবে না।",
			"email":    "%s একটি বৈধ ইমেইল হতে হবে।",
			"default":  "%s সঠিক নয়।",
		},
		FieldNames: map[string]string{
			"FirstName": "প্রথম নাম",
			"LastName":  "শেষ নাম",
			"Email":     "ইমেইল",
			"Password":  "পাসওয়ার্ড",
		},
	},
}

// CustomErrorMessage generates localized and dynamic error messages using struct
func CustomErrorMessage(err validator.FieldError, lang string) string {
	// Fallback to English if the specified language is not found
	localization, ok := localizations[lang]
	if !ok {
		localization = localizations["en"]
	}

	// Get the friendly name for the field, fallback to the original field name if not found
	field := localization.FieldNames[err.Field()]
	if field == "" {
		field = err.Field()
	}

	// Get the error message template, fallback to the default message
	messageTemplate := localization.Messages[err.Tag()]
	if messageTemplate == "" {
		messageTemplate = localization.Messages["default"]
	}

	// Return the formatted error message
	if err.Param() != "" {
		return fmt.Sprintf(messageTemplate, field, err.Param())
	}
	return fmt.Sprintf(messageTemplate, field)
}

func main() {
	lang := "en"

	// Define a sample struct to validate
	type User struct {
		FirstName string `validate:"required,min=3,max=80"`
		LastName  string `validate:"min=3,max=80"`
		Email     string `validate:"required,email"`
		Password  string `validate:"max=12"`
	}

	// Sample data with validation errors
	user := User{
		FirstName: "",                 // Too short
		LastName:  "L",                // Too short
		Email:     "invalid.com",      // Invalid email
		Password:  "VeryLongPassword", // Exceeds max length
	}

	// Initialize the validator
	validate := validator.New()

	// Validate the struct
	err := validate.Struct(user)
	if err != nil {
		// Iterate over validation errors
		for _, validationErr := range err.(validator.ValidationErrors) {
			fmt.Println(CustomErrorMessage(validationErr, lang))
		}
		fmt.Println()
	}
}
