package routers

import (
	"go-starter/internal/resources"

	"github.com/gin-gonic/gin"
)

//SetupRouter function will perform all route operations
func SetupRouter() *gin.Engine {
	router := gin.Default()

	userResource := resources.NewUserResource()

	router.Use(func(c *gin.Context) {
		// add header Access-Control-Allow-Origin
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	})

	//API route for health
	router.GET("/health", resources.HealthCheck)

	//API route for users
	users := router.Group("/users")

	users.POST("", userResource.AddUser)

	return router
}
