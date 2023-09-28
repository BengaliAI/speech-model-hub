package services

import (
	"speech-model-hub/internal/domains"

	"go.uber.org/fx"
)

var Module = fx.Module(
	"services",
	fx.Provide(
		fx.Annotate(
			NewServiceHandler,
			fx.As(new(domains.IFServices)),
		),
	),
)

func NewServiceHandler() *ServiceHandler {
	return &ServiceHandler{}
}
