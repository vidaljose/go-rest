package models

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"../config"
)

var db *sql.DB

func CreateConnection(){
	url := config.GetUrlDatabase()
	if connection, err := sql.Open("mysql", url) ; err != nil{
		panic(err)
	}else{
		db = connection
		fmt.Println("Conexion exitosa!")
	}
}

func CreateTables(){
	createTable("users",userSchema)
	existsTable("users")
}

//SI LA TABLA NO EXISTE LA CREA
func createTable(tableName, schema string){
	if !existsTable(tableName){
		Exec(schema)
	}else{
		TruncateTable(tableName)
	}
} 
func TruncateTable(tableName string){
	sql := fmt.Sprintf("TRUNCATE TABLE %s",tableName)
	Exec(sql)
}



func Exec(query string, args ...interface{})(sql.Result,error){
	result, err := db.Exec(query,args...)
	if err != nil{
		log.Println(err)
	}
	return result, err
}
func Query(query string, args ...interface{})(*sql.Rows,error){
	rows, err := db.Query(query,args...)
	if err != nil{
		log.Println(err)
	}
	return rows, err
}

func CloseConnection(){
	db.Close()
}
//verifica si hay conexion
func Ping(){
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

//verifica si la tabla existe 
func existsTable(tableName string) bool{
	//SHOW TABLES LIKE 'users'
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows,_ := Query(sql)
	return rows.Next() //next retorna true si esta bien 
}

	
