package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	secretKey = []byte("simpledouyin")
)

// func main() {
// 	tokenString, err := generateJWT()
// 	if err != nil {
// 		fmt.Println("Error generating JWT:", err)
// 		return
// 	}

// 	fmt.Println("Generated JWT:", tokenString)

// 	claims, err := parseJWT(tokenString)
// 	if err != nil {
// 		fmt.Println("Error parsing JWT:", err)
// 		return
// 	}

// 	fmt.Println("Parsed Claims:", claims)
// }

func GenerateJWT() (string, error) {
	// 创建一个新的 Token
	token := jwt.New(jwt.SigningMethodHS256)

	// 设置 Claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = 123
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// 使用密钥签名 Token 并返回字符串
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseJWT(tokenString string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}

	// 解析 Token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
