package main

import (
	"fmt"
	"./orm"
)

func main() {
	orm.CreateConnection()
	orm.CreateTables()
	
	user := orm.NewUser("Gato","123","gato@gmail.com")
	user.Save()
	
	//users := orm.GetUsers()
	//fmt.Println(users)
	
	//user = orm.GetUser(1)

	user.Username = "Prueba"
	user.Password = "pass"
	user.Email = "email@gmail.com"

	user.Save()
	user.Delete()
	fmt.Println(user)
	
	orm.CloseConnection()
}
