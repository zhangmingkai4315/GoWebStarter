package main

import (
	"controllers"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
        "github.com/codegangsta/negroni"
	_ "gopkg.in/mgo.v2"
)

func NotExistHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "NotExist 404")
}

func iconHandler(w http.ResponseWriter, r *http.Request) {
}

func main() {
	// 解析输入的参数
	port := flag.String("p", "8000", "listening port")
	// directory := flag.String("d", "public", "static file directory")
	flag.Parse()

	// 创建复用器
	//r := mux.NewRouter().StrictSlash(false)
	//
	//mux:=http.NewServeMux()
	//mux.HandleFunc("/favicon.ico",iconHandler)

	r:=mux.NewRouter()


	//设置静态文件目录
	//fs := http.FileServer(http.Dir("public"))
	//r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", controllers.LoggingHandler(fs)))

	n:=negroni.New()
	n.Use(negroni.HandlerFunc(controllers.StartTimeCostMiddleware))
	n.Use(negroni.NewStatic(http.Dir("/public")))


	// 路由设置
	r.HandleFunc("/login", controllers.GetLoginHandler).Methods("Get")
	r.HandleFunc("/login", controllers.PostLoginHandler).Methods("POST")
	//http.Handle("/login",controllers.AddTimeCostHandler(r))

	r.HandleFunc("/signup", controllers.GetSignUpHandler).Methods("Get")
	r.HandleFunc("/signup", controllers.PostSignUpHandler).Methods("POST")
	n.Use(negroni.HandlerFunc(controllers.EndTimeCostMiddleware))
	n.UseHandler(r)
	log.Println("Start the web server")
	server := &http.Server{
		Addr:           "localhost:" + *port,
		Handler:        n,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(server.ListenAndServe())

}
