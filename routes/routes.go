package routes

import (
	"jumia/controllers"

	"github.com/gin-gonic/gin"
)

func CarregaRotas() {

	router := gin.Default()

	router.GET("/countries", controllers.GetData)
	router.GET("/countries/:country", controllers.GetDataByCountry)
	router.GET("/countries/:country/:state", controllers.GetDataByState)

	router.Run("localhost:8080")
}
