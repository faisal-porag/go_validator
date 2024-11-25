# Custom data validation using Golang
Validation using Golang and validate


```go
type User struct {
	FirstName string `validate:"required,min=3,max=80"`
	LastName  string `validate:"max=80"`
	Email     string `validate:"required,email"`
}
```


```go
func CustomErrorMessage(err validator.FieldError) string {
	// Map struct field names to user-friendly field names
	fieldNames := map[string]string{
		"FirstName": "first name",
		"LastName":  "last name",
		"Email":     "email",
	}

	field := fieldNames[err.Field()] // Get the friendly name

	switch err.Tag() {
	case "required":
		return fmt.Sprintf("The %s is required.", field)
	case "min":
		return fmt.Sprintf("The %s must be at least %s characters long.", field, err.Param())
	case "max":
		return fmt.Sprintf("The %s must not exceed %s characters.", field, err.Param())
	case "email":
		return fmt.Sprintf("The %s must be a valid email address.", field)
	default:
		return fmt.Sprintf("The %s is invalid.", field)
	}
}
```


```go
func main() {
	validate := validator.New()

	// Creating an instance of User with invalid data for testing
	user := User{
		FirstName: "",                                                                      // Less than min length
		LastName:  "A very very very very very long last name exceeding eighty characters", // Exceeds max length
		Email:     "invalid-email",                                                         // Invalid email format
	}

	// Validate the struct
	err := validate.Struct(user)
	if err != nil {
		// Iterate through the validation errors and print custom error messages
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(CustomErrorMessage(err))
		}
	}
}
```


