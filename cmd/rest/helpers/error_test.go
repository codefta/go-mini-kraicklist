package helpers

import (
	"errors"
	"testing"
)

func TestNewBadRequestErrorStatusResp(t *testing.T){
	var (
		status int = 400
		message string = "This is error bad request"
	)

	result := NewBadRequestErrorResp(message)

	if result.StatusCode != status {
		t.Errorf("Status Code doesn't corresponded to expected value: %d", status)
	}
}

func TestNewBadRequestErrorErrTypeResp(t *testing.T) {
	var (
		errType string = "ERR_BAD_REQUEST"
		message string = "This is error bad request"
	)

	result := NewBadRequestErrorResp(message)

	if result.Err != errType {
		t.Errorf("Err Type doesn't corresponded to expected value: %s", errType)
	}
}

func TestNewBadRequestErrorMessageResp(t *testing.T) {
	message := "This is error bad request"

	result := NewBadRequestErrorResp(message)

	if result.Message != message {
		t.Errorf("Err Message doesn't corresponded to expected value: %s", message)
	}
}

func TestNewNotFoundErrorStatusResp(t *testing.T) {
	var (
		status int = 404
		message string = "This is not found error"
	)

	result := NewNotFoundErrorResp(message)

	if result.StatusCode != status {
		t.Errorf("Status Code doesn't corresponded to expected value: %d", status)
	}
}

func TestNewNotFoundErrorErrTypeResp(t *testing.T) {
	var (
		errType string = "ERR_NOT_FOUND"
		message string = "This is not found error"
	)

	result := NewNotFoundErrorResp(message)

	if result.Err != errType {
		t.Errorf("Err Type doesn't corresponded to expected value: %s", errType)
	}
}

func TestNewNotFoundErrorMessageResp(t *testing.T) {
	message := "This is not found error"

	result := NewNotFoundErrorResp(message)

	if result.Message != message {
		t.Errorf("Err Message doesn't corresponded to expected value: %s", message)
	}
}

func TestNewInternalServerErrorStatusCodeResp(t *testing.T) {
	status := 500
	message := errors.New("This is internal server error")

	result := NewInternalErrorResp(message)

	if result.StatusCode != status {
		t.Errorf("Status Code doesn't corresponded to expected value: %d", status)
	}
}

func TestNewInternalServerErrorErrTypeResp(t *testing.T) {
	errType := "ERR_INTERNAL_ERROR"
	message := errors.New("This is internal server error")

	result := NewInternalErrorResp(message)

	if result.Err != errType {
		t.Errorf("Err Type doesn't corresponded to expected value: %s", errType)
	}
}

func TestNewInternalServerErrorMessageResp(t *testing.T) {
	message := errors.New("This is internal server error")

	result := NewInternalErrorResp(message)

	if result.Message != message.Error() {
		t.Errorf("Err Message doesn't corresponded to expected value: %s", message)
	}
}