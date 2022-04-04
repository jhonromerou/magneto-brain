package dotenv

import (
	"github.com/jhonromerou/magneto-brain/src/domain"

	"github.com/google/wire"
)

var SetEnviroments = wire.NewSet(
	NewDotEnvEnvironmentReposity,
	wire.Bind(new(domain.EnvironmentRespository), new(*DotEnvEnvironmentReposity)),
)
