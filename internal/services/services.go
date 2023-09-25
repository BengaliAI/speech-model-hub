package services

import (
	"speech-model-hub/internal/domains"

	"go.uber.org/fx"
)

var Module = fx.Module(
	"services",
	fx.Provide(
		fx.Annotate(
			NewInferenceRequestHandler,
			fx.As(new(domains.IFRequestHandler)),
		),
	),
)
