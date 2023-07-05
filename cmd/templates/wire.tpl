//go:build wireinject
// +build wireinject

package {{.PkgName}}

import (
	"context"

	"github.com/google/wire"
)

func New(ctx context.Context, cfg Config) (* {{.StructName}}, error) {
	wire.Build(RegisterSet)
	return & {{.StructName}} {}, nil
}
