package router

import (
	"context"

	"github.com/Aranyak-Ghosh/spotigo/user_profile/controllers"
	"github.com/Aranyak-Ghosh/spotigo/user_profile/middlewares"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func SetupRoutes(userController controllers.UserController, logger *zap.SugaredLogger) *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.TransactionIdGenerator())
	router.Use(middlewares.DefaultCORSMiddleware())

	v1 := router.Group("api/v1")

	user := v1.Group("users")

	userController.RegisterRoutes(user)
	router.GET("/health-check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "alive",
		})
	})

	router.Run()

	return router
}

func registerHooks(lifecycle fx.Lifecycle, ginRouter *gin.Engine, logger *zap.SugaredLogger) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("Initializing server")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Terminating server")
			logger.Sync()
			return nil
		},
	})
}

var Module = fx.Options(fx.Provide(SetupRoutes), fx.Invoke(registerHooks))
