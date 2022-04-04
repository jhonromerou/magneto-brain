package domain

// Relative path with call in src/repositories
const DEFAULT_ENV_FILE_LOCATION = "../.env"

type EnvironmentRespository interface {
	Get(key string) string
}
