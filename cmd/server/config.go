package main

import (
	cfg "github.com/vseinstrumentiru/lego/v2/config"
	"github.com/vseinstrumentiru/lego/v2/log/handlers/console"
	"github.com/vseinstrumentiru/lego/v2/metrics/exporters"
	"github.com/vseinstrumentiru/lego/v2/metrics/tracing"
	"github.com/vseinstrumentiru/lego/v2/transport/http"
	"github.com/vseinstrumentiru/lego/v2/transport/postgres"

	"github.com/avoropaev/idp-project/config"
)

type Config struct {
	App      cfg.Application
	HTTP     http.Config
	Postgres postgres.Config
	Log      console.Config
	Tracing  tracing.Config
	Jaeger   exporters.Jaeger
	External config.External
}
