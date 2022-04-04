FROM golang:1.17

ENV PACKAGES_OMIT "node_modules"
ENV CI_COVERAGE_MIN 80
ENV PATH_LOGS "logs"
ENV PATH_PROJECT "/home/ubuntu/project"

ENTRYPOINT cd ${PATH_PROJECT} && ./tester-go.sh
