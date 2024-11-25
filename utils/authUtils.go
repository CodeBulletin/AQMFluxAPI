package utils

import (
	"fmt"
	"time"

	"math/rand"

	"github.com/codebulletin/AQMFluxAPI/types"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateTokens(username string, tokenSecret string, refreshTokenSecret string, tokenDuration time.Duration, refreshTokenDuration time.Duration) (string, string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, types.TokenClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenDuration)),
		},
	})

	tokenString, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		return "", "", err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256,  types.TokenClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(refreshTokenDuration)),
			ID: 	  "refresh",
		},
	})

	refreshTokenString, err := refreshToken.SignedString([]byte(refreshTokenSecret))
	if err != nil {
		return "", "", err
	}

	return tokenString, refreshTokenString, nil
}

func ValidateToken(tokenString string, secret string, username string) (bool, error) {
	// fmt.Print(secret)
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return false, err
	}

	if !token.Valid {
		return false, fmt.Errorf("token is invalid")
	}

	if claims["username"] != username {
		return false, fmt.Errorf("username does not match")
	}

	return token.Valid, nil
}

func RefreshToken(tokenString string, secret string, duration time.Duration) (string, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", fmt.Errorf("token is invalid")
	}

	// claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(duration))
	claims["exp"] = jwt.NewNumericDate(time.Now().Add(duration))

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	newTokenString, err := newToken.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return newTokenString, nil
}

func GenerateSecrets (length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	secret := make([]byte, length)
	for i := range secret {
		secret[i] = charset[rand.Intn(len(charset))]
	}
	return string(secret), nil
}