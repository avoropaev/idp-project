package main

import (
	cfg "github.com/vseinstrumentiru/lego/v2/config"
	"github.com/vseinstrumentiru/lego/v2/log/handlers/console"
)

type Config struct {
	App cfg.Application
	Log console.Config
}
