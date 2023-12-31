.DEFAULT_GOAL := build

brew:
	brew install golangci-lint
	brew install staticcheck


clean:
	rm -rf dist/
	rm -rf tmp/

mod:
	go mod download
	go mod tidy

check:
	staticcheck  ./...

lint:
	go vet ./...
	golangci-lint run

lint-fix:
	go vet ./...
	golangci-lint run --fix

build-release:
	mkdir -p dist
	go build -o dist/opsinsights-exporter ./pkg/cmd/opsinsights-exporter

build-docker:
	docker build -t opsinsights-exporter .

run-release:
	go run ./pkg/cmd/opsinsights-exporter

run-docker:
	docker rm opsinsights-exporter
	docker run --name opsinsights-exporter opsinsights-exporter

build:  clean mod lint-fix build-release

run:  clean mod lint-fix run-release

docker:  clean mod lint-fix run-docker

# disallow any parallelism (-j) for Make. This is necessary since some
# commands during the build process create temporary files that collide
# under parallel conditions.
.NOTPARALLEL:

.PHONY: clean mod lint lint-fix release alll