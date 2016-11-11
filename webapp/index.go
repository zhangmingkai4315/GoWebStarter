package main

import (
	"webapp/routers"
	"flag"
	"log"
	"net/http"
	"time"
        "github.com/codegangsta/negroni"
	_ "gopkg.in/mgo.v2"
	"webapp/common"
)



func main() {
	// parse the input params
	port := flag.String("p", "8000", "listening port")
	// directory := flag.String("d", "public", "static file directory")
	flag.Parse()

	common.StartUpInit()
	router:=routers.InitRouters()

	//setting up the middleware for global http.
	n:=negroni.Classic()
	n.UseHandler(router)
	log.Printf("Start the web server in port:%v",*port)
	server := &http.Server{
		Addr:           "localhost:" + *port,
		Handler:        n,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(server.ListenAndServe())
}
