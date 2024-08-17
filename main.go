package main

import (
	"another_todo_app/app/controllers"
	"another_todo_app/app/models"
	"fmt"
)

func main() {

	// User構造体のName,Email、Passwordに値を入れて,これらの値をCreateUserに入れる
	fmt.Println(models.Db)
	controllers.StartMainServer()
}
