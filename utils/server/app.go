package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-simple-api/cmd/core/auth"
	"go-simple-api/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	httpServer *http.Server

	authRepository *auth.Repository
}

func NewApp() *App {
	db := initDB()

	return &App{
		authRepository: auth.NewAuthRepository(db, "users"),
	}
}

func (a *App) Run(port string) error {
	router := gin.Default()
	router.Use(gin.Recovery(), gin.Logger())

	api := router.Group("/api")

	auth.RegisterHTTPEndpoints(api, a.authRepository)

	// HTTP Server
	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}

func initDB() *mongo.Database {
	clientOptions := options.Client().ApplyURI(config.Env.MongoUrl)
	client, err := mongo.Connect(clientOptions)

	if err != nil {
		log.Fatalf("Error occured while establishing connection to mongoDB")
	}

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return client.Database(config.Env.DbName)
}
