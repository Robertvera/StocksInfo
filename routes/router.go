package routes

import (
	"github.com/gorilla/mux"
	"stocksinfo.biz/app"
	"stocksinfo.biz/controllers"
)

func NewRouter(a *app.App) *mux.Router {
	router := mux.NewRouter()

	// Controllers
	etf := controllers.NewETFController(a)

	router.HandleFunc("/fetch/etf/{id}", etf.fetchETFInfo).Methods("GET")

	return router
}
