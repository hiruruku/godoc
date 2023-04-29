# Description: Makefile for backend
# HostとDockerの両方で動く前提だが、基本Hostで使う

APP_NAME=myapp
GOOS=linux
GOARCH=amd64

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main main.go

.PHONY: test
test:
	go test -v ./...

.PHONY: clean
clean:
	rm -f ${APP_NAME}

.PHONY: run
run:
	go run main.go

.PHONY: deps
deps:
	go get -v -d ./...

.PHONY: build-linux
build-linux:
	GOOS=${GOOS} GOARCH=${GOARCH} go build -o ${APP_NAME}-${GOOS}-${GOARCH}

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: all
all: clean deps build
