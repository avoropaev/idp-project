package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	legoapp "github.com/vseinstrumentiru/lego/v2/app"
	cfg "github.com/vseinstrumentiru/lego/v2/config"
	"github.com/vseinstrumentiru/lego/v2/di"
	"github.com/vseinstrumentiru/lego/v2/log/handlers/console"
	"github.com/vseinstrumentiru/lego/v2/module"
	"logur.dev/logur"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	ctx := context.Background()

	legoapp.NewRuntime(
		legoapp.CommandMode(),
		legoapp.Provide(
			module.Pipeline,
			di.ProvideCommand(NewCommand(ctx)),
		),
		legoapp.WithConfig(&Config{
			App: cfg.Application{
				Name:       os.Getenv("APP_NAME"),
				DataCenter: os.Getenv("DATA_CENTER"),
				DebugMode:  stringEnvToBool(os.Getenv("DEBUG_MODE")),
				LocalMode:  stringEnvToBool(os.Getenv("LOCAL_MODE")),
			},
			Log: console.Config{
				Depth:      -1,
				Format:     console.JSONFormat,
				TimeFormat: time.RFC3339Nano,
				Level:      logur.Trace,
				Color:      false,
				Stop:       false,
			},
		}),
	).Run()
}

func stringEnvToBool(s string) bool {
	if s == "true" || s == "1" {
		return true
	}

	return false
}
