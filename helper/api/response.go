package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PrintBadRequestResponse(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"Success": false,
		"Message": err.Error(),
	})
}

func PrintSuccessResponse(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Message": msg,
	})
}

func PrintSuccessResponseWithData(ctx *gin.Context, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Message": msg,
		"Data":    data,
	})
}
