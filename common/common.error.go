package common

import "net/http"

type CustomError struct {
	error      string
	statusCode int
}

func (c CustomError) Error() string {
	return c.error
}

func (c CustomError) StatusCode() int {
	return c.statusCode
}

func BadRequestError(err string) CustomError {
	error := CustomError{error: err, statusCode: http.StatusBadRequest}
	return error
}

func InternalError(err string) CustomError {
	error := CustomError{error: err, statusCode: http.StatusInternalServerError}
	return error
}

func Created(err string) CustomError {
	error := CustomError{error: err, statusCode: http.StatusCreated}
	return error
}
