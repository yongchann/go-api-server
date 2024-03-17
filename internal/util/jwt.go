package util

import (
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"go-api-server/config"
	"log"
	"time"
)

func CretaeJwt(info any) string {
	tkn := jwt.New(jwt.SigningMethodHS256)
	claims := tkn.Claims.(jwt.MapClaims)
	claims["info"] = info
	claims["exp"] = time.Now().Add(time.Minute * 20).Unix()

	sTkn, err := tkn.SignedString([]byte(config.Get().Jwt.SecretKey))
	if err != nil {
		log.Println(err)
		return ""
	}
	return sTkn
}

func ParseJwtToken(signedToken string) (map[string]interface{}, error) {
	token, err := jwt.Parse(signedToken, jwtKeyFunc)
	if err != nil {
		return nil, errors.New("failed to parse token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(config.Get().Jwt.SecretKey), nil
}
