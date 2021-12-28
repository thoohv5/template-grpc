package app

import "github.com/thoohv5/template-grpc/internal/pkg/config"

type IApp interface {
	GetConfig() config.IConfig
	Run(addr ...string) error
}
