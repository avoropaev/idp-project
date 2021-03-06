package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/avoropaev/idp-project/config"
	"github.com/joho/godotenv"
	legoapp "github.com/vseinstrumentiru/lego/v2/app"
	cfg "github.com/vseinstrumentiru/lego/v2/config"
	"github.com/vseinstrumentiru/lego/v2/log/handlers/console"
	"github.com/vseinstrumentiru/lego/v2/metrics/exporters"
	"github.com/vseinstrumentiru/lego/v2/module"
	"github.com/vseinstrumentiru/lego/v2/transport/http"
	"github.com/vseinstrumentiru/lego/v2/transport/postgres"
	"logur.dev/logur"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	legoapp.NewRuntime(
		legoapp.ServerMode(),
		legoapp.Provide(
			module.HTTPServer,
			module.HTTPClient,
			module.MuxRouter,
			module.PostgresPack,
			module.Pipeline,
		),
		legoapp.WithConfig(&Config{
			App: cfg.Application{
				Name:       os.Getenv("APP_NAME"),
				DataCenter: os.Getenv("DATA_CENTER"),
				DebugMode:  stringEnvToBool(os.Getenv("DEBUG_MODE")),
				LocalMode:  stringEnvToBool(os.Getenv("LOCAL_MODE")),
			},
			HTTP: http.Config{
				Port:            atoi(os.Getenv("HTTP_PORT")),
				ShutdownTimeout: 5 * time.Second,
			},
			Postgres: postgres.Config{
				DSN: "postgresql://" + os.Getenv("POSTGRES_USER") + ":" + os.Getenv("POSTGRES_PASS") + "@" +
					os.Getenv("POSTGRES_HOST") + ":" + os.Getenv("POSTGRES_PORT") + "/" + os.Getenv("POSTGRES_NAME"),
			},
			Log: console.Config{
				Depth:      -1,
				Format:     console.JSONFormat,
				TimeFormat: time.RFC3339Nano,
				Level:      logur.Trace,
				Color:      false,
				Stop:       false,
			},
			Jaeger: exporters.Jaeger{
				Addr: os.Getenv("JAEGER_ADDR"),
			},
			External: config.External{
				S1: os.Getenv("S1URL"),
				S2: os.Getenv("S2URL"),
			},
		}),
	).Run(app{})
}

func stringEnvToBool(s string) bool {
	if s == "true" || s == "1" {
		return true
	}

	return false
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return i
}
