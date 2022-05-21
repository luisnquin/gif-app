.PHONY: build

build:
	@echo 🗃
	@go build -o ./build/server ./src/server/cmd/main.go

image:
	@echo 🐋
	@docker build . -t server:latest --quiet

run:
	@./build/server

compose-up:
	@echo 🚀
	@docker-compose up -d