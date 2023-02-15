package util

import (
	"doubletoken/model"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var accessSecret = []byte("qyrzr")
var refreshSecret = []byte("ar")

const AccessTokenTime = time.Second * 30
const RefreshTokenTime = time.Second * 40

func GenToken(id uint, username string) (string, string) {
	accessClaim := model.CustomClaims{
		ID:       id,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(AccessTokenTime)),
			Issuer:    "ginfo",
		},
	}

	refreshClaim := model.CustomClaims{
		ID:       id,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(RefreshTokenTime)),
			Issuer:    "ginfo",
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaim)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaim)

	accessTokenSigned, err := accessToken.SignedString(accessSecret)
	if err != nil {
		fmt.Println("获取accessToken失败，Secret错误")
		return "", ""
	}

	refreshTokenSigned, err := refreshToken.SignedString(refreshSecret)
	if err != nil {
		fmt.Println("获取refreshToken失败，Secret错误")
		return "", ""
	}

	return accessTokenSigned, refreshTokenSigned
}

func ParseToken(accessTokenString, refreshTokenString string) (*model.CustomClaims, bool, error) {
	accessToken, err := jwt.ParseWithClaims(accessTokenString, &model.CustomClaims{}, func(token *jwt.Token) (any, error) {
		return accessSecret, nil
	})
	if err != nil {
		return nil, false, nil
	}
	if claims, ok := accessToken.Claims.(*model.CustomClaims); ok && accessToken.Valid {
		return claims, false, nil
	}

	refreshToken, err := jwt.ParseWithClaims(refreshTokenString, &model.CustomClaims{}, func(token *jwt.Token) (any, error) {
		return refreshSecret, nil
	})
	if err != nil {
		return nil, false, err
	}

	if claims, ok := refreshToken.Claims.(*model.CustomClaims); ok && refreshToken.Valid {
		return claims, true, nil
	}
	return nil, false, errors.New("invalid token")

}
