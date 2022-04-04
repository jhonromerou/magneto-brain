PACKAGES := $(shell go list ./... | grep -v **/node_modules/ )

.PHONY: vendor
vendor:
	@go mod vendor
	@go mod tidy

.PHONY: stage
stage:
	@./scripts/deploy-lambdas context stage

.PHONY: mocks
mocks:
	@./scripts/create-mocks

.PHONY: coverage-review
coverage-review:
	@./scripts/go-coverage.sh "--top"

# Ejecuta la validacion de codigo en un docker
.PHONY: tester
tester:
	@./tester.sh

# Desplegar proyecto con docker
.PHONY: deployer
deployer:
	@./deployer.sh
	

# Limpiar lambdas
# Si el despliegue presente fallo del estilo 'could not be found' se recomienda limpiar
BIN_MODULES="../../../node_modules/.bin/serverless"
.PHONY: lambda-purge
lambda-purge:
	@cd context/functions/analysis && $(BIN_MODULES) --stage stage remove
	@cd context/functions/stats && $(BIN_MODULES) --stage stage remove
	@cd context/functions/store && ../../../node_modules/.bin/serverless --stage stage remove