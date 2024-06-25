package repositories

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewAdminRepository),
	fx.Provide(NewUserRepository),
	fx.Provide(NewRoleRepository),
	fx.Provide(NewCustomerRepository),
)
