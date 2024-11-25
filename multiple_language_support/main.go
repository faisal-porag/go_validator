package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

// Localization structure for error messages and field names
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

// CustomErrorMessage generates a localized error message
func CustomErrorMessage(err validator.FieldError, lang string) string {
	// Fallback to English if language is not found
	localization, ok := localizations[lang]
	if !ok {
		localization = localizations["en"]
	}

	// Friendly field name
	field := localization.FieldNames[err.Field()]
	if field == "" {
		field = err.Field()
	}

	// Error message template
	messageTemplate := localization.Messages[err.Tag()]
	if messageTemplate == "" {
		messageTemplate = localization.Messages["default"]
	}

	// Format and return the error message
	if err.Param() != "" {
		return fmt.Sprintf(messageTemplate, field, err.Param())
	}
	return fmt.Sprintf(messageTemplate, field)
}

// ValidateAndTranslate validates a struct and returns all localized error messages
func ValidateAndTranslate(data interface{}, lang string) []string {
	validate := validator.New()
	err := validate.Struct(data)
	if err != nil {
		var errorMessages []string
		for _, validationErr := range err.(validator.ValidationErrors) {
			errorMessages = append(errorMessages, CustomErrorMessage(validationErr, lang))
		}
		return errorMessages
	}
	return nil
}

func main() {
	// Choose a language
	lang := "en" // Change to "en" for English

	// Define a struct to validate
	type User struct {
		FirstName string `validate:"required,min=3,max=80"`
		LastName  string `validate:"min=3,max=80"`
		Email     string `validate:"required,email"`
		Password  string `validate:"max=12"`
	}

	// Example user data with errors
	user := User{
		FirstName: "porag",            // Missing (required)
		LastName:  "L",                // Too short
		Email:     "invalid.com",      // Invalid email
		Password:  "VeryLongPassword", // Exceeds max length
	}

	// Validate the struct and get error messages
	errors := ValidateAndTranslate(user, lang)
	if errors != nil {
		for _, msg := range errors {
			fmt.Println(msg)
		}
	} else {
		fmt.Println("Validation passed!")
	}
}


