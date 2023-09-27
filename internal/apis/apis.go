package apis

import (
	"speech-model-hub/internal/domains"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type SpeechHandler struct {
	logger  *zap.SugaredLogger
	repo    domains.IFModelRepository
	service domains.IFRequestHandler
}

type SpeechHandlerParams struct {
	fx.In
	Logger  *zap.SugaredLogger
	Repo    domains.IFModelRepository
	Service domains.IFRequestHandler
}

func NewSpeechHandler(logger *zap.SugaredLogger, repo domains.IFModelRepository, service domains.IFRequestHandler) *SpeechHandler {
	return &SpeechHandler{
		logger:  logger,
		repo:    repo,
		service: service,
	}
}

func (handler *SpeechHandler) GetModelList(c *gin.Context) {
	handler.logger.Info("GetModelList")
	res, err := handler.repo.GetModelList()
	if err != nil {
		handler.logger.Error(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"data": res,
	},
	)

}

func (handler *SpeechHandler) RegisterRouterGroup(path string, app *gin.RouterGroup) {
	group := app.Group(path)
	group.GET("/", handler.GetModelList)
}

var Module = fx.Module(
	"apis",
	fx.Provide(NewSpeechHandler),
)
