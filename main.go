package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/souvik666/go_todo/Config"
	"github.com/souvik666/go_todo/Models"
	"github.com/souvik666/go_todo/Routes"
)

func main() {
	Config.MakeDBConnection()
	r := Routes.InitializeRouter()

	Config.DB.AutoMigrate(&Models.Todo{})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
