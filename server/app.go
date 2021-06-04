package server

import (
	"context"
	"github.com/TunahanPehlivan13/go-mongo/record"
	recordHttp "github.com/TunahanPehlivan13/go-mongo/record/delivery/http"
	recordRepo "github.com/TunahanPehlivan13/go-mongo/record/repository/mongo"
	"github.com/TunahanPehlivan13/go-mongo/record/usecase"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	httpServer *http.Server

	recordUseCase record.UseCase
}

func NewApp() *App {
	database := initMongoDB()

	recordRepo := recordRepo.NewRecordRepository(database, "records")

	return &App{
		recordUseCase: usecase.NewRecordUseCase(recordRepo),
	}
}

func (app *App) Run(port string) error {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	recordHttp.RegisterHTTPEndpoints(router, app.recordUseCase)

	app.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := app.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return app.httpServer.Shutdown(ctx)
}

func initMongoDB() *mongo.Database {
	clientOptions := options.Client().ApplyURI("mongodb+srv://challenge-xzwqd.mongodb.net/?retryWrites=true").
		SetAuth(options.Credential{
			Username: "challengeUser", Password: "WUMglwNBaydH8Yvu",
		})

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatalf("Error occured while establishing connection to mongoDB")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return client.Database("getir-case-study")
}
