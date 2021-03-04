package errorhandler

import (
	"net/http"
)

/*
TypeStatus const type status
*/
type TypeStatus string

/*
NewBadRequestError
NewFoundError
NewUnauthorizedError
NewInternalServer
*/
const (
	NewBadRequestError   TypeStatus = "NEW_BAD_REQUEST_ERROR"
	NewFoundError        TypeStatus = "NEW_NOT_FOUND_ERROR"
	NewUnauthorizedError TypeStatus = "NEW_UNATHORIZED_ERROR"
	NewInternalServer    TypeStatus = "NEW_ERROR_INTERNAL_SERVER"
)

/*
ErrorResponse struct reponse
*/
type ErrorResponse struct {
	ErrorMessage string        `json:"message"`
	ErrorStatus  int           `json:"status"`
	Error        string        `json:"error"`
	Cause        []interface{} `json:"causes"`
}

/*
ErrorResponseHandler func generate response
*/
func ErrorResponseHandler(err error, typeStatus TypeStatus) ErrorResponse {

	// Se crea respuesta
	response := ErrorResponse{
		ErrorMessage: err.Error(),
		ErrorStatus:  500,
		Error:        "UNDEFINED",
		Cause:        []interface{}{},
	}

	//Se selecciona el tipo segun parametro para completar respuesta
	switch typeStatus {
	case NewBadRequestError:
		response.ErrorStatus = http.StatusBadRequest
		response.Error = string(NewBadRequestError)
		break

	case NewFoundError:
		response.ErrorStatus = http.StatusNotFound
		response.Error = string(NewFoundError)
		break

	case NewUnauthorizedError:
		response.ErrorStatus = http.StatusUnauthorized
		response.Error = string(NewUnauthorizedError)
		break

	default:
		response.ErrorStatus = http.StatusInternalServerError
		response.Error = string(NewInternalServer)
	}

	return response

}

/*
NewErrorResponse error new response
*/
func NewErrorResponse(message string, status int, err string, causes []interface{}) ErrorResponse {
	return ErrorResponse{
		ErrorMessage: message,
		ErrorStatus:  status,
		Error:        err,
		Cause:        causes,
	}
}
