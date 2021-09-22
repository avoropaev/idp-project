package main

import (
	app "github.com/vseinstrumentiru/lego/v2/app"
	"github.com/vseinstrumentiru/lego/v2/server"
)

func main() {
	server.Run(application{}, app.WithConfig(&config{}))
}
