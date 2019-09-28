package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func main() {
	r := gin.New()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	r.Use(cors.New(corsConfig))
	initializeRoutes(r)

	err := r.Run(fmt.Sprintf(":%v", 9000))
	if err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}

func initializeRoutes(r *gin.Engine) {
	apiGroup := r.Group("api")
	apiGroup.GET("/repo/", func(c *gin.Context) {
		client := resty.New()
		resp, err := client.R().
			EnableTrace().
			Get(c.Query("repoUrl"))

		if err != nil {
			c.Status(500)
			return
		}

		c.JSON(200, resp)
	})
}
