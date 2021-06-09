package helpers

import (
	"net/http"
)

func NewBadRequestErrorResp(message string) APIResp {
	return NewErrorResp(http.StatusBadRequest, "ERR_BAD_REQUEST", message)
}

func NewInternalErrorResp(err error) APIResp {
	return NewErrorResp(http.StatusInternalServerError, "ERR_INTERNAL_ERROR", err.Error())
}

func NewNotFoundErrorResp(message string) APIResp {
	return NewErrorResp(http.StatusNotFound, "ERR_NOT_FOUND", message)
}
