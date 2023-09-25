package main

import (
	"context"
	"speech-model-hub/internal/apis"
	"speech-model-hub/internal/configs"
	"speech-model-hub/internal/infrastructure"
	"speech-model-hub/internal/services"
	"speech-model-hub/pkg/loggerFx"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func home(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}

func NewGINApp() *gin.Engine {
	gin.SetMode(gin.TestMode)
	app := gin.New()
	// app.Use(gin.Logger())
	app.Use(gin.Recovery())
	return app
}

func Provides() fx.Option {
	return fx.Provide(
		NewGINApp,
		configs.NewConfig,
	)
}

func Invokes(lifecycle fx.Lifecycle, app *gin.Engine, apiHandler *apis.SpeechHandler, config configs.AppConfig) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				app.GET("/", home)
				apiHandler.RegisterRouterGroup("/models", app)
				go app.Run(":" + "8080")
				return nil
			},
			OnStop: func(context.Context) error {
				return nil
			},
		},
	)
}

func main() {
	fx.New(
		Provides(),
		apis.Module,
		services.Module,
		infrastructure.Module,
		loggerFx.Module,
		fx.Invoke(Invokes),
	).Run()

}
