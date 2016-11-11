package routers

import (
	"github.com/gorilla/mux"
	"webapp/controllers"
	"log"
)

func SetUserRouter(router *mux.Router) *mux.Router{

	router.HandleFunc("/login", controllers.GetLoginHandler).Methods("Get")
	router.HandleFunc("/login", controllers.PostLoginHandler).Methods("POST")

	router.HandleFunc("/logout", controllers.GetLogoutHandler).Methods("Get")

	router.HandleFunc("/signup", controllers.GetSignUpHandler).Methods("Get")
	router.HandleFunc("/signup", controllers.PostSignUpHandler).Methods("POST")
	log.Println("Loading user interface done.")
	return router;
}