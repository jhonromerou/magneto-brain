//go:build wireinject

package di

import (
	"github.com/jhonromerou/magneto-brain/context/functions/analysis/v1/internal"

	"github.com/google/wire"
)

func Initialize() (*internal.Handler, error) {
	wire.Build(internal.SetAnalysisFunction)
	return &internal.Handler{}, nil
}
