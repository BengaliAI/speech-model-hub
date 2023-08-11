package main

import "github.com/gin-gonic/gin"

func constructGINApp() *gin.Engine {
	gin.SetMode(gin.TestMode)
	app := gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	return app
}

func main() {
	app := constructGINApp()
	app.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})
	app.Run(":8080")
}
