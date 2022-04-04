.PHONY: tests clean build stage

BIN_NODE_MODULES := "../../../node_modules/.bin/serverless"

tests: 
	@./../../../scripts/go-test.sh

clean:
	@rm -rf ./bin

build:
	@./../../../scripts/go-build.sh

stage: build
	$(BIN_NODE_MODULES) --stage stage deploy
