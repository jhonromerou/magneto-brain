#bin/bash

rm -rf coverage
mkdir coverage

if [ "$1" = "--top" ]; then
    PACKAGES=$(go list ./... | grep -v **/node_modules/)
else
    PACKAGES=$(go list ./...)
fi

go test -covermode=count -coverprofile=./coverage/coverage.out -timeout 60s $PACKAGES

if [ -f "./coverage/coverage.out" ]; then
    go tool cover -html=./coverage/coverage.out
fi