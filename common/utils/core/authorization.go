/**
 * @author:       wangxuebing
 * @fileName:     authorization.go
 * @date:         2022/11/5 22:16
 * @description:
 */

package core

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
)

const (
	CtxAuth           = "Authorization" //JWT授权标识
	CtxKeyJwtUserId   = "jwtUserId"     //用户ID
	CtxKeyWxJwtUserId = "jwtWxUserId"   //微信用户ID
)

// 根据ctx解析token中的jwtUserId
func GetJwtUserId(ctx context.Context) string {
	userId := ctx.Value(CtxKeyJwtUserId).(string)
	return userId
}

// 根据ctx解析token中的jwtWxUserId
func GetWxJwtMainUserId(ctx context.Context) string {
	mainUserId := ctx.Value(CtxKeyWxJwtUserId).(string)
	return mainUserId
}

func GeneratorJwtToken(secretKey string, iat, seconds int64, userId string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["iss"] = "lynn"
	claims[CtxKeyJwtUserId] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
