package main

import (
	"jumia/routes"
)

func main() {

	r := routes.LoadRoutes()

	r.Run("localhost:8080")

}
