package jwtx

import (
	"context"
	"github.com/golang-jwt/jwt"
)

func GetToken(secretKey string, iat, seconds, uid int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["uid"] = uid
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func GetUid(ctx context.Context) int64 {
	uid, _ := ctx.Value("uid").(int64)
	return uid
}
