package validator

import "regexp"

var (
	EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

//define validator that contains a map of validator errors
type Validator struct {
	Errors map[string]string
}

//New is a helper which creates a new validator with an empty errors map

func New() *Validator {
	return &Validator{
		Errors: make(map[string]string),
	}
}

//return true if errors map is empty
func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

//adderror adds an error message to the map so long the entry does not exist
func (v *Validator) AddError(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

//check adds an error message to the map only if a validation check is not ok
func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

//generic function which returns true if a specific value is in a list
func PermittedValue[T comparable](value T, permittedValues ...T) bool {
	for i := range permittedValues {
		if value == permittedValues[i] {
			return true
		}
	}
	return false
}

//matched returns true if a string value matches a specific regexp pattern

func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

//Generic function which returns true if all values in a slice are unique
func Unique[T comparable](values []T) bool {
	uniqueValues := make(map[T]bool)

	for _, value := range values {
		uniqueValues[value] = true
	}
	return len(values) == len(uniqueValues)
}
