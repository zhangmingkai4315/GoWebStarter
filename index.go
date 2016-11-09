package main

import (
	"controllers"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "gopkg.in/mgo.v2"
)

func NotExistHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "NotExist 404")
}

func main() {
	// 解析输入的参数
	port := flag.String("p", "8000", "listening port")
	// directory := flag.String("d", "public", "static file directory")
	flag.Parse()

	// 创建复用器

	r := mux.NewRouter().StrictSlash(false)

	//设置静态文件目录
	fs := http.FileServer(http.Dir("public"))
	r.Handle("/public", fs)

	// 路由设置
	r.HandleFunc("/login", controllers.GetLoginHandler).Methods("Get")
	r.HandleFunc("/login", controllers.PostLoginHandler).Methods("POST")
	r.HandleFunc("/signup", controllers.GetSignUpHandler).Methods("Get")
	r.HandleFunc("/signup", controllers.PostSignUpHandler).Methods("POST")

	log.Println("Start the web server")

	server := &http.Server{
		Addr:           "localhost:" + *port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(server.ListenAndServe())

}
