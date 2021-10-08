package main

import (
	legoapp "github.com/vseinstrumentiru/lego/v2/app"
	"github.com/vseinstrumentiru/lego/v2/server"
	"github.com/vseinstrumentiru/lego/v2/transport/http"
)

func main() {
	server.Run(
		app{HTTP: http.NewDefaultConfig()},
		legoapp.WithConfig(&config{}),
	)
}
