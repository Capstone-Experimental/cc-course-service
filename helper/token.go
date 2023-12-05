package helper

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func VerifyToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, err
	}
	return claims, nil
}

func CheckAuthorization(c *fiber.Ctx) string {
	header := c.Get("Authorization")
	if header == "" {
		return "Unauthorized 4"
	}
	parts := strings.Split(header, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "Unauthorized 5"
	}
	token := parts[1]
	return token
}

func GetIDFromToken(tokenString string) (string, error) {
	claims, err := VerifyToken(tokenString)
	if err != nil {
		return "", err
	}
	return claims.StandardClaims.Id, nil
}

func GetUsernameFromToken(tokenString string) (string, error) {
	claims, err := VerifyToken(tokenString)
	if err != nil {
		return "", err
	}
	return claims.Username, nil
}
