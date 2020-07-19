package main

import (
	"./models"
	"fmt"
)

func main() {
	models.CreateConnection()
	models.Ping()
	result := models.ExistsTable("users")
	fmt.Println(result)
	models.CloseConnection()
}
