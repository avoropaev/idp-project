package main

import (
	cfg "github.com/vseinstrumentiru/lego/v2/config"
	"github.com/vseinstrumentiru/lego/v2/metrics/tracing"
	"github.com/vseinstrumentiru/lego/v2/multilog"
	"github.com/vseinstrumentiru/lego/v2/multilog/log"
	"github.com/vseinstrumentiru/lego/v2/transport/http"
	"github.com/vseinstrumentiru/lego/v2/transport/postgres"

	"github.com/avoropaev/idp-project/config"
)

type Config struct {
	App      *cfg.Application
	HTTP     *http.Config
	Postgres *postgres.Config
	Log      *log.Config
	Logger   *multilog.Config
	Tracing  *tracing.Config
	External config.External
}
