package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/hongthai0101/golang-gin/entity"
)

var validate *validator.Validate

//AuthController is for auth logic
type AuthController struct{}

//Login is to process login request
func (auth *AuthController) Login(c *gin.Context) {

	//var loginInfo entity.User
	//if err := c.ShouldBindJSON(&loginInfo); err != nil {
	//	c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
	//	return
	//}
	////TODO
	//userService := services.UserService{}
	//user, errf := userService.Find(&loginInfo)
	//if errf != nil {
	//	c.AbortWithStatusJSON(401, gin.H{"error": "Not found"})
	//	return
	//}
	//
	//err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInfo.Password))
	//if err != nil {
	//	c.AbortWithStatusJSON(402, gin.H{"error": "Email or password is invalid."})
	//	return
	//}
	//
	//token, err := user.GetJwtToken()
	//if err != nil {
	//	c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
	//	return
	//}
	//-------
	c.JSON(200, gin.H{
		"token": "token",
	})
}

//Profile is to provide current user info
func (auth *AuthController) Profile(c *gin.Context) {
	user := c.MustGet("user").(*entity.User)

	c.JSON(200, gin.H{
		"user_name": user.Name,
		"email":     user.Email,
	})
}

//Signup is for user signup
func (auth *AuthController) Signup(c *gin.Context) {
	validate = validator.New()
	type signupInfo struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
		Name     string `json:"name"`
	}
	var info signupInfo

	err := c.ShouldBindJSON(&info)
	if err != nil {
		c.AbortWithStatusJSON(422, gin.H{
			"error": err.Error(),
		})
		return
	}
	//err = validate.Struct(info)
	//if err != nil {
	//	c.AbortWithStatusJSON(423, gin.H{
	//		"error": err.Error(),
	//	})
	//	return
	//}
	//hash, err := bcrypt.GenerateFromPassword([]byte(info.Password), bcrypt.MinCost)
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//user := entity.User{
	//	Email: info.Email,
	//	Password: string(hash),
	//	Name: info.Name,
	//}
	//userService := services.UserService{}
	//err = userService.Create(&user)
	//if err != nil {
	//	c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
	//} else {
	//	c.JSON(200, gin.H{"result": "ok"})
	//}
	return
}
