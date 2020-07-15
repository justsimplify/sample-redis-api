package main

import (
	"log"
	"redisapi/health"
	"redisapi/redis"
	"redisapi/utils"
)

func main() {
	r := utils.GinRouter()
	health.GetHealth(r)
	redis.AddKeyToRedis(r)
	redis.DeleteKeyFromRedis(r)
	redis.GetKeyFromRedis(r)

	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}