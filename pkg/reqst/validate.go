package reqst

import (
	"github.com/go-playground/validator/v10"
)

func IsValid[T any](payload T) error {
	//validation
	validate := validator.New()
	err := validate.Struct(payload)
	return err
}
