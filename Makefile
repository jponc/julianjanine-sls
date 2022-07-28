.PHONY: build clean deploy

CMD_DIR = ./cmd

build:
	@for f in $(shell ls ${CMD_DIR}); do echo Building $${f} && env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/$${f} cmd/$${f}/*.go; done

clean:
	rm -rf ./bin

deploy-prod: clean build
	sls deploy --stage=prod --verbose
