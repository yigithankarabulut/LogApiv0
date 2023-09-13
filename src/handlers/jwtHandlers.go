package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func JwtCreate(str string, c *fiber.Ctx) (error, string) {
	var env []byte = []byte(os.Getenv("SECRET_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Name":      str,
		"ExpiresAt": time.Now().Add(time.Hour * 6).Unix(),
	})
	tokenString, err := token.SignedString(env)
	if err != nil {
		return err, ""
	}
	return nil, tokenString
}

func JwtAccess(tokenString string, adminFlag bool) bool {
	var env []byte = []byte(os.Getenv("SECRET_KEY"))
	var admin string = os.Getenv("ADMIN_KEY")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return env, nil
	})
	if err != nil || !token.Valid {
		return false
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false
	}
	expirationTime := time.Unix(int64(claims["ExpiresAt"].(float64)), 0)
	currentTime := time.Now()
	if currentTime.After(expirationTime) {
		return false
	}
	if adminFlag == true {
		if claims["Name"] != admin {
			return false
		}
	}
	return true
}

func JwtGetUsername(jwtKey string) (string, error) {
	var env []byte = []byte(os.Getenv("SECRET_KEY"))
	token, err := jwt.Parse(jwtKey, func(token *jwt.Token) (interface{}, error) {
		return env, nil
	})
	if err != nil || !token.Valid {
		return "", fmt.Errorf("jwt parse error")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("jwt parse error")
	}
	return claims["Name"].(string), nil
}
