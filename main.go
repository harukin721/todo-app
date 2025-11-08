package main

import (
	"fmt"
	"github/harukin721/todo-app/app/models"
)

func main() {
	fmt.Println(models.Db)

	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@example.com"
	// u.Password = "password"
	// fmt.Println(u)

	// u.CreateUser()

	// user, _ := models.GetUser(2)
	// user.CreateTodo("First Todo")

	fmt.Println("----- Todo -----")
	t, _ := models.GetTodo(1)
	fmt.Println(t)
}
