package main

import (
	"jumia/routes"
)

func main() {

	routes.LoadRoutes()

	routes.LoadRoutes().Run("localhost:8080")

}
