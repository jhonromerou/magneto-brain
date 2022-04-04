package internal

import (
	"github.com/jhonromerou/magneto-brain/src/infrastructure/logruslogger"
	"github.com/jhonromerou/magneto-brain/src/repositories"

	"github.com/google/wire"
)

// Set groups dependencies for the creation of aws services
var SetStoreFunction = wire.NewSet(
	logruslogger.SetLogrusLogger,
	repositories.SetAnalysisAndStatsRepository,
	NewHandler,
)
