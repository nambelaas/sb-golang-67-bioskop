package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sb-golang-67-bioskop/repository"
	"github.com/sb-golang-67-bioskop/structs"
)

type BioskopServiceInterface interface {
	CreateBioskop(ctx *gin.Context) (result structs.Bioskop, err error)
	GetAllBioskop(ctx *gin.Context) (result []structs.Bioskop, err error)
	GetBioskopById(ctx *gin.Context) (result structs.Bioskop, err error)
	UpdateBioskop(ctx *gin.Context) (err error)
	DeleteBioskop(ctx *gin.Context) (err error)
}

type BioskopService struct {
	repo repository.BioskopRepositoryInterface
}
