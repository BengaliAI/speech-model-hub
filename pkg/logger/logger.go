package logger

import (
	"time"

	ginzap "github.com/gin-contrib/zap"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var Sugar *zap.SugaredLogger
var Logger *zap.Logger

func Setup(app *gin.Engine) {
	var config zap.Config
	config = zap.NewProductionConfig()
	Logger, _ = config.Build()
	Sugar = Logger.Sugar()
	app.Use(ginzap.Ginzap(Logger, time.RFC3339, true))
}

func TestSetup() {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = nil
	config.DisableStacktrace = true
	Logger, _ := config.Build()
	Sugar = Logger.Sugar()
}
