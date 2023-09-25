package infrastructure

import (
	"speech-model-hub/internal/domains"

	"go.uber.org/fx"
)

var Module = fx.Module(
	"infrastructure",
	fx.Provide(
		NewDBInstance,
		fx.Annotate(
			NewModelRepository,
			fx.As(new(domains.IFModelRepository)),
		),
	),
)
