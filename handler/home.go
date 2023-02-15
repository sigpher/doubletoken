package handler

import "github.com/gin-gonic/gin"

func HomeHandler(ctx *gin.Context) {
	id := ctx.MustGet("id").(uint)
	username := ctx.MustGet("username").(string)
	ctx.JSON(200, gin.H{
		"code": 2000,
		"msg":  "success",
		"data": gin.H{
			"id":       id,
			"username": username,
		},
	})

}
