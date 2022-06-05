.PHONY: docs
.PHONY: env
.PHONY: build

build:
	@go build -o ./build/server ./src/server/cmd/main.go
	@echo 🗃 builded

run:
	@./build/server

image:
	@echo 🐋
	@docker build . -t server:latest --quiet

store:
	@docker-compose up -d postgres_db redis_cache

docs:
	@redoc-cli serve docs/openapi.yaml

changelog:
	@git log --pretty=format:' - %s **(%h)** [See commit.](https://github.com/luisnquin/gif-app/commit/%H)<br>' > CHANGELOG.md
	@sed  -i '1i # Meow app - Changelog' CHANGELOG.md

env:
	@if [ ! -f "./venv/bin/activate" ]; then virtualenv venv; fi; source ./venv/bin/activate; pip freeze > requirements.txt

migration:
	@python ./tools/migration/main.py
