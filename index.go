package main

import (
	"controllers"
	"flag"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
        "github.com/codegangsta/negroni"
	_ "gopkg.in/mgo.v2"
)



func main() {
	// 解析输入的参数
	port := flag.String("p", "8000", "listening port")
	// directory := flag.String("d", "public", "static file directory")
	flag.Parse()


	router:=mux.NewRouter()
	authRouter:=mux.NewRouter().PathPrefix("/admin").Subrouter().StrictSlash(true)
	authRouter.HandleFunc("/",controllers.AdminIndexHandler)
	router.PathPrefix("/admin").Handler(
		negroni.New(
			negroni.HandlerFunc(controllers.AdminMiddleware),
			negroni.Wrap(authRouter)))

	//设置静态文件目录
	fs := http.FileServer(http.Dir("public"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/",fs))

	n:=negroni.New()
	n.Use(negroni.HandlerFunc(controllers.StartTimeCostMiddleware))


	// 路由设置
	router.HandleFunc("/login", controllers.GetLoginHandler).Methods("Get")
	router.HandleFunc("/login", controllers.PostLoginHandler).Methods("POST")

	router.HandleFunc("/logout", controllers.GetLogoutHandler).Methods("Get")
	//http.Handle("/login",controllers.AddTimeCostHandler(r))

	router.HandleFunc("/signup", controllers.GetSignUpHandler).Methods("Get")
	router.HandleFunc("/signup", controllers.PostSignUpHandler).Methods("POST")
	n.Use(negroni.HandlerFunc(controllers.EndTimeCostMiddleware))
	n.UseHandler(router)


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
