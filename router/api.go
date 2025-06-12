package router

import (
	"github.com/sb-golang-67-bioskop/module/bioskop"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	{
		bioskopRoute := router.Group("/bioskop")
		bioskopRoute.POST("/", bioskop.CreateBioskopRouter)
		bioskopRoute.GET("/", bioskop.GetAllBioskop)
		bioskopRoute.GET("/:id", bioskop.GetBioskopById)
		bioskopRoute.PUT("/:id", bioskop.UpdateBioskop)
		bioskopRoute.DELETE("/:id", bioskop.DeleteBioskop)
	}
}
