package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/impact-hack-sos-barnbyar/api/config"
	"github.com/impact-hack-sos-barnbyar/api/internal/healthcheck"
	"github.com/impact-hack-sos-barnbyar/api/internal/platform/api/host"
	"github.com/impact-hack-sos-barnbyar/api/internal/platform/api/route"
	"github.com/impact-hack-sos-barnbyar/api/internal/platform/lifecycle"
)

// @title		SOS API
// @version	1.0
func main() {
	ctx := context.Background()
	slog.Info("starting the service. Fetching application config")

	cfg := config.Get()
	slog.Info("configs were fetched. Instantiating router")

	router := route.NewRouter()
	router.
		WithAPIDocumentation().
		WithRoute(route.NewRoute("/health", route.GET, healthcheck.GetStatus))

	slog.With("port", cfg.Port).Info("running web host")
	host := host.New(cfg.Port, router.GetRouter())
	host.Run()

	lifecycle.ListenForApplicationShutDown(func() {
		slog.Info("terminating the web host")
		if err := host.Terminate(ctx); err != nil {
			slog.With("err", err.Error()).Error("error during host termination")
		}

	}, make(chan os.Signal))
}
