package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hongthai0101/golang-gin/routers"
	"github.com/hongthai0101/golang-gin/utils"
	"io"
	"os"

	_ "github.com/hongthai0101/golang-gin/docs"
)

func outputLog()  {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	outputLog()
	router := routers.InitRoute()
	port := utils.Env("SERVER_PORT", "8080")
	router.Run(":" + port)
}
