package config

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"log"
)
type Database struct {
	Client *redis.Client
}

var (
	ErrNil = errors.New("no matching record found in redis database")
	Ctx    = context.TODO()
)

//RedisInstance func
func RedisInstance() (*Database, error) {

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	if err := client.Ping(Ctx).Err(); err != nil {
		log.Printf("Connect redis error")
		return nil, err
	}
	return &Database{
		Client: client,
	}, nil
}