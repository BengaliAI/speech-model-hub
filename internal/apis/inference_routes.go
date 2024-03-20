package apis

import (
	"mime/multipart"
	"net/http"
	"speech-model-hub/internal/domains"
	"speech-model-hub/internal/services"
	"speech-model-hub/pkg/etypes"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SpeechHandler struct {
	logger  *zap.SugaredLogger
	repo    domains.IFModelRepository
	service services.IFServices
}

func (handler *SpeechHandler) GetModelList(c *gin.Context) {
	handler.logger.Info("GetModelList")
	res, err := handler.repo.GetModelList()
	if err != nil {
		handler.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"data": res,
	},
	)
}

type InferenceRequest struct {
	DisplayName string                `binding:"required" form:"display_name"`
	File        *multipart.FileHeader `binding:"required" form:"file"`
}

func (handler *SpeechHandler) GetInference(c *gin.Context) {
	handler.logger.Info("GetInference")
	body := InferenceRequest{}
	if err := c.Bind(&body); err != nil {
		handler.logger.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	savedFilePath, err := handler.service.SaveFile(body.File)
	if err != nil {
		handler.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	handler.logger.Info(savedFilePath)
	aiModel, err := handler.repo.GetModelByDisplayName(body.DisplayName)
	if err != nil {
		handler.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": etypes.ERROR_MODEL_NOT_FOUND})
		return
	}
	handler.logger.Info(aiModel)
	inference, err := handler.service.SendRequest(aiModel, body.File)
	if err != nil {
		handler.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": etypes.ERROR_INFERENCE_FAILED})
		return
	}
	c.JSON(200, gin.H{"transcript": inference})
}

func (handler *SpeechHandler) RegisterRouterGroup(path string, app *gin.RouterGroup) {
	group := app.Group(path)
	group.GET("/", handler.GetModelList)
	group.POST("/inference", handler.GetInference)
}
