package routes

import (
	"github.com/Charan-456/funcs/handlers"
	"github.com/Charan-456/funcs/middleware"
	"github.com/gorilla/mux"
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
