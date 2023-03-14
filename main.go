package main

import (
	"fmt"
	"goAPI/configs"
	"goAPI/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}
	app := gin.Default()

	routes.UserRoutes(app)
	routes.UrlRoutes(app)

	app.Run(fmt.Sprintf("localhost:%s", configs.GetServerPort()))
}
