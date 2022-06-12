FROM golang:1.18.2-alpine3.16 AS server

ARG project_name=gif-app

WORKDIR /home

# Essential pkgs
RUN apk add\
    bash\
    git\
    python3\
    python3-dev\
    py3-pip\
    gcc\
    g++\
    libpq-dev

RUN echo 'alias python="python3"' >> ~/.bashrc
RUN echo 'alias py="python3"' >> ~/.bashrc

# Simulating path
RUN mkdir -p ./${project_name}/src

# Project preparation
COPY ./src/server ./${project_name}/src/server
COPY ./config-server.json .
COPY ./schemas.sql .
COPY ./go.mod ./${project_name}
COPY ./go.sum ./${project_name}
COPY ./private.rsa.key .
COPY ./public.rsa.key .
COPY ./tools ./tools
COPY ./docs ./docs

COPY ./requirements.txt .

# Project build
RUN (cd ${project_name}; go mod tidy)
RUN (cd ${project_name}; go build -o ../server ./src/server/cmd/main.go)
RUN pip install -r requirements.txt

# Cleaning source files
RUN rm -rf ${project_name}
RUN apk del \
    python3-dev\
    gcc\
    g++\
    libpq-dev

# Public and private keys
# RUN ssh-keygen -t rsa -b 4096 -m PEM -f private.rsa.key
# RUN openssl rsa -in private.rsa.key -pubout -outform PEM -out public.rsa.key

EXPOSE 5800

CMD [ "./server" ]
