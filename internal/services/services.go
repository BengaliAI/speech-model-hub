package services

import (
	"go.uber.org/fx"
)

var Module = fx.Module(
	"services",
	fx.Provide(
		fx.Annotate(
			NewServiceHandler,
			fx.As(new(IFServices)),
		),
	),
)

func NewServiceHandler() *ServiceHandler {
	return &ServiceHandler{}
}
