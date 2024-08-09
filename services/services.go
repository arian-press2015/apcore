package services

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewAdminService),
	fx.Provide(NewUserService),
	fx.Provide(NewRoleService),
	fx.Provide(NewCustomerService),
	fx.Provide(NewNotificationService),
)
