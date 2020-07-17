package models 

import (
	"net/http"
	"fmt"
	"encoding/json"
	)
type Response struct{
	Status int `json:"status"`
	Data interface{} `json:"data"`
	Message string `json:"message"`
	contentType string
	writer http.ResponseWriter
}

func GetDefaultResponse(w http.ResponseWriter) Response{
	return Response{Status:http.StatusOK ,writer:w,contentType:"application/json"}
}

func (this *Response) NotFound(message string){
	this.Status = http.StatusNotFound
	this.Data = nil
	this.Message =message 
}
func (this *Response) Send(){
	this.writer.Header().Set("Content-Type",this.contentType)
	this.writer.WriteHeader(this.Status)

	output,_:=json.Marshal(&this)
	fmt.Fprintf(this.writer,string(output))
}
