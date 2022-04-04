#bin/bash

PACKAGES=$(go list ./...)

go test -count=1 -cover -timeout 60s ${PACKAGES}