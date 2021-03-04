package errorhandler

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/julianskyline/motorcars-users/domain/exceptions"
	"github.com/pkg/errors"
)

/*
ErrorHandler Manejador de errores
Return funcion gin HandlerFunc para manejar errores en el router
*/
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Ejecuta los manejadores pendientes en ejecutar dentro del presente manejador
		c.Next()

		//Se verifica la lista de errores generados
		err := c.Errors.Last()
		//Se retorna porque no existe error
		if err == nil {
			return
		}

		//Se define la respuesta por defecto
		errorResponse := getErrorResponse(err, exceptions.UnknownError)

		//Se verifica el tipo de error si es generico
		if _, ok := errors.Cause(err.Err).(exceptions.GeneralException); ok {

			//Se castea el tipo de error
			exception, ok := err.Err.(*exceptions.GeneralException)

			//Se valida si la conversion fue correcta para solicitar respuesta
			if ok {
				errorResponse = getErrorResponse(err, exception.ErrorType)
			}

		} else {
			errorResponse = getErrorResponse(err, exceptions.ServerError)
		}
		//Se hace log de errores
		log.Fatal(errorResponse.ErrorMessage, err)
		//Se serializa el json del error
		c.JSON(errorResponse.ErrorStatus, errorResponse)

	}
}

func ErrorHandlerSingle(err error, c *gin.Context) {
	exception := err.(exceptions.GeneralException)
	errorResponse := getErrorResponse(exception, exception.ErrorType)
	//Se hace log de errores
	log.Fatal(errorResponse.ErrorMessage, exception)
	//Se serializa el json del error
	c.JSON(errorResponse.ErrorStatus, errorResponse)
}

/*
	func que verifica el tipo de error y consulta la respuesta
*/
func getErrorResponse(err error, errorType exceptions.ErrorType) ErrorResponse {

	if exceptions.DataError == errorType ||
		exceptions.BusinessError == errorType {
		//Verifica si el error es de tipo badrequest
		return ErrorResponseHandler(err, NewBadRequestError)
	} else if exceptions.PermissionDenied == errorType {
		//Verifica si el error es de tipo 401
		return ErrorResponseHandler(err, NewUnauthorizedError)
	} else {
		//Abarca errores de servidor
		return ErrorResponseHandler(err, NewInternalServer)
	}

}
