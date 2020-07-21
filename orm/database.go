package orm

import (
	_ "github.com/go-sql-driver/mysql"
	"../config"
	"fmt"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func CreateConnection(){
	url := config.GetUrlDatabase()
	if connection, err := gorm.Open("mysql", url) ; err != nil{
		panic(err)
	}else{
		db = connection
		fmt.Println("Conexion exitosa!")
	}
}
func CloseConnection(){
	db.Close()
}

func CreateTables(){
	db.DropTableIfExists(&User{}) //simulamos un TRUNCATE
	db.CreateTable(&User{})
}
