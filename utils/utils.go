package utils

import (
	"apcore/config"
	"apcore/utils/fileupload"
	"apcore/utils/jwt"
	"apcore/utils/keystore"
	"apcore/utils/otp"
	"apcore/utils/sms"
	"time"

	"go.uber.org/fx"
)

const (
	OTP_LENGTH         = 6
	OTP_EXPIRY_MINUTES = 2 * time.Minute
)

var Module = fx.Options(
	fx.Provide(jwt.NewJWTService),
	fx.Provide(provideRedisAddress),
	fx.Provide(keystore.NewKeyStore),
	fx.Provide(fileupload.NewLocalFileUploader),
	fx.Provide(fileupload.NewExtensionValidator),
	fx.Provide(sms.NewSmsSender),
	fx.Provide(
		provideOTPLength,
		provideOTPExpiry,
	),
	fx.Provide(otp.NewOTPService),
)

func provideRedisAddress(config *config.Config) string {
	return config.Redis.Url
}

func provideOTPLength() int {
	return OTP_LENGTH
}

func provideOTPExpiry() time.Duration {
	return OTP_EXPIRY_MINUTES
}
