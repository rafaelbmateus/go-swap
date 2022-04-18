compose = docker-compose -f build/docker-compose.yml

.PHONY: build
build: ##@development Build docker image.
	$(compose) build

.PHONY: up
up: build ##@development Build and start development environment in background.
	$(compose) up -d

.PHONY: logs
logs: ##@development Follows development logs [service="svc1 svc2..."].
	$(compose) logs -f --tail=100 $(service)

.PHONY: shell
shell: ##@development Start a shell session within the container.
	$(compose) run --rm app /bin/sh

lint_version ?= v1.40-alpine
.PHONY: lint
lint: ##@development Runs static analysis code.
	docker run --rm \
		-v $(shell pwd):/app \
		-w /app \
		golangci/golangci-lint:$(lint_version) \
		golangci-lint run --timeout 3m

.PHONY: test
test: build ##@development Runs the tests.
	$(compose) run --rm app go test ./...

.PHONY: stop
stop: ##@development Stop development environment.
	$(compose) down -v --remove-orphans
