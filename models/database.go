package models

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

const username string = "root"
const password string = ""
const host string = "localhost"
const port int = 3306
const database string = "project_go_web"

func CreateConnection(){
	if connection, err := sql.Open("mysql",generateURL()) ; err != nil{
		panic(err)
	}else{
		db = connection
		fmt.Println("Conexion exitosa!")
	}
}
func CloseConnection(){
	db.Close()
}
//conexion mysql
//<username>:<password>@tcp(<host>:<port>)/<database>
func generateURL() string{
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",username,password,host,port,database)
}

	
