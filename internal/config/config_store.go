package config

import (
	"sync"
)

const (
	DefaultGrpcPort uint32 = 8090
)

type ConfigStore struct {
	GrpcPort uint32
}

var (
	store *ConfigStore
	once  sync.Once
)

func Get() *ConfigStore {
	return store
}

func Initialize() {
	once.Do(func() {
		store = &ConfigStore{
			GrpcPort: DefaultGrpcPort,
		}
	})
}
