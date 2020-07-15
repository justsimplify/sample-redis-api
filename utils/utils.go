package utils

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"log"
)

var ginEngine *gin.Engine
var redisClient *redis.Client

func GinRouter() *gin.Engine {
	if ginEngine == nil {
		fmt.Println("Creating new router ...")
		ginEngine = gin.Default()
		ginEngine.RedirectTrailingSlash = false
		ginEngine.RemoveExtraSlash = true
	}
	return ginEngine
}

func GenerateContext() context.Context {
	return context.Background()
}

func GetRedis(c *gin.Context) *redis.Client {
	redisHost := c.GetHeader("redis_host")
	redisPort := c.GetHeader("redis_port")
	if redisClient == nil {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     redisHost + ":" + redisPort,
			Password: "", // no password set
			DB:       0,  // use default DB
		})
	}
	_, err := redisClient.Ping(GenerateContext()).Result()
	if err != nil {
		log.Fatalf("redis client error: %s", err)
	}
	return redisClient
}