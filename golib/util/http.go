package util

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func BindJson[T any](r *http.Request, v T) error {
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		return err
	}
	return nil
}

func ShouldBindJson[T any](r *http.Request, v T) error {
	err := BindJson(r, v)
	if err != nil {
		return err
	}
	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(v)
	if err != nil {
		return err
	}
	return nil
}
