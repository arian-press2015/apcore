package jwt

import (
	"apcore/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService struct {
	jwtKey          []byte
	jwtExpireAt     time.Duration
	refreshExpireAt time.Duration
}

type Claims struct {
	Phone string `json:"phone"`
	jwt.StandardClaims
}

type AuthMessage struct {
	AccessToken  string
	RefreshToken string
}

func NewJWTService(cfg *config.Config) (*JWTService, error) {
	jwtKey := []byte(cfg.Jwt.JwtSecret)
	expireDuration, err := time.ParseDuration(cfg.Jwt.JwtExpireAt)
	if err != nil {
		return nil, err
	}

	refreshExpireDuration, err := time.ParseDuration(cfg.Jwt.RefreshExpireAt)
	if err != nil {
		return nil, err
	}

	return &JWTService{
		jwtKey:          jwtKey,
		jwtExpireAt:     expireDuration,
		refreshExpireAt: refreshExpireDuration,
	}, nil
}

func (s *JWTService) GenerateToken(phone string) (*AuthMessage, error) {
	accessExpireAt := time.Now().Add(s.jwtExpireAt)
	refreshExpireAt := time.Now().Add(s.refreshExpireAt)

	accessToken, err := s.GenerateJWT(phone, accessExpireAt)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.GenerateJWT(phone, refreshExpireAt)
	if err != nil {
		return nil, err
	}

	auth := &AuthMessage{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return auth, nil
}

func (s *JWTService) GenerateJWT(phone string, expireAt time.Time) (string, error) {
	claims := &Claims{
		Phone: phone,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireAt.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *JWTService) VerifyJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return s.jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return claims, nil
}
