package exceptions

/*
ErrorType Tipo de error
*/
type ErrorType string

/*
UnknownError
DataError
permissionDenied
ServerError
*/
const (
	UnknownError     ErrorType = "UNKNOWN"
	DataError        ErrorType = "DATA_ERROR"
	PermissionDenied ErrorType = "PERMISSION_DENIED"
	ServerError      ErrorType = "SERVER_ERROR"
	BusinessError    ErrorType = "BUSINESS_ERROR"
)
