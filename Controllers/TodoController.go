package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/souvik666/go_todo/Models"
	"github.com/souvik666/go_todo/Repository"
)

func CreateTodo(c *gin.Context) {
	var todo Models.Todo
	c.BindJSON(&todo)
	err := Repository.CreateTodo(&todo)
	if err != nil {

		c.AbortWithStatus(http.StatusExpectationFailed)
	}
	c.JSON(http.StatusOK, &todo)
}

func GetAllTodo(c *gin.Context) {
	var todo []Models.Todo
	err := Repository.GetALlTodo(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusExpectationFailed)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func GetSingleTodo(c *gin.Context) {
	var todo Models.Todo
	id := c.Params.ByName("id")
	err := Repository.GetOneTodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusExpectationFailed)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteSingleTodo(c *gin.Context) {
	var todo Models.Todo
	id := c.Params.ByName("id")
	err := Repository.DeleteOneTodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusExpectationFailed)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func UpdateSingleTodo(c *gin.Context) {
	var todo Models.Todo
	err := Repository.UpdateOneTodo(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusExpectationFailed)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}
