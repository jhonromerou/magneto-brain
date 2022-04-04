PACKAGES := $(shell go list ./...)

tests:
	@go test -count=1 -cover -timeout 60s $(PACKAGES)
