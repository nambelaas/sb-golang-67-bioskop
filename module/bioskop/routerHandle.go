package bioskop

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBioskopRouter(ctx *gin.Context) {
	var (
		bioskopRepo    = NewRepository()
		bioskopService = NewService(bioskopRepo)
	)

	data, err := bioskopService.CreateBioskop(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": err.Error(),
			"Data":    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Message": "berhasil menambahkan data bioskop",
		"Data":    data,
	})
}
