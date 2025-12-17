package responses

import "github.com/gin-gonic/gin"

func JSONSuccess(ctx *gin.Context, msg string) {
	ctx.JSON(200, gin.H{"message": msg})
}

func JSONCreated(ctx *gin.Context, msg string) {
	ctx.JSON(201, gin.H{"message": msg})
}

func JSONData(ctx *gin.Context, code int, obj interface{}) {
	ctx.JSON(code, obj)
}
