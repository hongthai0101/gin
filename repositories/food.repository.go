package repositories

import (
	"context"
	"github.com/hongthai0101/golang-gin/config"
	"github.com/hongthai0101/golang-gin/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type FoodRepository struct {

}

var foodCollection *mongo.Collection = config.OpenCollection(config.Client, entity.FoodCollection)

func (foodRepo FoodRepository) Save(food *entity.Food) (*entity.Food, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	food = entity.CreatingFood(food)
	result, insertErr := foodCollection.InsertOne(ctx, food)
	if insertErr != nil {
		return nil, insertErr
	}
	data := foodCollection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&food)
	if data == nil {
		return nil, data
	}
	return food, nil
}

func (foodRepo FoodRepository) Find(
	where interface{},
	limit *int64,
	offset *int64,
	sort interface{},
	) ([]*entity.Food, error) {
	var foods []*entity.Food

	option := options.FindOptions{
		Limit: limit,
		Skip: offset,
		Sort: sort,
	}
log.Printf("option", *limit, *offset)
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	cursor, err := foodCollection.Find(ctx, where, &option)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &foods)
	if err != nil {
		return nil, err
	}
	return foods, nil
}