package log

import (
	"github.com/google/wire"

	"github.com/thoohv5/template-grpc/internal/pkg/config"
	"github.com/thoohv5/template-grpc/pkg/log"
	"github.com/thoohv5/template-grpc/pkg/log/adapter"
)

// ProviderSet is log providers.
var ProviderSet = wire.NewSet(
	New,
)

func New(cf config.IConfig) log.ILogger {
	return adapter.New(log.Zap)
}
