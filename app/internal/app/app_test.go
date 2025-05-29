package app

import (
	"testing"

	"github.com/yaedinorog/quote_api/app/config"
)

func TestLoadConfig(t *testing.T) {
	config.LoadConfig("config.yaml")

	if config.Cfg == nil {
		t.Error("config didnt load")
	}
	if config.Cfg.Server.Host != "localhost" {
		t.Error("config loaded wrong or host set as not local host")
	}
}
