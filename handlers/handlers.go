package handlers

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"../models"
	"strconv"
	"encoding/json"
	//"encoding/xml"
	//"gopkg.in/yaml.v2"
)

func GetUsers(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Se obtienen todos los usuarios")
}

func GetUser(w http.ResponseWriter, r *http.Request){
	//user := models.User{Id:1,Username:"Jose",Password:"jose123"}
	w.Header().Set("Content-type","aplication/json")
	//w.Header().Set("Content-type","text/xml")
	vars := mux.Vars(r)
	userId,_ := strconv.Atoi(vars["id"])
	user,err := models.GetUser(userId)
	if err!=nil{
		w.WriteHeader(http.StatusNotFound)
	}
	output,_:=json.Marshal(&user)
	fmt.Fprintf(w,string(output))
}

func CreateUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Se crea el usuario")
}

func UpdateUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Se actualiza un usuario")
}

func DeleteUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Se elimina un usuario")
}
