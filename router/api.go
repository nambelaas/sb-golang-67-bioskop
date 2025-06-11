package router

import (
	"github.com/sb-golang-67-bioskop/module/bioskop"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	router.POST("/bioskop", bioskop.CreateBioskopRouter)
}
