package services

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hongthai0101/golang-gin/config"
	"github.com/hongthai0101/golang-gin/entity"
	"github.com/hongthai0101/golang-gin/repositories"
	"github.com/hongthai0101/golang-gin/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

//FoodService is to handle Food relation db query
type FoodService struct{
	repository repositories.FoodRepository
}

var foodCollection = config.OpenCollection(config.Client, entity.FoodCollection)

//Create is to register new Food
func (foodService FoodService) Create(food *entity.Food) (*entity.Food, error) {
	redis, _ := config.RedisInstance()
	err := redis.Client.Set(context.TODO(), config.RedisKeyFood, nil, config.RedisExpirationFood).Err()
	if err != nil {
		log.Printf("foodService cache error %v", err)
	}
	return foodService.repository.Save(food)
}

//Find Food
func (foodService FoodService) Find(c *gin.Context) ([]*entity.Food, error) {
	redis, _ := config.RedisInstance()
	data, err := redis.Client.Get(c, config.RedisKeyFood).Result()

	var result []*entity.Food

	if err != nil {
		log.Printf("redis khong co")
		result, _ = foodService.repository.Find(bson.M{}, utils.GetLimit(c), utils.GetSkip(c), bson.M{})
		dataCache, _ := json.Marshal(result)
		err := redis.Client.Set(c, config.RedisKeyFood, dataCache, 0).Err()
		if err != nil {
			log.Printf("foodService cache error 111 %v", err)
		}
	}else {
		log.Printf("redis da co")
		errUnmarshal := json.Unmarshal([]byte(data), &result)

		if  errUnmarshal != nil {
			log.Printf("errUnmarshal redis khong co")
			result, _  = foodService.repository.Find(bson.M{}, utils.GetLimit(c), utils.GetSkip(c), bson.M{})
			err := redis.Client.Set(c, config.RedisKeyFood, result, config.RedisExpirationFood).Err()
			if err != nil {
				log.Printf("foodService cache error 2222 %v", err)
			}
		}
	}
	return result, nil
}

func (foodService FoodService) Update(food *entity.Food) (*entity.Food, error) {
	var updatedFood *entity.Food

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	update := bson.M{
		"$set": food,
	}

	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		Upsert:         &upsert,
		ReturnDocument: &after,
	}

	err := foodCollection.FindOneAndUpdate(ctx, bson.M{"_id": food.ID}, update, &opt).Decode(&updatedFood)
	if err != nil {
		log.Printf("Could not save Task: %v", err)
		return nil, err
	}
	return updatedFood, nil
}

func (foodService FoodService) Delete(id string) (bool, error) {
	var updatedFood *entity.Food
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}

	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		Upsert:         &upsert,
		ReturnDocument: &after,
	}

	err = foodCollection.FindOneAndUpdate(ctx, bson.M{"_id": oid}, bson.M{
		"$set": bson.M{
			"deleted_at": time.Now(),
		},
	}, &opt).Decode(&updatedFood)
	if err != nil {
		log.Printf("Could not delete food: %v", err)
		return false, err
	}

	return true, nil
}
