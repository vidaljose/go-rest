package handlers

import (
	"fmt"
	"net/http"
	//"github.com/gorilla/mux"
	"../models"
	"encoding/json"
)

func GetUsers(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Se obtienen todos los usuarios")
}

func GetUser(w http.ResponseWriter, r *http.Request){
	user := models.User{Id:1,Username:"Jose",Password:"jose123"}
	output ,err := json.Marshal(&user)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Fprintf(w,string(output))
}

func CreateUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Se crea un usuario")
}

func UpdateUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Se actualiza un usuario")
}

func DeleteUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Se elimina un usuario")
}
