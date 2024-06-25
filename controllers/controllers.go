package controllers

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewAuthController),
	fx.Provide(NewUserController),
	fx.Provide(NewRoleController),
	fx.Provide(NewAdminAuthController),
	fx.Provide(NewPingController),
	fx.Provide(NewCustomerController),
	fx.Provide(NewControllers),
)

type Controllers struct {
	AdminAuthController *AdminAuthController
	AuthController      *AuthController
	UserController      *UserController
	RoleController      *RoleController
	PingController      *PingController
	CustomerController  *CustomerController
}

func NewControllers(
	adminAuthController *AdminAuthController,
	authController *AuthController,
	userController *UserController,
	roleController *RoleController,
	pingController *PingController,
	CustomerController *CustomerController,
) *Controllers {
	return &Controllers{
		AdminAuthController: adminAuthController,
		AuthController:      authController,
		UserController:      userController,
		RoleController:      roleController,
		PingController:      pingController,
		CustomerController:  CustomerController,
	}
}
