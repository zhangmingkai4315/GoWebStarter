package routers

import (
	"github.com/gorilla/mux"
	"webapp/controllers"
	"log"
)

func SetPublicRouter(router *mux.Router) *mux.Router{

	router.HandleFunc("/", controllers.GetIndexPageHandler).Methods("Get")
	//router.HandleFunc("/about", controllers.GetAboutPageHandler).Methods("GET")

	log.Println("Loading public interface done.")
	return router;
}