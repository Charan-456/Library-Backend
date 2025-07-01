package routes

import (
	"github.com/gorilla/mux"
	"gitub.com/Charan-456/funcs/handlers"
	"gitub.com/Charan-456/funcs/middleware"
)

func Routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/signIn", handlers.SignUp).Methods("POST")
	r.HandleFunc("/AllUsers", handlers.GetAllUserNames).Methods("GET")
	r.HandleFunc("/Login", handlers.Login)
	sub := r.PathPrefix("/api").Subrouter()
	sub.Use(middleware.JwtMiddleware)
	sub.HandleFunc("/welcome", handlers.Books)

	return r
}
