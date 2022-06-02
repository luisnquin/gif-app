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

docs:
	@redoc-cli serve api/docs/openapi.yaml

changelog:
	@git log --pretty=format:' - %s [See commit.](https://github.com/luisnquin/meow-app/commit/%H)<br>' > CHANGELOG.md
	@sed  -i '1i # Meow app - Changelog' CHANGELOG.md
