package main

import (
	"GolangExampleProject/config"

	"github.com/gin-gonic/gin"
)

func init() {
	configFile := config.ReadYml()
	gin.SetMode(configFile.GetString("server.ginMode"))
}
func main() {
	router := gin.Default()
	v1 := router.Group("v1/")
	v1.POST("/test")
	v1.GET("/test")
	router.Run()

}
