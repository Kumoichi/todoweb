package main

import (
	"fmt"
	"todo_app/app/models"
)

func main() {
	//package instance
	//Config構造体にアクセス

	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")

	fmt.Println(models.Db)

}
