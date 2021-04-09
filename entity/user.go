package entity

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/hongthai0101/golang-gin/utils"
	"time"
)

//User struct is to handle user data
type User struct {
	Email              string `idx:"{email},unique" json:"email" binding:"required,email"`
	Password           string `json:"password" binding:"required"`
	Name               string `json:"name" binding:"required,min=2,max=100"`
	CreatedAt          *time.Time `json:"created_at" default:CURRENT_TIMESTAMP"`
	UpdatedAt          *time.Time `json:"updated_at" default:CURRENT_TIMESTAMP"`
	VerifiedAt *time.Time
}

//GetJwtToken returns jwt token with user email claims
func (user *User) GetJwtToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": string(user.Email),
	})
	secretKey := utils.Env("TOKEN_KEY", "")
	tokenString, err := token.SignedString([]byte(secretKey))
	return tokenString, err
}

type UserList struct {
	User []User `json:"user"`
}

//CreateUser is create new user request model
type CreateUser struct {
	Name  string `json:"name" minLength:"4" example:"test"`
	Email string `json:"email" binding:"required" example:"test@gmail.com"`
}