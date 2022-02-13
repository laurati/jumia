package routes

import (
	"jumia/controllers"

	"github.com/gin-gonic/gin"
)

func LoadRoutes() *gin.Engine {

	router := gin.Default()

	router.GET("/countries", controllers.GetData)
	router.GET("/countries/:country", controllers.GetDataByCountry)
	router.GET("/countries/:country/:state", controllers.GetDataByState)

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// router.Run("localhost:8080")
	return router
}
