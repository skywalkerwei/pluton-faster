package jwtx

import (
	"context"
	"github.com/golang-jwt/jwt"
)

var CtxKeyJwtUserId = "jwtUserId"   //用户端
var CtxKeyJwtAdminId = "jwtAdminId" //后端

func GetToken(secretKey string, iat, seconds, uid int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[CtxKeyJwtUserId] = uid
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func GetAdminToken(secretKey string, iat, seconds, uid int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[CtxKeyJwtAdminId] = uid
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func GetUid(ctx context.Context) int64 {
	uid, _ := ctx.Value(CtxKeyJwtUserId).(int64)
	return uid
}

func GetAdminUid(ctx context.Context) int64 {
	uid, _ := ctx.Value(CtxKeyJwtAdminId).(int64)
	return uid
}
