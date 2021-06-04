package http

import (
	"github.com/TunahanPehlivan13/go-mongo/record"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, recordUseCase record.UseCase) {
	handler := NewHandler(recordUseCase)

	records := router.Group("/records")
	{
		records.POST("", handler.Post)
	}
}
