package main

import (
	"net/http"
	"errors"
	"github.com/gin-gonic/gin"
)

type todo struct {
	ID string	`json:"id"`
	Item string		`json:"title"`
	Completed bool		`json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Read book", Completed: false},
	{ID: "2", Item: "Workout", Completed: false},
	{ID: "3", Item: "Completed project", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func getTodoById(id string) (*todo, error) {
	for i, t := range todos{
		if t.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("Todo not found")
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message" : "Todo not found"})
		return 
	}

	context.IndentedJSON(http.StatusOK, todo) 
}

func addTodos(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}

func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message" : "Todo not found"})
		return 
	}

	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("todos/:id", toggleTodoStatus)
	router.POST("/todos", addTodos)
	router.Run("localhost:9890")
}