package controllers

import (
	"jumia/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//all customers
func GetData(c *gin.Context) {
	customers := models.DataJumia()
	c.IndentedJSON(http.StatusOK, customers)
}

//customer by country
func GetDataByCountry(c *gin.Context) {
	country := c.Param("country")
	i := 0
	for _, a := range models.DataJumia() {
		if strings.EqualFold(a.Country, country) {
			c.IndentedJSON(http.StatusOK, a)
			i++
		}
	}
	if i == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "customer not found"})
	}
}

//customer by country and by state
func GetDataByState(c *gin.Context) {
	state := c.Param("state")
	country := c.Param("country")
	i := 0
	for _, a := range models.DataJumia() {
		if strings.EqualFold(a.Country, country) && strings.EqualFold(a.State, state) {
			c.IndentedJSON(http.StatusOK, a)
			i++
		}
	}
	if i == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "customer not found"})
	}
}
