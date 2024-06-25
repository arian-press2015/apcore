package utils

import (
	"apcore/utils/fileupload"
	"apcore/utils/jwt"
	"apcore/utils/keystore"
	"apcore/utils/otp"
	"apcore/utils/sms"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(jwt.NewJWTService),
	fx.Provide(otp.NewOTPService),
	fx.Provide(keystore.NewKeyStore),
	fx.Provide(fileupload.NewLocalFileUploader),
	fx.Provide(fileupload.NewExtensionValidator),
	fx.Provide(sms.NewSmsSender),
)
