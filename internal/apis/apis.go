package apis

import (
	"speech-model-hub/internal/domains"
	"speech-model-hub/internal/services"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type SpeechHandlerParams struct {
	fx.In
	Logger  *zap.SugaredLogger
	Repo    domains.IFModelRepository
	Service services.IFServices
}

func NewSpeechHandler(logger *zap.SugaredLogger, repo domains.IFModelRepository, service services.IFServices) *SpeechHandler {
	return &SpeechHandler{
		logger:  logger,
		repo:    repo,
		service: service,
	}
}

var Module = fx.Module(
	"apis",
	fx.Provide(NewSpeechHandler),
)
