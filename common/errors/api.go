// Package commonerrors holds whole logic of APIErrors
package commonerrors

import (
	"fmt"
	"net/http"
)

type Category string

const (
	CategoryValidation       Category = "VALIDATION"
	CategoryNotFound         Category = "NOT_FOUND"
	CategoryInternal         Category = "INTERNAL"
	CategoryRequestTimeout   Category = "REQUEST_TIMEOUT"
	CategoryDuplicate        Category = "DUPLICATE"
	CategoryMethodNotAllowed Category = "METHOD_NOT_ALLOWED"
)

type API struct {
	Category      Category
	StatusCode    int
	Message       string
	IsOperational bool
}

func NewAPI(statusCode int, category Category, message string, isOperational bool) *API {
	return &API{
		Category:      category,
		StatusCode:    statusCode,
		Message:       message,
		IsOperational: isOperational,
	}
}

func (apiE *API) Error() string {
	return fmt.Sprintf("api error: %s", apiE.Message)
}

func NotFound(msg string) *API {
	return &API{Category: CategoryNotFound, StatusCode: http.StatusNotFound, IsOperational: true, Message: msg}
}

func Validation(msg string) *API {
	return &API{
		Category: CategoryValidation, StatusCode: http.StatusUnprocessableEntity, IsOperational: true, Message: msg,
	}
}

func InvalidJSONFormat() *API {
	return &API{
		Category:      CategoryValidation,
		StatusCode:    http.StatusUnprocessableEntity,
		IsOperational: true,
		Message:       "provided invalid json format",
	}
}

func RequestTimeout() *API {
	return &API{
		Category: CategoryRequestTimeout, StatusCode: http.StatusUnauthorized, IsOperational: true, Message: "",
	}
}
