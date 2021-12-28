//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"

	"github.com/thoohv5/template-grpc/internal/data"
	"github.com/thoohv5/template-grpc/internal/pkg/app"
	"github.com/thoohv5/template-grpc/internal/pkg/config"
	"github.com/thoohv5/template-grpc/internal/pkg/log"
	"github.com/thoohv5/template-grpc/internal/router"
	"github.com/thoohv5/template-grpc/internal/server"
	"github.com/thoohv5/template-grpc/internal/service"
	pap "github.com/thoohv5/template-grpc/pkg/app"
)

// initApp init application.
func initApp() (pap.IApp, func(), error) {
	panic(wire.Build(config.ProviderSet, log.ProviderSet, data.ProviderSet, service.ProviderSet, router.ProviderSet, server.ProviderSet, app.ProviderSet))
}
