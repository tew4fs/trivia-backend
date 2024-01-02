default: build

Local-Image-Name = trivia-backend

build:
	@ go build -o $(Local-Image-Name) main.go

run: build
	@ ./bin/trivia-backend

docker-build: 
	docker build . -t $(Local-Image-Name)

docker-compose:
	docker compose up

docker-run: docker-build docker-compose