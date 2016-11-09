package controllers

import (
	"net/http"
	"time"
	"github.com/gorilla/context"
)

//func LoggingHandler(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
//		start:=time.Now()
//		log.Printf("Start %s %s",r.Method,r.URL.Path)
//		next.ServeHTTP(w,r)
//		log.Printf("Completed %s in %v",r.URL.Path,time.Since(start))
//	})
//}
//
//func AddTimeCostHandler(next http.Handler) http.Handler{
//	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
//		start:=time.Now()
//		next.ServeHTTP(w,r)
//		w.Header().Set("X-Serve-Time",time.Since(start).String())
//		log.Printf("Record time cost in %v",time.Since(start))
//	})
//}


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

