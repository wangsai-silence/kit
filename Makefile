CURRENT_DIR=$(pwd)

PROJ = kit

MODULE = "kit"

.PHONY: build, install

build:
	go build -o build/bin/${PROJ} main/kit.go

install:
	go install main/kit.go