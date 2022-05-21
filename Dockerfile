FROM golang:1.18.2-alpine3.15 AS server

ARG project_name=meow-app

WORKDIR /home

# Essential pkgs
RUN apk add openssl
RUN apk add bash
RUN apk add curl
RUN apk add git

# Simulating path
RUN mkdir -p ./${project_name}/src

# Project preparation
COPY ./src/server ./${project_name}/src/server
COPY ./server.json .
COPY ./go.mod ./${project_name}
COPY ./go.sum ./${project_name}

# Project build
RUN (cd ${project_name}; go mod tidy)
RUN (cd ${project_name}; go build -o ../server ./src/server/cmd/main.go)

# Cleaning source files
RUN rm -rf ${project_name}

# Public and private keys
RUN openssl genrsa -out private.rsa 1024
RUN openssl rsa -in private.rsa -pubout > public.rsa.pub

EXPOSE 5800

CMD [ "./server" ]