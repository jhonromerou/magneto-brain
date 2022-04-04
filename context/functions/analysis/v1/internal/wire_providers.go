package internal

import (
	"github.com/jhonromerou/magneto-brain/src/infrastructure/logruslogger"
	"github.com/jhonromerou/magneto-brain/src/repositories"

	"github.com/google/wire"
)

var SetAnalysisFunction = wire.NewSet(
	logruslogger.SetLogrusLogger,
	repositories.SetQueueAnalysisRepository,
	NewHandler,
)
