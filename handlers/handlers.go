package handlers 

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"../models"
	"strconv"
	//"encoding/json"
	//"encoding/xml"
	//"gopkg.in/yaml.v2"
)

func GetUsers(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Se obtienen todos los usuarios")
}

func GetUser(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	userId,_ := strconv.Atoi(vars["id"])
	user,err := models.GetUser(userId)
	response := models.GetDefaultResponse(w)
	if err!=nil{
		response.NotFound(err.Error())
		w.WriteHeader(http.StatusNotFound)
	}else{
		response.Data = user
	}
	response.Send()
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
