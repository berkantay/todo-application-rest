package router

import (
	"github.com/berkantay/todo-app-example/internal/todo"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func SetupRouter(todoHandler *todo.Handler) {
	r = gin.Default()
	r.POST("/todo", todoHandler.CreateTodo)
	r.DELETE("/todo", todoHandler.DeleteTodo)
	r.GET("/todo", todoHandler.GetAllTodo)
}

func Run(address string) error {
	return r.Run(address)
}
