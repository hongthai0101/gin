package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/hongthai0101/golang-gin/entity"
	"github.com/hongthai0101/golang-gin/services"
	"net/http"

	_ "github.com/hongthai0101/golang-gin/types"
)

// create a validator object
var validateFoo = validator.New()

//AuthController is for auth logic
type FoodController struct{}

// ListUser is the handler of list user endpoint
// @Summary List users
// @Description list all the users based on filter given
// @Tags user
// @Produce  json
// @Param q query string true "q"
// @Success 200 {object} entity.UserList
// @Router /users/ [get]
func (c *FoodController) Save(ctx *gin.Context) {
	var food entity.Food
	//bind the object that comes in with the declared varaible. thrrow an error if one occurs
	if err := ctx.BindJSON(&food); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// use the validation package to verify that all items coming in meet the requirements of the struct
	if err := validateFoo.Struct(food); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	service := services.FoodService{}

	service.Create(&food)
	ctx.JSON(200, food)
	return
}

// CreateUser is create user endpoint handler
// @Summary create user
// @Description create a new user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body entity.CreateUser true "create user"
// @Success 200 {object} entity.User
// @Failure 400 {object} types.Swagger.HTTPError
// @Router /user/ [post]
func (c *FoodController ) Find(ctx *gin.Context)  {
	service := services.FoodService{}
	foods, _ := service.Find(ctx)
	ctx.JSON(200, foods)
	return
}

func (c *FoodController ) Delete(ctx *gin.Context)  {
	service := services.FoodService{}
	service.Delete(ctx.Param("id"))
	ctx.JSON(204, nil)
	return
}

