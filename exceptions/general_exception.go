package exceptions

import (
	"github.com/pkg/errors"
)

/*
GeneralException Exception general
*/
type GeneralException struct {
	ErrorType    ErrorType
	ErrorMessage string
	ErrorService string
	Source       string
	Cause        []interface{}
}

func (ge GeneralException) Error() string {
	return ge.ErrorMessage
}

/*
NewError create new error
*/
func NewError(err error) GeneralException {
	return GeneralException{
		ErrorType:    UnknownError,
		ErrorMessage: err.Error(),
		Source:       errors.Wrap(err, err.Error()).Error(),
	}
}

/*
NewErrorDetailService create new error
*/
func NewErrorDetailService(err error, errorType ErrorType, errorService string) GeneralException {
	return GeneralException{
		ErrorType:    errorType,
		ErrorMessage: err.Error(),
		ErrorService: errorService,
		Source:       errors.Wrap(err, err.Error()).Error(),
		Cause:        []interface{}{},
	}
}
