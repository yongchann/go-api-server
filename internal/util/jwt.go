package util

import (
	"github.com/golang-jwt/jwt"
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
