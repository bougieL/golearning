# /usr/bin/bash
go mod vendor
go mod tidy
gin -p 8080 run ./main.go