APP?=./cmd/app
BIN?=./bin/github.com/lenvendo/ig-absolut-api
PATH_ROJECT?=github.com/lenvendo/ig-absolut-api

VERSION?=0.1.0
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

.PHONY: proto
proto:
	sh ./scripts/protoc-gen.sh

.PHONY: clean
clean:
	rm -f ${BIN}

.PHONY: build
build: clean
	CGO_ENABLED=0 go build -ldflags "-s -w \
		-X ${PATH_ROJECT}/pkg/health.Version=${VERSION} \
		-X ${PATH_ROJECT}/pkg/health.Commit=${COMMIT} \
		-X ${PATH_ROJECT}/pkg/health.BuildTime=${BUILD_TIME}" \
		-a -installsuffix cgo -o ${BIN} ${APP}

.PHONY: run
run: build
	${BIN}

.PHONY: test
test:
	go test -v -race ./...
