package main

import (
	"fmt"
	"github/harukin721/todo-app/app/models"
)

func main() {
	fmt.Println(models.Db)

	u := &models.User{}
	u.Name = "test3"
	u.Email = "test3@example.com"
	u.Password = "password"
	fmt.Println(u)

	u.CreateUser()

	// user, _ := models.GetUser(2)
	// user.CreateTodo("First Todo")

	// fmt.Println("----- Todo -----")
	// t, _ := models.GetTodo(1)
	// fmt.Println(t)

	user, _ := models.GetUser(3)
	user.CreateTodo("Third Todo")

	// todos, _ := models.GetTodos()
	// for _, todo := range todos {
	// 	fmt.Println(todo)
	// }

	user2, _ := models.GetUser(2)
	todos, _ := user2.GetTodosByUser()
	for _, todo := range todos {
		fmt.Println(todo)
	}
}
