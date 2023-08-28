.PHONY: build

export GO111MODULE=on
export GOPROXY=https://arti.freewheel.tv/api/go/go
BINARY_NAME=creative_e2e_test


default: build

build:
	rm -rf ./bin && mkdir -p ./bin
	go build -ldflags "-X main.revision=$(REV) -X main.goVersion=$(GO_VER)" -o ./bin/${BINARY_NAME}

build-send:
  GOOS=linux GOARCH=amd64 go build -o sendtopic  *.go
gen-proto:
	proto/gen_proto.sh
