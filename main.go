package main

import (
	"log"

	"stocksinfo.biz/app"
	"stocksinfo.biz/config"
	"stocksinfo.biz/routes"
)

func main() {
	cfg, err := config.New("config/app.json")
	if err != nil {
		log.Fatal(err)
	}
	app := app.New(cfg)
	router := routes.NewRouter(app)
	app.Run(router)
}
