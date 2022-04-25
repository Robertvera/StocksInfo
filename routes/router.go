package routes

import (
	"github.com/gorilla/mux"
	"stocksinfo.biz/app"
	"stocksinfo.biz/controllers"
)

func NewRouter(a *app.App) *mux.Router {
	router := mux.NewRouter()

	// Controllers
	security := controllers.NewSecurityController(a)

	router.HandleFunc("/fetch/security/{id}", security.FetchSecurityInfo).Methods("GET")
	router.HandleFunc("/save/security/{id}", security.SaveSecurityInfo).Methods("POST")

	return router
}
