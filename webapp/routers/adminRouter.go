package routers

import (
	"github.com/gorilla/mux"
	"webapp/controllers"
	"github.com/codegangsta/negroni"
	"log"
	"webapp/common"
)

func SetAdminRouter(router *mux.Router) *mux.Router{

	authRouter:=mux.NewRouter().PathPrefix("/admin").Subrouter().StrictSlash(true)
	authRouter.HandleFunc("/",controllers.AdminIndexHandler)
	router.PathPrefix("/admin").Handler(
		negroni.New(
			negroni.HandlerFunc(common.AdminMiddleware),
			negroni.Wrap(authRouter)))
	log.Println("Loading admin interface done.")
	return router;
}
