PACKAGE=github.com/magaldima/bizday
CURRENT_DIR=$(shell pwd)
DIST_DIR=${CURRENT_DIR}/dist

.PHONY: build test clean

build:
	go build -v -o ${DIST_DIR}/server .

test:
	go test $(shell go list ./... | grep -v /vendor/) -race -short -v

clean:
	-rm -rf ${CURRENT_DIR}/dist
