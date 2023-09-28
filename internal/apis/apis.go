package apis

import (
	"speech-model-hub/internal/domains"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type SpeechHandlerParams struct {
	fx.In
	Logger  *zap.SugaredLogger
	Repo    domains.IFModelRepository
	Service domains.IFServices
}

func NewSpeechHandler(logger *zap.SugaredLogger, repo domains.IFModelRepository, service domains.IFServices) *SpeechHandler {
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
