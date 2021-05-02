package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"stocksinfo.biz/config"
)

type App struct {
	Config config.Config
}

func New(cfg config.Config) *App {
	return &App{cfg}
}

func (a *App) Run(router *mux.Router) {
	port := a.Config.Port
	addr := fmt.Sprintf(":%v", port)
	fmt.Printf("Listening on port: %d\n", port)
	log.Fatal(http.ListenAndServe(addr, router))
}
