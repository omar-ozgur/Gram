package config

import (
	"github.com/gorilla/mux"
	"github.com/omar-ozgur/gram/app/controllers"
	"github.com/omar-ozgur/gram/middleware"
	"github.com/urfave/negroni"
)

func InitRouter() (n *negroni.Negroni) {
	authorizationHandler := middleware.JWTMiddleware.Handler

	r := mux.NewRouter()

	r.Handle("/signup", controllers.UsersCreate).Methods("POST")
	r.Handle("/login", controllers.UsersLogin).Methods("POST")
	r.Handle("/profile", authorizationHandler(controllers.UsersProfile)).Methods("Get")
	r.Handle("/users", controllers.UsersIndex).Methods("GET")
	r.Handle("/users/search", controllers.UsersSearch).Methods("POST")
	r.Handle("/users/{id}", controllers.UsersShow).Methods("GET")
	r.Handle("/users/{id}", authorizationHandler(controllers.UsersUpdate)).Methods("PUT")
	r.Handle("/users/{id}", authorizationHandler(controllers.UsersDelete)).Methods("DELETE")

	n = negroni.New(negroni.HandlerFunc(middleware.CustomMiddleware), negroni.NewLogger())
	n.UseHandler(r)

	return
}
