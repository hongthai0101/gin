package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hongthai0101/golang-gin/controllers"
	"github.com/hongthai0101/golang-gin/middlewares"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setAuthRoute(router *gin.Engine) {
	authController := new(controllers.AuthController)
	router.POST("/login", authController.Login)
	router.POST("/signup", authController.Signup)

	authGroup := router.Group("/")
	authGroup.GET("/profile", authController.Profile)
}

func setFoodRoute(router *gin.Engine) {
	foodController := new(controllers.FoodController)
	router.POST("/foods", foodController.Save)
	router.GET("/foods", foodController.Find)
	router.DELETE("/foods/:id", foodController.Delete)
}

// InitRoute ..
func InitRoute() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Use(gin.Recovery(), middlewares.Authentication())
	setAuthRoute(router)
	setFoodRoute(router)
	return router
}