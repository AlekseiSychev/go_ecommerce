package validations

import (
	"github.com/go-playground/validator/v10"
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	case "min":
		return "Minimum number of characters " + fe.Param()
	case "email":
		return "Еmail should contains @ , example@mail.com " + fe.Param()
	case "max":
		return "Maximum number of characters  " + fe.Param()
	case "alphaunicode":
		return "Only alphabet supported " + fe.Param()
	case "numeric":
		return "Only numbers supported " + fe.Param()
	}
	
	return "Unknown error"
}


