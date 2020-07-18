package handlers 

import (
	//"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"../models"
	"strconv"
	"encoding/json"
	//"encoding/xml"
	//"gopkg.in/yaml.v2"
)

func GetUsers(w http.ResponseWriter, r *http.Request){
	models.SendData(w,models.GetUsers())
}

func GetUser(w http.ResponseWriter, r *http.Request){

	if user, err := getUserByRequest(r);err!=nil{
		models.SendNotFound(w) //nuevo
	}else{
		models.SendData(w,user) //nuevo
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request){
	user := models.User{}
	decoder := json.NewDecoder(r.Body)
	
	if err:=decoder.Decode(&user);err != nil{
		models.SendUnprocessableEntity(w)
	}else{
		models.SendData(w,models.SaveUser(user))
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request){
	user, err := getUserByRequest(r)
	
	if err != nil {
		models.SendNotFound(w)
		return 
	}
	userResponse := models.User{}
	decoder := json.NewDecoder(r.Body)
	
	if err := decoder.Decode(&userResponse);err != nil{
		models.SendUnprocessableEntity(w)
		return
	}
	
	user = models.UpdateUser(user, userResponse.Username,userResponse.Password)
	models.SendData(w,user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request){
	if user,err := getUserByRequest(r);err != nil{
		models.SendNotFound(w)
	}else{
		models.DeleteUser(user.Id)
		models.SendNoContent(w)
		
	}
}

func getUserByRequest(r *http.Request) (models.User,error){
	vars := mux.Vars(r)
	userId,_ := strconv.Atoi(vars["id"])
	if user,err := models.GetUser(userId); err != nil{
		return user,err
	}else{
		return user, nil
	}
}
