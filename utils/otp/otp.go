package otp

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"apcore/utils/keystore"

	"github.com/go-redis/redis/v8"
)

const (
	otpChars = "0123456789"
)

type OTPService struct {
	keyStore *keystore.KeyStore
	length   int
	expiry   time.Duration
}

func NewOTPService(keyStore *keystore.KeyStore, length int, expiry time.Duration) *OTPService {
	return &OTPService{
		keyStore: keyStore,
		length:   length,
		expiry:   expiry,
	}
}

func (s *OTPService) generateOTP() (string, error) {
	otp := make([]byte, s.length)
	for i := 0; i < s.length; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(otpChars))))
		if err != nil {
			return "", fmt.Errorf("failed to generate OTP: %w", err)
		}
		otp[i] = otpChars[index.Int64()]
	}
	return string(otp), nil
}

func (s *OTPService) Generate(ctx context.Context, phone string) (string, error) {
	otp, err := s.generateOTP()
	if err != nil {
		return "", err
	}

	otpKey := fmt.Sprintf("phone-login-otp:%s", phone)
	err = s.keyStore.Set(ctx, otpKey, otp, s.expiry)
	if err != nil {
		return "", fmt.Errorf("failed to store OTP in Redis: %w", err)
	}

	return otp, nil
}

func (s *OTPService) Verify(ctx context.Context, phone, otp string) (bool, error) {
	otpKey := fmt.Sprintf("phone-login-otp:%s", phone)
	storedOTP, err := s.keyStore.Get(ctx, otpKey)
	if err == redis.Nil {
		return false, fmt.Errorf("OTP not found or expired")
	} else if err != nil {
		return false, fmt.Errorf("failed to get OTP from Redis: %w", err)
	}

	if storedOTP != otp {
		return false, nil
	}

	err = s.keyStore.Del(ctx, otpKey)
	if err != nil {
		return false, fmt.Errorf("failed to delete OTP from Redis: %w", err)
	}

	return true, nil
}
