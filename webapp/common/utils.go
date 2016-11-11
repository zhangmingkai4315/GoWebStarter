package common

import (
	"os"
	"log"
	"encoding/json"
	"net/http"
)

type configuration struct {
	Database databaseConfig `json:"database"`
}
type databaseConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
	User string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}
var AppConfig configuration
func initConfig(){
	file,err:=os.Open("config.json")
	defer file.Close()
	if err!=nil{
		log.Fatalf("Error:Open config file fail %s\n",err)
	}
	decoder:=json.NewDecoder(file)
	AppConfig=configuration{}
	err=decoder.Decode(&AppConfig)
	if err!=nil{
		log.Fatalf("Error:Loading config file fail %s\n",err)
	}
	log.Println("Loading config file done.")
}

type ResponseMessage struct {
	Error string `json:"error"`
	Message string `json:"message"`
	HttpStatus int `json:"status"`
	Data interface{} `json:"data"`
}

type ResponseObj struct {
	Data ResponseMessage `json:"data"`
	Status bool `json:status`
}

func DisplayAppError(w http.ResponseWriter, handlerError error, message string,code int){
	var errorMessage string
	if handlerError==nil{
		errorMessage=""
	}else{
		errorMessage=handlerError.Error()
	}
	errObj:=ResponseMessage{
		Error:errorMessage,
		Message:message,
		HttpStatus:code,
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	if response,err:=json.Marshal(ResponseObj{Data:errObj,Status:false});err==nil{
		w.Write(response)
	}
}

func DisplayAppData(w http.ResponseWriter, data interface{}){
	messageObj:=ResponseMessage{
		Error: "",
		Message:"",
		HttpStatus:200,
		Data:data,
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(200)

	if response,err:=json.Marshal(ResponseObj{Data:messageObj,Status:true});err==nil{
		w.Write(response)
	}else{
		DisplayAppError(w,err,"object turn to json fail",500)
	}
}
