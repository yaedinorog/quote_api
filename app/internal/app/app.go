package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yaedinorog/quote_api/app/config"
	"github.com/yaedinorog/quote_api/app/internal/router"
)

func Start() {
	config.LoadConfig("config.yaml")
	router.HTTPSetup()

	log.Print("server started on localhost:8000")
	err := http.ListenAndServe(fmt.Sprintf(":%v", config.Cfg.Server.Port), nil)

	if err != nil {
		log.Fatalf("Server starting error: %v", err)
	}
}
