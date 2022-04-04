HOST := "http://localhost:9000"
PROJECT_KEY := "magneto-brain"
SONAR_PATH=${HOME}/tmp/sonar
SHELL := /bin/bash
PATH := ${SONAR_PATH}/bin:$(PATH)

init:
	@docker-compose up -d
	@echo "Open http://127.0.0.1:9000"

scanner:
	@sonar-scanner

install:
	@mkdir -p ${SONAR_PATH} \
		&& cd ${SONAR_PATH} \
		&& curl -fsSL https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-4.6.2.2472-linux.zip -o sonar.zip \
		&& unzip sonar.zip \
		&& mv sonar-scanner-4.6.2.2472-linux/* ./  \
		&& rm -rf sonar-scanner-4.6.2.2472-linux \
		&& rm -rf sonar.zip