dev:
	docker-compose -f ./build/docker-compose.dev.yml up

vendor:
	go mod vendor
	go mod tidy

build:
	docker build -o ./dist/golearning ./main.go

release:
	docker-compose -f ./build/docker-compose.prod.yml