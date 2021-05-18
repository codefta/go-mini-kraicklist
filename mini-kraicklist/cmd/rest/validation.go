package main

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator"
)

func listValidation (list *List) error{
	v := validator.New()

	err := v.Struct(list)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return nil
		}

		for _, err := range err.(validator.ValidationErrors) {

			if err.Field() == "Title" {
				return errors.New("The `title` cannot be empty")
			} else if err.Field() == "Body" {
				return errors.New("The `body` cannot be empty")
			} else if err.Field() == "Tags" {
				return errors.New("The `tags` cannot be empty")
			}
		}
	}
	return nil
}