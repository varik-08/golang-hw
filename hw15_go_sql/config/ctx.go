package config

import (
	"context"
	"sync"
)

var (
	ctx     context.Context
	onceCTX sync.Once
)

func InitCTX() {
	onceCTX.Do(func() {
		ctx = context.Background()
	})
}

func GetCTX() context.Context {
	if ctx == nil {
		InitCTX()
	}

	return ctx
}
