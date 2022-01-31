package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func injectTransactionId(ctx *gin.Context) {
	ctx.Request.Header.Set("transactionId", uuid.NewString())
}

func TransactionIdGenerator() gin.HandlerFunc {
	return injectTransactionId
}
