package config

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthConfig struct {
	AuthSecret	      	string
	TokenExpiry 		int
}

type JWTClaims struct {
	UserID string `json:"user_id"`
	UserEmail string `json:"user_email"`
	UserName string `json:"user_name"`
	UserCreatedAt int64 `json:"created_at"`
	UserUpdatedAt int64  `json:"updated_at"`
	jwt.RegisteredClaims
}

func NewAuthConfig(env Env) *AuthConfig {
	return &AuthConfig{
		AuthSecret:      	env.AUTH_SECRET,
		TokenExpiry: 		env.TOKEN_EXPIRY,
	}
}

func (a *AuthConfig) GenerateToken(userID string, userName string, userEmail string, UserCreatedAt int64, UserUpdatedAt int64) (string, error) {
	claims := JWTClaims{
		UserID: userID,
		UserEmail: userEmail,
		UserName: userName,
		UserCreatedAt: UserCreatedAt,
		UserUpdatedAt: UserUpdatedAt,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(a.TokenExpiry))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(a.AuthSecret))
}

func (a *AuthConfig) ValidateToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(a.AuthSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}