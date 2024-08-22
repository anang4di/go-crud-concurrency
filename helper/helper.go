package helper

import (
	"fmt"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

func FormatValidationError(validationErrors validator.ValidationErrors) []string {
	var errors []string

	for _, ve := range validationErrors {
		msg := ve.Error()

		switch ve.Tag() {
		case "required":
			msg = fmt.Sprintf("%s tidak boleh kosong", ve.Field())
		case "email":
			msg = fmt.Sprintf("%s tidak valid", ve.Field())
		case "gt":
			msg = fmt.Sprintf("%s tidak boleh kurang dari 0", ve.Field())
		}
		errors = append(errors, msg)
	}

	return errors
}

func FormatValidationErrors(err error) []string {
	var errors []string

	if sliceErr, ok := err.(binding.SliceValidationError); ok {
		for _, e := range sliceErr {
			if validationErrors, ok := e.(validator.ValidationErrors); ok {
				errors = append(errors, FormatValidationError(validationErrors)...)
			} else {
				errors = append(errors, e.Error())
			}
		}
	} else if validationErrors, ok := err.(validator.ValidationErrors); ok {
		errors = append(errors, FormatValidationError(validationErrors)...)
	} else {
		errors = append(errors, err.Error())
	}

	return errors
}

func FormatErrors(errs []error) []string {
	var errors []string

	for _, e := range errs {
		msg := e.Error()
		errors = append(errors, msg)
	}
	return errors
}
