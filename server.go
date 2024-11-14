package main

import (
	"context"
	"restapi/config"
	"restapi/handler"
	"restapi/middleware"
	"restapi/repository"
	"restapi/service"
	"restapi/utils/http_response"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupServer(router *gin.Engine) {
	ctx := context.Background()

	mongoConn := config.NewMongoConn(ctx)
	defer mongoConn.Close(ctx)
	mongoDb := mongoConn.Database(config.Envs.MONGO_DB_NAME)

	responseWriter := http_response.NewResponseWriter()

	// repositories
	userRepo := repository.NewUserRepo(mongoDb.Collection("users"))

	// services
	authService := service.NewAuthService(userRepo)

	// handlers
	authHandler := handler.NewAuthHandler(responseWriter, authService)
	_ = authHandler
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"error":   false,
			"message": "pong!",
		})
	})

	secureRouter := router.Group("/")
	{
		secureRouter.Use(middleware.AuthMiddleware(responseWriter, authService))
	}

	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
