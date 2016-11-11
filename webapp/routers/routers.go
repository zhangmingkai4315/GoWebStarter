package routers

import "github.com/gorilla/mux"

func InitRouters() *mux.Router{
	router:=mux.NewRouter()

	SetStaticFileRouter(router)

	//setting up the pubic access routers like / /about ...
	SetPublicRouter(router)

	//setting up the user routers,like /login /logout and /signup
	SetUserRouter(router)

	//setting up the admin routers like /admin/...
	SetAdminRouter(router)

	//setting up the api routers ,/api/version or add your customer subrouters
	SetApiRouter(router)

	return router
}