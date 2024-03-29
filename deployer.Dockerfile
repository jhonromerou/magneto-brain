FROM golang:1.17

ENV PATH_PROJECT "/home/ubuntu/project"

RUN apt-get update && \
    apt install npm -y

RUN apt-get install unzip -y
RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip" \
    && unzip awscliv2.zip \
    && ./aws/install

ENTRYPOINT cd ${PATH_PROJECT} && ./deployer-go.sh
