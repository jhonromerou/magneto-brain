//go:build wireinject

package di

import (
	"github.com/jhonromerou/magneto-brain/context/functions/stats/v1/internal"

	"github.com/google/wire"
)

func Initialize() (*internal.Handler, error) {
	wire.Build(internal.SetStatsFunction)
	return &internal.Handler{}, nil
}
