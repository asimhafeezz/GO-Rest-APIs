package helper

import (
	"errors"
	"net/http"

	"fmt"

	"github.com/gin-gonic/gin"

	"restapi/data"
)

func GetTodos(context *gin.Context){
	context.IndentedJSON(http.StatusOK, data.Todos)
}

func AddTodo(context *gin.Context) {
	var newTodo data.Todo

	fmt.Println("before", newTodo)
	
	if err := context.BindJSON(&newTodo); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	
	fmt.Println("after", newTodo)
	data.Todos = append(data.Todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func GetTodo(context *gin.Context){
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

// mini helper
func getTodoById(id string) (*data.Todo, error) {
	for i, t := range data.Todos {
		if t.ID == id {
			return &data.Todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}