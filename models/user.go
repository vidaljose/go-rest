package models 

import "errors"

type User struct{
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

const userSchema string = `CREATE TABLE users(
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(30) NOT NULL,
	password VARCHAR(64) NOT NULL,
	email VARCHAR(40),
	created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`

type Users []User //supongo que estructura para lista de objetos

var users = make(map[int]User)

func SetDefaultUser(){
	user := User{1,"JoseVidal","123"}
	users[user.Id] = user
}

func GetUser(userId int)(User,error){
	if user,ok:=users[userId];ok{
		return user,nil
	}
	return User{}, errors.New("No se encuentra el usuario")
}

//retornar todos los usuarios
func GetUsers() Users{
	list := Users{}
	for _,user := range users{
		list = append(list,user)
	}
	return list
}

func SaveUser(user User) User{
	user.Id = len(users) + 1
	users[user.Id] = user
	return user
}

func UpdateUser(user User,username string,password string) User{
	user.Username = username
	user.Password = password
	users[user.Id] = user
	return user
}

func DeleteUser(id int){
	delete(users, id) //primer parametro mapa, segundo llave
}
