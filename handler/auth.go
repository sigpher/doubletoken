package handler

import (
	"doubletoken/model"
	"doubletoken/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthHandler(ctx *gin.Context) {
	var user model.User
	err := ctx.ShouldBind(&user)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "无效参数",
		})
		return
	}

	if !(user.Username == "choi" && user.ID == 1) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 2002,
			"msg":  "鉴权失败",
		})
		return
	}

	accessTokenString, refreshTokenString := util.GenToken(user.ID, user.Username)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": gin.H{
			"accessToken":  accessTokenString,
			"refreshToken": refreshTokenString,
		},
	})

}
