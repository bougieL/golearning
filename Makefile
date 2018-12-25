image:
	docker build -f ./dockerfile/golang.dockerfile . -t golang:golearning
	# docker build -f ./dockerfile/mysql.dockerfile . -t mysql:golearning

dev:
	docker-compose -f ./build/docker-compose.dev.yml up

stop:
	docker-compose -f ./build/docker-compose.dev.yml stop
	docker-compose -f ./build/docker-compose.prod.yml stop

vendors:
	go mod vendor
	go mod tidy

build:
	docker build -o ./dist/golearning ./main.go

release:
	docker-compose -f ./build/docker-compose.prod.yml -d

# 导出数据
dump:
