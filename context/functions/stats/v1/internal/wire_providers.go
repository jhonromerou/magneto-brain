package internal

import (
	"github.com/jhonromerou/magneto-brain/src/infrastructure/logruslogger"
	"github.com/jhonromerou/magneto-brain/src/repositories"

	"github.com/google/wire"
)

var SetStatsFunction = wire.NewSet(
	logruslogger.SetLogrusLogger,
	repositories.SetStatsRepository,
	NewHandler,
)
