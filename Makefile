.PHONY: build clean deploy test test-ci lint format

build:
	GOOS=linux GOARCH=amd64 go build -o bin/converttowebp ./cmd/lambda/converttowebp/main.go

clean:
	rm -rf ./bin

deploy: clean build
	npm run deploy

remove:
	npm run remove

test:
	go clean -testcache
	go test -p 1 -v $$(go list ./... | grep -v /node_modules/)

test-ci:
	go clean -testcache
	go test -p 1 -v -coverprofile coverage.out -covermode atomic $$(go list ./... | grep -v /node_modules/)

lint:
	go vet ./...
	golangci-lint run ./...

format:
	gofmt -l -s -w .
	goimports -w -l ./
