package main

import (
	"fmt"
	"./models"
)

func main() {
	models.CreateConnection()
	models.Ping()
	models.CreateTables()
	
	models.CreateUser("Gato1","MIMIU","miau_mi@gmail.com")
	models.CreateUser("Gato2","MIMIU","miau_mi@gmail.com")
	models.CreateUser("Gato3","MIMIU","miau_mi@gmail.com")
	
	//user := models.GetUser(2)
	//fmt.Println(user)
	
	users := models.GetUsers()
	fmt.Println(users)
	
	models.CloseConnection()
}
