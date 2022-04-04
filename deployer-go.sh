#!/usr/bin/env bash

# Configura los secretos para aws
source config/aws_credentials.sh

# Dependencias
rm -rf node_modules
npm install

# Definir modulos
rm -rf vendor
make vendor

# Desplegar
make stage

echo -e "\n Endpoint para POST analysis"
grep "POST -" logs/aws-deployer-context.log

echo -e "\n Endpoint para GET stats"
grep "GET -" logs/aws-deployer-context.log