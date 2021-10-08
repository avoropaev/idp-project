package main

import (
	"github.com/avoropaev/idp-project/config"
	legoapp "github.com/vseinstrumentiru/lego/v2/app"
	"github.com/vseinstrumentiru/lego/v2/server"
	"github.com/vseinstrumentiru/lego/v2/transport/http"
)

func main() {
	server.Run(
		app{HTTP: http.NewDefaultConfig()},
		legoapp.WithConfig(&Config{
			External: config.External{S1: "http://localhost:8080", S2: "http://localhost:8080"},
		}),
	)
}
