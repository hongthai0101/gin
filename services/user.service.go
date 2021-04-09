package services

import (
	"github.com/hongthai0101/golang-gin/entity"
)

//UserService is to handle user relation db query
type UserService struct{}

//Create is to register new user
func (userService UserService) Create(user *entity.User)  {

}

// Delete a user from DB
func (userService UserService) Delete(email string)  {

}

//Find user
func (userService UserService) Find(user *entity.User)  {

}

//Find user from email
func (userService UserService) FindByEmail(email string) {
	return
}
