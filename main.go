package main

import (
	"github.com/gin-gonic/gin"
	"wibble/database"
	"wibble/request"
	"wibble/service"
)

func Engine(db database.Server) *gin.Engine {
	engine := gin.Default()

	engine.GET("/name", func(c *gin.Context) {
		r, requestErr := request.ParseNameRequest(c)
		if requestErr != nil {
			c.String(400, requestErr.Error())
			return
		}

		name, serviceErr := service.Name(*r, db)
		if serviceErr != nil {
			c.String(404, serviceErr.Error())
		}

		c.String(200, name)
	})

	return engine
}

func main() {
	e := Engine(database.MySqlServer{})
	e.Run(":8080")
}
