package http

import (
	"github.com/TunahanPehlivan13/go-mongo/item"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Item struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Handler struct {
	useCase item.UseCase
}

func NewHandler(useCase item.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (handler *Handler) Post(ctx *gin.Context) {
	inp := new(Item)
	if err := ctx.BindJSON(inp); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := handler.useCase.Persist(inp.Key, inp.Value)

	if err != nil {
		log.Printf("Error occured while saving item with msg -> (%s)", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, inp)
}

func (handler *Handler) Get(ctx *gin.Context) {
	query := ctx.Request.URL.Query()
	key := query.Get("key")

	item, err := handler.useCase.Get(key)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	ctx.JSON(http.StatusOK, item)
}
