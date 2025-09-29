package service

import (
	"github.com/Nezent/mig-test/internal/domain/user"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"service",
	fx.Provide(
		NewUserService,
		fx.Annotate(
			NewUserService,
			fx.As(new(user.UserService)),
		),
	),
)
