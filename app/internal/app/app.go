package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yaedinorog/quote_api/app/config"
	"github.com/yaedinorog/quote_api/app/internal/router"
)

func Start() {
	config.LoadConfig()
	router.HTTPSetup()

	log.Print("server started on localhost:8000")
	http.ListenAndServe(fmt.Sprintf(":%v", config.Cfg.Server.Port), nil)
}
