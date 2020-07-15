package health

import (
	"github.com/gin-gonic/gin"
	"redisapi/utils"
)

func GetHealth(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, utils.Response{
			Message: "healthy",
			Error:   nil,
		})
	})
}