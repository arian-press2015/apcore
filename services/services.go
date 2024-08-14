package services

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewAdminService),
	fx.Provide(NewUserService),
	fx.Provide(NewRoleService),
	fx.Provide(NewCustomerService),
	fx.Provide(NewNotificationService),
	fx.Provide(NewAlbumService),
	fx.Provide(NewMenuService),
	fx.Provide(NewCategoryService),
	fx.Provide(NewProductService),
	fx.Provide(NewIngredientService),
	fx.Provide(NewCommentService),
)
