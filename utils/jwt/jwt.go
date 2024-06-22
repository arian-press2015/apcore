package jwt

import (
	"apcore/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService struct {
	jwtKey      []byte
	jwtExpireAt time.Duration
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func NewJWTService(cfg *config.Config) (*JWTService, error) {
	jwtKey := []byte(cfg.Jwt.JwtSecret)
	expireDuration, err := time.ParseDuration(cfg.Jwt.JwtExpireAt)
	if err != nil {
		return nil, err
	}

	return &JWTService{
		jwtKey:      jwtKey,
		jwtExpireAt: expireDuration,
	}, nil
}

func (s *JWTService) GenerateJWT(email string) (string, error) {
	expirationTime := time.Now().Add(s.jwtExpireAt)
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
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
