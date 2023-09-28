package loggerFx

import (
	"speech-model-hub/internal/configs"
	"time"

	ginzap "github.com/gin-contrib/zap"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Sugar *zap.SugaredLogger
var Logger *zap.Logger

func NewLogger(app *gin.Engine, cfg configs.AppConfig) *zap.SugaredLogger {
	if cfg.IsProduction() {
		Setup(app)
	} else {
		TestSetup(app)
	}
	return Sugar
}

func Setup(app *gin.Engine) {
	var config zap.Config
	config = zap.NewProductionConfig()
	Logger, _ = config.Build()
	Sugar = Logger.Sugar()
	app.Use(ginzap.Ginzap(Logger, time.RFC3339, true))
}

func TestSetup(app *gin.Engine) {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.DisableStacktrace = true
	Logger, _ := config.Build()
	Sugar = Logger.Sugar()
	app.Use(ginzap.Ginzap(Logger, time.RFC3339, true))
}

var Module = fx.Provide(NewLogger)
