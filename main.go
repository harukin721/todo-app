package main

import (
	"fmt"
	"github/harukin721/todo-app/app/models"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")

	fmt.Println(models.Db)

	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@example.com"
	// u.Password = "password"
	// fmt.Println(u)

	// u.CreateUser()

	fmt.Println("----- Get User -----")
	u, _ := models.GetUser(1)
	fmt.Println(u)

	fmt.Println("----- Update User -----")
	u.Name = "test2"
	u.Email = "test2@example.com"
	u.UpdateUser()
	u, _ = models.GetUser(1)
	fmt.Println(u)
}
