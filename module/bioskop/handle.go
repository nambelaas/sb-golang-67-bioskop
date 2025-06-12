package bioskop

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sb-golang-67-bioskop/helper/api"
	"github.com/sb-golang-67-bioskop/repository"
	"github.com/sb-golang-67-bioskop/service"
)

var (
	bioskopRepo    = repository.NewBioskopRepository()
	bioskopService = service.NewBioskopService(bioskopRepo)
)

func CreateBioskopRouter(ctx *gin.Context) {
	data, err := bioskopService.CreateBioskop(ctx)
	if err != nil {
		api.PrintBadRequestResponse(ctx, err)
		return
	}

	api.PrintSuccessResponseWithData(ctx, "berhasil menambahkan data bioskop", data)
}

func GetAllBioskop(ctx *gin.Context) {
	data, err := bioskopService.GetAllBioskop(ctx)
	if err != nil {
		api.PrintBadRequestResponse(ctx, err)
		return
	}

	api.PrintSuccessResponseWithData(ctx, "berhasil mendapatkan data bioskop", data)
}

func GetBioskopById(ctx *gin.Context) {
	data, err := bioskopService.GetBioskopById(ctx)
	if err != nil {
		api.PrintBadRequestResponse(ctx, err)
		return
	}

	api.PrintSuccessResponseWithData(ctx, fmt.Sprintf("berhasil mendapatkan data bioskop %s", ctx.Param("id")), data)
}

func UpdateBioskop(ctx *gin.Context) {
	err := bioskopService.UpdateBioskop(ctx)
	if err != nil {
		api.PrintBadRequestResponse(ctx, err)
		return
	}

	api.PrintSuccessResponse(ctx, fmt.Sprintf("berhasil update data bioskop %s", ctx.Param("id")))
}

func DeleteBioskop(ctx *gin.Context) {
	err := bioskopService.DeleteBioskop(ctx)
	if err != nil {
		api.PrintBadRequestResponse(ctx, err)
		return
	}

	api.PrintSuccessResponse(ctx, fmt.Sprintf("berhasil hapus data bioskop %s", ctx.Param("id")))
}
