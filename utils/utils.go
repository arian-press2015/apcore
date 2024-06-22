package utils

import (
	"apcore/utils/jwt"
	"apcore/utils/keystore"
	"apcore/utils/otp"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(jwt.NewJWTService),
	fx.Provide(otp.NewOTPService),
	fx.Provide(keystore.NewKeyStore),
)
