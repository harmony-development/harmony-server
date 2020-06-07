package http

import "github.com/go-playground/validator/v10"

// HarmonyValidator validates request bindings
type HarmonyValidator struct {
	Validator *validator.Validate
}

// Validate checks if a specific struct has valid fields
func (hv HarmonyValidator) Validate(i interface{}) error {
	return hv.Validator.Struct(i)
}
