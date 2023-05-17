package main

import (
	"restapi/helper"

	"github.com/gin-gonic/gin"
)



func main() {
	router := gin.Default()
	router.GET("/todos", helper.GetTodos)
	router.GET("/todos/:id", helper.GetTodo)
	router.POST("/todos", helper.AddTodo)
	router.Run("localhost:9090")
}