//go:build wireinject

package di

import (
	"github.com/jhonromerou/magneto-brain/context/functions/store/v1/internal"

	"github.com/google/wire"
)

func Initialize() (*internal.Handler, error) {
	wire.Build(internal.SetStoreFunction)
	return &internal.Handler{}, nil
}
