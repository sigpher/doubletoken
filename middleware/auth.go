package middleware

import (
	"doubletoken/util"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		if authHeader == "" {
			ctx.JSON(200, gin.H{
				"code": 2003,
				"msg":  "请求头中auth为空",
				"data": errors.New("请求头中auth为空"),
			})
			ctx.Abort()
			return

		}
		parts := strings.Split(authHeader, " ")

		if !(len(parts) == 3 && parts[0] == "Bearer") {
			ctx.JSON(200, gin.H{
				"code": 2004,
				"msg":  "请求头中auth格式有误",
				"data": errors.New("请求头中auth格式有误"),
			})
			ctx.Abort()
			return
		}

		parseToken, needUpdate, err := util.ParseToken(parts[1], parts[2])
		if err != nil {
			ctx.JSON(200, gin.H{
				"code": 2005,
				"msg":  "无效的Token",
				"data": err,
			})
			ctx.Abort()
			return
		}
		if needUpdate {
			parts[1], parts[2] = util.GenToken(parseToken.ID, parseToken.Username)
			ctx.JSON(200, gin.H{
				"code": 200,
				"msg":  "鉴权成功",
				"data": gin.H{
					"accessToken":  parts[1],
					"refreshToken": parts[2],
				},
			})

		}

		ctx.Set("id", parseToken.ID)
		ctx.Set("username", parseToken.Username)
		ctx.Next()
	}
}
