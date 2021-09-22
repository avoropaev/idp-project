package main

import multilog "github.com/vseinstrumentiru/lego/v2/multilog"

type application struct {
	// Log will injected automatically
	Log multilog.Logger
}

func (app application) Providers() []interface{} {
	return []interface{}{
		// add your constructors here...
	}
}
func (app application) ConfigureService() error {
	// here you can build your service...
	return nil
}
