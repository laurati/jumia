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

	//to test routes_test.go
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	return router
}
