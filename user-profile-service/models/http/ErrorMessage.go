package http

import "fmt"

type ECODE int

type ErrorResponse struct {
	StatusCode    int    `json:"-"`
	ErrorCode     ECODE  `json:"code"`
	ErrorMessage  string `json:"errorMessage"`
	ErrorDetails  string `json:"errorDetails"`
	TransactionId string `json:"transactionId"`
}

func (er *ErrorResponse) Error() string {
	return fmt.Sprintf("%d: %s", er.ErrorCode, er.ErrorMessage)
}

const (
	// Error codes
	//NOTE: These are not the same as the HTTP status codes.

	//NO ERROR
	ECODE_OK ECODE = iota

	//INVALID INPUT ERRORS
	ECODE_ENTITY_NOT_FOUND ECODE = iota + 1000
	ECODE_MISSING_DATA
	ECODE_VALIDATION_ERROR
	ECODE_CONFLICT_ERROR

	//SYSTEM ERRORS
	ECODE_DATABASE_ERROR ECODE = iota + 2000

	//UNKNOWN ERRORS
	ECODE_UNKNOWN_ERROR ECODE = iota + 9000
)

func (e ECODE) String() string {
	switch e {
	case ECODE_OK:
		return "OK"
	case ECODE_ENTITY_NOT_FOUND:
		return "ENTITY_NOT_FOUND"
	case ECODE_MISSING_DATA:
		return "MISSING_DATA"
	case ECODE_VALIDATION_ERROR:
		return "VALIDATION_ERROR"
	case ECODE_CONFLICT_ERROR:
		return "CONFLICT_ERROR"
	case ECODE_DATABASE_ERROR:
		return "DATABASE_ERROR"
	case ECODE_UNKNOWN_ERROR:
		return "UNKNOWN_ERROR"
	default:
		return "UNKNOWN_ERROR"
	}
}
