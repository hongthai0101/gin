package repositories

import (
	"github.com/hongthai0101/golang-gin/entity"
)

type UserRepository interface {
	Save(video entity.Video)
	Update(video entity.Video)
	Delete(video entity.Video)
	Find() []entity.Video
	FindById(id string)
	CloseDB()
}

type database struct {
}

func Save(user entity.User) {

}

func FindById(id string)  {

}