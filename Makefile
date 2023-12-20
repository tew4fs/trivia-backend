default: build

build:
	@ go build -o bin/golang-api-skeleton main.go

run: build
	@ ./bin/golang-api-skeleton