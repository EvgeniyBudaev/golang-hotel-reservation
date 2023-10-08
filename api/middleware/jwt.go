package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func JWTAuthentication(ctx *fiber.Ctx) error {
	token, ok := ctx.GetReqHeaders()["X-Api-Token"]
	if !ok {
		return fmt.Errorf("unauthorized")
	}
	claims, err := validateToken(token)
	if err != nil {
		return err
	}
	expiresFloat := claims["expires"].(float64)
	expires := int64(expiresFloat)
	// Check token expiration
	if time.Now().Unix() > expires {
		return fmt.Errorf("token expired")
	}
	return ctx.Next()
}

func validateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		secret := os.Getenv("JWT_SECRET")
		return []byte(secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("unauthorized")
	}
	if !token.Valid {
		fmt.Println("invalid token")
		return nil, fmt.Errorf("unauthorized")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("unauthorized")
	}
	return claims, nil
}
