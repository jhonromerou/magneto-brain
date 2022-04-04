package dotenv

import (
	"log"

	"github.com/joho/godotenv"
)

var LOCAL_ENVS = map[string]string{
	"DB_TABLE_ANALYSIS":   "analysis",
	"DB_TABLE_STATS":      "stats",
	"QUEUE_NAME_ANALYSIS": "magneto_brain_validator",
}

type DotEnvEnvironmentReposity struct {
}

func (e *DotEnvEnvironmentReposity) From(envPath string) error {
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Error loading enviroment file %s", envPath)
	}
	return nil
}

func (e *DotEnvEnvironmentReposity) Get(key string) string {
	// TODO: usar cuando se lea el archivo
	//return os.Getenv(key)
	return LOCAL_ENVS[key]
}

func NewDotEnvEnvironmentReposity() *DotEnvEnvironmentReposity {
	dotenv := &DotEnvEnvironmentReposity{}
	// TODO: se podria mejorar haciendo implementacion con variables externa para usar .env.testing por ejemplo
	// TODO: leer variables desde archivo y no quemadas.
	//dotenv.From(domain.DEFAULT_ENV_FILE_LOCATION)
	return dotenv
}
