package redis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"redisapi/utils"
)

func AddKeyToRedis(r *gin.Engine) {
	r.GET("/add/:key/:value", func(c *gin.Context) {
		redisClient := utils.GetRedis(c)
		r, err := addKey(redisClient, c.Param("key"), c.Param("value"))
		c.JSON(200, utils.Response{
			Message: r,
			Error:   err,
		})
	})
}

func DeleteKeyFromRedis(r *gin.Engine) {
	r.GET("/delete/:key", func(c *gin.Context) {
		redisClient := utils.GetRedis(c)
		r, err := deleteKey(redisClient, c.Param("key"))
		c.JSON(200, utils.Response{
			Message: r,
			Error:   err,
		})
	})
}

func GetKeyFromRedis(r *gin.Engine) {
	r.GET("/get/:key", func(c *gin.Context) {
		redisClient := utils.GetRedis(c)
		r, err := getKey(redisClient, c.Param("key"))
		c.JSON(200, utils.Response{
			Message: r,
			Error:   err,
		})
	})
}

func addKey(redisClient *redis.Client, k, v string) (interface{}, error) {
	return redisClient.Set(utils.GenerateContext(), k, v, 0).Result()
}

func deleteKey(redisClient *redis.Client, k string) (interface{}, error) {
	return redisClient.Del(utils.GenerateContext(), k).Result()
}

func getKey(redisClient *redis.Client, k string) (interface{}, error) {
	return redisClient.Get(utils.GenerateContext(), k).Result()
}
