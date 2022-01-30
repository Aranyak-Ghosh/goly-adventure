package errorhandling

import (
	"errors"

	validationErrors "github.com/Aranyak-Ghosh/spotify/errors"
	httpModels "github.com/Aranyak-Ghosh/spotify/models/http"
	"gorm.io/gorm"
)

func HandleDatabaseError(foundError error, transactionId string) *httpModels.ErrorResponse {
	if errors.Is(foundError, gorm.ErrRecordNotFound) {
		return generateErrorObject(httpModels.ECODE_ENTITY_NOT_FOUND, foundError.Error(), transactionId)
	} else if errors.Is(foundError, gorm.ErrInvalidData) || errors.Is(foundError, gorm.ErrInvalidField) {
		return generateErrorObject(httpModels.ECODE_VALIDATION_ERROR, foundError.Error(), transactionId)
	} else if errors.Is(foundError, gorm.ErrModelValueRequired) || errors.Is(foundError, gorm.ErrPrimaryKeyRequired) || errors.Is(foundError, validationErrors.ErrUUIDValidationFailed) {
		return generateErrorObject(httpModels.ECODE_MISSING_DATA, foundError.Error(), transactionId)
	} else {
		return generateErrorObject(httpModels.ECODE_DATABASE_ERROR, foundError.Error(), transactionId)
	}
}

func generateErrorObject(errorCode httpModels.ECODE, errorDetails string, transactionId string) *httpModels.ErrorResponse {
	return &httpModels.ErrorResponse{
		ErrorCode:     errorCode,
		ErrorMessage:  errorCode.String(),
		ErrorDetails:  errorDetails,
		TransactionId: transactionId,
	}
}
