package utils

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type RedisEnv struct {
	RedisAddr     string `yaml:"redisAddr"`
	RedisPassword string `yaml:"redisPassword"`
	RedisDB       int    `yaml:"redisDB"`
}

func NewClient(ctx context.Context) *redis.Client {
	file, e := ioutil.ReadFile("env.yaml")
	if e != nil {
		log.Fatal(e)
	}
	var redisEnv []RedisEnv
	err2 := yaml.Unmarshal(file, &redisEnv)
	if err2 != nil {
		log.Fatal(err2)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisEnv[0].RedisAddr,
		Password: redisEnv[0].RedisPassword,
		DB:       redisEnv[0].RedisDB,
	})

	// Perform basic diagnostic to check if the connection is working
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalln("Redis connection was refused")
	}

	//fmt.Println(pong)
	return rdb
}
