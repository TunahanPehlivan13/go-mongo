package http

import (
	"github.com/TunahanPehlivan13/go-mongo/item"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, itemUseCase item.UseCase) {
	handler := NewHandler(itemUseCase)

	records := router.Group("/in-memory")
	{
		records.POST("", handler.Post)
		records.GET("", handler.Get)
	}
}
