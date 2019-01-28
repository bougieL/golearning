# image:
	# docker build -f ./dockerfile/golang.dockerfile . -t golang:golearning
	# docker build -f ./dockerfile/mysql.dockerfile . -t mysql:golearning

dev:
	docker-compose -f ./build/docker-compose.dev.yml up

stop:
	docker-compose -f ./build/docker-compose.dev.yml stop
	docker-compose -f ./build/docker-compose.pro.yml stop

vendors:
	go mod vendor
	go mod tidy

build:
	docker build -o ./dist/golearning ./main.go

release:
	docker-compose -f ./build/docker-compose.prod.yml -d

# export data from container
# docker exec golearning_mysql sh -c 'mysqldump -uroot -p$MYSQL_ROOT_PASSWORD $MYSQL_DATABASE > /mysqldump/$MYSQL_DATABASE.sql'
dump:
	# make sure the golearning_mysql is running
	docker exec golearning_mysql sh -c 'mysqldump -uroot -p$$MYSQL_ROOT_PASSWORD $$MYSQL_DATABASE > /mysqldump/$$MYSQL_DATABASE.sql'
