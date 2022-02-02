package errorhandling

import (
	"errors"
	"net/http"

	validationErrors "github.com/Aranyak-Ghosh/spotigo/user_profile/errors"
	httpModels "github.com/Aranyak-Ghosh/spotigo/user_profile/models/http"
	"gorm.io/gorm"
)

func HandleDatabaseError(foundError error, transactionId string) *httpModels.ErrorResponse {
	if errors.Is(foundError, gorm.ErrRecordNotFound) {
		return generateErrorObject(http.StatusNotFound, httpModels.ECODE_ENTITY_NOT_FOUND, foundError.Error(), transactionId)
	} else if errors.Is(foundError, gorm.ErrInvalidData) || errors.Is(foundError, gorm.ErrInvalidField) || errors.Is(foundError, validationErrors.ErrInvalidUUID) {
		return generateErrorObject(http.StatusBadRequest, httpModels.ECODE_VALIDATION_ERROR, foundError.Error(), transactionId)
	} else if errors.Is(foundError, gorm.ErrModelValueRequired) || errors.Is(foundError, gorm.ErrPrimaryKeyRequired) {
		return generateErrorObject(http.StatusBadRequest, httpModels.ECODE_MISSING_DATA, foundError.Error(), transactionId)
	} else {
		return generateErrorObject(http.StatusInternalServerError, httpModels.ECODE_DATABASE_ERROR, foundError.Error(), transactionId)
	}
}

func generateErrorObject(statusCode int, errorCode httpModels.ECODE, errorDetails string, transactionId string) *httpModels.ErrorResponse {
	return &httpModels.ErrorResponse{
		StatusCode:    statusCode,
		ErrorCode:     errorCode,
		ErrorMessage:  errorCode.String(),
		ErrorDetails:  errorDetails,
		TransactionId: transactionId,
	}
}
