package main

import (
	"discourse_bench/discourse"
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	Url := os.Getenv("DISCOURSE_URL")
	ApiKey := os.Getenv("DISCOURSE_API_KEY")
	AdminUserName := os.Getenv("DISCOURSE_ADMIN_USERNAME")

	sdk := discourse.Discourse{
		Url:           Url,
		ApiKey:        ApiKey,
		AdminUserName: AdminUserName,
	}

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("sdk", sdk)
		c.Next()
	})

	r.POST("create-user", func(c *gin.Context) {
		var user discourse.CreateDiscourseUserDto
		c.BindJSON(&user)

		sdk := c.MustGet("sdk").(discourse.Discourse)

		response, err := sdk.CreateUser(&user)

		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
		}

		var jsonRes map[string]interface{}
		json.Unmarshal(response, &jsonRes)

		if jsonRes["success"] == false {
			c.JSON(400, jsonRes)
			return
		}

		c.JSON(200, jsonRes)
	})

	r.GET("latest-posts", func(c *gin.Context) {
		sdk := c.MustGet("sdk").(discourse.Discourse)

		response, err := sdk.GetLatestPosts()

		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
		}

		var jsonRes map[string]interface{}
		err1 := json.Unmarshal(response, &jsonRes)

		if err1 != nil {
			c.JSON(500, gin.H{
				"error": err1.Error(),
			})
		}

		c.JSON(200, jsonRes)

	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	r.Run(":5000")
}
