package routers

import (
	"github.com/gorilla/mux"
	"webapp/controllers"
	"github.com/codegangsta/negroni"
	"log"
	"webapp/common"
)

func SetApiRouter(router *mux.Router) *mux.Router{
	apiRouter:=mux.NewRouter().PathPrefix("/api").Subrouter().StrictSlash(true)

	apiRouter.HandleFunc("/version", controllers.GetApiVersionHandler).Methods("Get")



	//You can add any api routers here, and all will be authenticated after usering negroni wraper.
	//If you decide to show it public ,just modify the next line.

	router.PathPrefix("/api").Handler(
		negroni.New(negroni.HandlerFunc(common.AdminMiddleware),
		negroni.Wrap(apiRouter)))
	log.Println("Loading api interface done.")

	return router;
}