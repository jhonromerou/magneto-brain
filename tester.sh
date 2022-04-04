#!/bin/bash

PATH_LOGS="logs"
IMAGE_NAME="magneto-brain-go"
DOCKERFILE_PATH="tester.Dockerfile"
DCOKER_PATH="."
LOG_BUILD="${PATH_LOGS}/docker-tester-build.log"
PATH_PROJECT="/home/ubuntu/project"

mkdir -p ${PATH_LOGS}


echo "**** Docker building..."
echo "View logs in: ${LOG_BUILD}"
docker build -t ${IMAGE_NAME} -f ${DOCKERFILE_PATH} ${DCOKER_PATH} > ${LOG_BUILD}

build_success=$(grep "Successfully" ${LOG_BUILD})
if [ "${build_success}" ];  then
    echo "Successfully"
fi

docker run --rm -it \
    --name "${IMAGE_NAME}" \
    -v ${PWD}:${PATH_PROJECT} \
    ${IMAGE_NAME}
