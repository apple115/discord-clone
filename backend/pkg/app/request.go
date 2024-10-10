package app

import (
	"log"

	"github.com/go-playground/validator/v10"
)

func MarkError(errors []*validator.FieldError) {
	for _, err := range errors {
		log.Printf("err:%v", err)
	}
}
