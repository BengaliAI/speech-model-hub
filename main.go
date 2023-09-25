package main

import (
	"speech-model-hub/internal/apis"
	"speech-model-hub/internal/domains"
	"speech-model-hub/internal/infrastructure"
	"speech-model-hub/internal/services"
	"speech-model-hub/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func NewGINApp() *gin.Engine {
	gin.SetMode(gin.TestMode)
	app := gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	return app
}

func home(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}

// func Inintialize() {
// 	app := NewGINApp()
// 	app.GET("/", home)
// 	db := infrastructure.NewDBInstance("mongodb://localhost:27017", "test")
// 	repo := infrastructure.NewModelRepository(db)
// 	service := services.NewInferenceRequestHandler()
// 	aiHandler := apis.NewSpeechHandler(logger.Sugar, repo, service)
// 	aiHandler.RegisterToGroup("/models", app)
// 	// app.GET("/models", aiHandler.GetModelList)
// 	app.Run(":8080")
// }

func Provide() fx.Option {
	return fx.Provide(
		NewGINApp,
		logger.NewLogger,
		infrastructure.NewDBInstance,
		func() infrastructure.DBParams {
			return infrastructure.DBParams{
				MongoURL:     "mongodb://localhost:27017",
				DatabaseName: "test",
			}
		},
		fx.Annotate(
			infrastructure.NewModelRepository,
			fx.As(new(domains.IFModelRepository)),
		),
		fx.Annotate(
			services.NewInferenceRequestHandler,
			fx.As(new(domains.IFRequestHandler)),
		),
		apis.NewSpeechHandler,
	)
}

func InitializeFx() {
	fx.New(
		Provide(),
		fx.Invoke(
			func(app *gin.Engine, aiHandler *apis.SpeechHandler) {
				app.GET("/", home)
				aiHandler.RegisterToGroup("/models", app)
				app.Run(":8080")
			},
		),
	)

}

func main() {
	InitializeFx()
	// Inintialize()
}
