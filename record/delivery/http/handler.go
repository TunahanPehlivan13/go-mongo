package http

import (
	"github.com/TunahanPehlivan13/go-mongo/models"
	"github.com/TunahanPehlivan13/go-mongo/record"

	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type Record struct {
	Key        string    `json:"key"`
	CreatedAt  time.Time `json:"createdAt"`
	TotalCount int       `json:"totalCount"`
}

type Handler struct {
	useCase record.UseCase
}

func NewHandler(useCase record.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

type filterInput struct {
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	MinCount  int       `json:"minCount"`
	MaxCount  int       `json:"maxCount"`
}

func (handler *Handler) Post(ctx *gin.Context) {
	inp := new(filterInput)
	if err := ctx.BindJSON(inp); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	records, err := handler.useCase.GetRecords(ctx.Request.Context(),
		inp.StartDate, inp.EndDate, inp.MinCount, inp.MaxCount)

	if err != nil {
		log.Printf("Error occured while getting records with msg -> (%s)", err)
		writeResponse(ctx, records, http.StatusBadRequest, err.Error(), 1)
		return
	}
	if len(records) == 0 {
		writeResponse(ctx, records, http.StatusNotFound, record.ErrRecordNotFound.Error(), 1)
		return
	}
	writeResponse(ctx, records, http.StatusOK, "success", 0)
}

func writeResponse(ctx *gin.Context, records []*models.Record, status int, msg string, code int) {
	ctx.JSON(status, &getResponse{
		Code:    code,
		Msg:     msg,
		Records: toRecords(records),
	})
}

type getResponse struct {
	Code    int       `json:"code"`
	Msg     string    `json:"msg"`
	Records []*Record `json:"records"`
}

func toRecords(records []*models.Record) []*Record {
	out := make([]*Record, len(records))

	for i, r := range records {
		out[i] = toRecord(r)
	}
	return out
}

func toRecord(r *models.Record) *Record {
	return &Record{
		Key:        r.Key,
		CreatedAt:  r.CreatedAt,
		TotalCount: r.TotalCount,
	}
}
