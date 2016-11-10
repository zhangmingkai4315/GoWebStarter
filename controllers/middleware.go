package controllers

import (
	"net/http"
	"time"
	"github.com/gorilla/context"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"log"
	"github.com/dgrijalva/jwt-go/request"
)

func StartTimeCostMiddleware(w http.ResponseWriter,r *http.Request,next http.HandlerFunc){
		start:=time.Now()
		context.Set(r,"start_time",start)
		next(w,r)
}

func EndTimeCostMiddleware(w http.ResponseWriter,r *http.Request,next http.HandlerFunc){
	start:=context.Get(r,"start_time")
	w.Header().Set("X-Serve-Time",time.Since(start.(time.Time)).String())
	next(w,r)
}



func AdminMiddleware(w http.ResponseWriter,r *http.Request,next http.HandlerFunc){
	token, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return  verifyKey,nil
	})
	if err!=nil{
		switch err.(type) {
		case *jwt.ValidationError:
			vErr:=err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintln(w,"Token Expired")
				return
			default:
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w,"Error while parsing token!")
				log.Printf("ValidationError:%+v\n",vErr.Errors)
				return
			}
		default:
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w,"Error while parsing token!")
			log.Printf("ValidationError:%+v\n",err)
			return
		}
	}
	if token.Valid{
		next(w,r)
	}else{
		response:=Response{"Invalid token"}
		jsonResponse(response,w)
	}
}

