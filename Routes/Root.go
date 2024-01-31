package Routes

import (
	"github.com/gin-gonic/gin"
	"github.com/souvik666/go_todo/Controllers"
)

func InitializeRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("health", Controllers.GetHealth)
		v1.POST("todo", Controllers.CreateTodo)
		v1.GET("todo", Controllers.GetAllTodo)
		v1.GET("todo/:id", Controllers.GetSingleTodo)
		v1.PATCH("todo/:id", Controllers.UpdateSingleTodo)
		v1.DELETE("todo/:id", Controllers.DeleteSingleTodo)

	}
	return r
}
