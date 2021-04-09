package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const FoodCollection = "foods"

type Food struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name       string            `bson:"name" json:"name" validate:"required,min=2,max=100"`
	Price      float64           `bson:"price" json:"price" validate:"required"`
	FoodImage  string            `bson:"food_image" json:"foodImage" validate:"required"`
	CreatedAt  time.Time          `bson:"created_at" json:"createdAt"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updatedAt"`
	DeletedAt *time.Time          `bson:"deleted_at" json:"deletedAt"`
}

func CreatingFood(food *Food) *Food {
	food.UpdatedAt = time.Now()
	food.CreatedAt = time.Now()
	food.DeletedAt = nil
	return food
}