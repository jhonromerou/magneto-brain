#!/usr/bin/env bash

rm -rf vendor
make vendor

if [[ -z "$CI_COVERAGE_MIN" ]]; then
    export CI_COVERAGE_MIN=80
fi

if [[ -z "$PACKAGES_OMIT" ]]; then
    export PACKAGES_OMIT="node_modules"
fi

echo -e "**** Obteniendo packages, se omiten: $PACKAGES_OMIT"
PACKAGES=$(go list ./... | grep -vE "$PACKAGES_OMIT")

echo -e "\n**** Ejecutando pruebas..."
go test -v -count=1 -timeout 600s \
    -coverpkg=./... \
    -coverprofile=${PATH_LOGS}/coverage.out \
    -covermode=atomic $PACKAGES \
    > ${PATH_LOGS}/report.out

if [[ $(grep "FAIL" ${PATH_LOGS}/report.out) != "" ]]; then
    echo -e "\n**** Test Failed"
    EXITCODE=255
else
    echo -e "\n**** Test Passed"
fi

# Omitir los wire_gen del coverage
 grep -vE "wire_gen" ${PATH_LOGS}/coverage.out > ${PATH_LOGS}/coverage-filtered.out

# calcular el coverage
COVERAGE=$(go tool cover \
    -func ${PATH_LOGS}/coverage-filtered.out \
    | tail -1 | grep "total:" \
    | awk '{ print substr($3, 1, length($3)-1) }'
)


if [[ "${COVERAGE%.*}" -lt "$CI_COVERAGE_MIN" ]]; then
    echo -e "\n**** Coverage fallo: $COVERAGE is menor que el minimo esperado: $CI_COVERAGE_MIN"    
    EXITCODE=255
else
    echo -e "\n**** Coverage pasado: $COVERAGE is mayor que el minimo: $CI_COVERAGE_MIN"
fi

rm -rf vendor
