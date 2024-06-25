package mfa

import (
	"github.com/pquerna/otp/totp"
)

func GenerateTOTPSecret(phone string) (string, string, error) {
	secret, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "APCORE",
		AccountName: phone,
	})
	if err != nil {
		return "", "", err
	}

	qrCodeURL := secret.URL()
	return secret.Secret(), qrCodeURL, nil
}

func VerifyTOTPCode(secret, code string) bool {
	return totp.Validate(code, secret)
}
