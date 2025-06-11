package main

import (
	"fmt"
	"github.com/sb-golang-67-bioskop/config"
	"github.com/sb-golang-67-bioskop/database/connection"
	"github.com/sb-golang-67-bioskop/router"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	config.Init()

	connection.Init()
	defer connection.DbConn.Close()

	InitRouter()
}

func InitRouter() {
	web := gin.Default()

	router.Init(web)

	port := fmt.Sprintf(":%d", viper.GetInt("app.port"))

	web.Run(port)
}
