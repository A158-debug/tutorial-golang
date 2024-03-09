package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type todo struct {
	ID        string `json:id`
	Item      string `json:item`
	Completed bool   `json:completed`
}

var Todo_List = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Record videos", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, Todo_List)
}

func addTodos(context *gin.Context) {
	var newTodo todo
	err := context.BindJSON(&newTodo)
	if err != nil {
		return
	}

	Todo_List = append(Todo_List, newTodo)
	context.IndentedJSON(http.StatusCreated, Todo_List)
}

func getTodoById(context *gin.Context) {
	id := context.Param("id")
	for idx, item := range Todo_List {
		if item.ID == id {
			context.IndentedJSON(http.StatusOK, Todo_List[idx])
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, errors.New("todo not found"))
}

func main() {
	fmt.Println("Web serve in Go")
	router := gin.Default()

	router.GET("/", getTodos)
	router.GET("/:id", getTodoById)
	router.POST("/", addTodos)

	router.Run("localhost:8080")
}
