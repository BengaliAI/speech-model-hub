package main

import (
	"context"
	"speech-model-hub/internal/apis"
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

type Configs struct {
	fx.Out

	MongoDBURL   string `name:"mongo_url"`
	DatabaseName string `name:"database_name"`
}

func NewConfigs() Configs {
	return Configs{
		MongoDBURL:   "mongodb://localhost:27017",
		DatabaseName: "test",
	}
}

func Provide() fx.Option {
	return fx.Provide(
		NewGINApp,
		NewConfigs,
	)
}

func Invoke(lifecycle fx.Lifecycle, app *gin.Engine, apiHandler *apis.SpeechHandler) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				app.GET("/", home)
				apiHandler.RegisterRouterGroup("/models", app)
				go app.Run(":8080")
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
		Provide(),
		apis.Module,
		services.Module,
		infrastructure.Module,
		loggerFx.Module,
		fx.Invoke(Invoke),
	).Run()

}
