package server

import (
	"context"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-simple-api/cmd/core/auth"
	"go-simple-api/cmd/core/post"
	"go-simple-api/config"
	"go-simple-api/docs"
	"go-simple-api/utils/constants"
	"go-simple-api/utils/middleware"
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
	postRepository *post.Repository
}

func NewApp() *App {
	db := initDB()

	return &App{
		authRepository: auth.NewRepository(db, constants.Collections.User),
		postRepository: post.NewRepository(db, constants.Collections.Post),
	}
}

func (app *App) Run(port string) error {
	docs.SwaggerInfo.Title = "Simple backend"
	docs.SwaggerInfo.Description = "Backend API"
	docs.SwaggerInfo.Version = "1.0"

	router := gin.Default()
	router.Use(gin.Recovery(), gin.Logger())
	router.GET("/swagger/*any", middleware.BaseAuthSwagger(), ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")

	auth.RegisterHTTPEndpoints(api, app.authRepository)
	post.RegisterHTTPEndpoints(api, app.postRepository)

	// HTTP Server
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
