package main

import (
	"fmt"

	configFile "testingMongoGB/ConfigFile"
	DentistDAO "testingMongoGB/DAO/Dentist"
	userDAO "testingMongoGB/DAO/User"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/dentists", func(c *gin.Context) {
		var configProps configFile.Config = configFile.ReadConfigFile()

		fmt.Println(configProps)

		c.JSON(200, DentistDAO.GetDentists(configProps.Mongo))
	})

	r.GET("/users", func(c *gin.Context) {
		var configProps configFile.Config = configFile.ReadConfigFile()

		fmt.Println(configProps)

		c.JSON(200, userDAO.GetUsers(configProps.Mongo))
	})
	r.Run(":3000")
}
