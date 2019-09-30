package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	pkg "github.com/ja1984/lib-leak/backend/handlers/package"
)

func main() {
	r := gin.New()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	r.Use(cors.New(corsConfig))
	r.Use(errorHandler)
	initializeRoutes(r)

	err := r.Run(fmt.Sprintf(":%v", 9000))
	if err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}

func initializeRoutes(r *gin.Engine) {
	pkg.Routes(r.Group("api"))
}

func errorHandler(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		c.JSON(-1, c.Errors) // -1 == not override the current error code
	}
}
