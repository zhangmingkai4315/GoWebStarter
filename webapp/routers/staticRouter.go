package routers

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
)

func SetStaticFileRouter(router *mux.Router) *mux.Router{
	//setting the static file dir to public and access by /static/ prefix!
	fs := http.FileServer(http.Dir("public"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/",fs))
	log.Println("Loading static interface done.")
	return router

}