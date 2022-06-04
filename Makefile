.PHONY: docs
.PHONY: env
.PHONY: build

build:
	@go build -o ./build/server ./src/server/cmd/main.go
	@echo 🗃 builded

image:
	@echo 🐋
	@docker build . -t server:latest --quiet

run:
	@./build/server

compose-up:
	@echo 🚀
	@docker-compose up -d

docs:
	@redoc-cli serve docs/openapi.yaml

changelog:
	@git log --pretty=format:' - %s **(%h)** [See commit.](https://github.com/luisnquin/meow-app/commit/%H)<br>' > CHANGELOG.md
	@sed  -i '1i # Meow app - Changelog' CHANGELOG.md

env:
	@if [ ! -f "./venv/bin/activate" ]; then virtualenv venv; fi; source ./venv/bin/activate; pip freeze > requirements.txt

migration:
	@python ./tools/migration/main.py
