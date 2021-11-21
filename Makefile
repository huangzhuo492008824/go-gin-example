include .env

default: run

gen:
	swag init

run-cron:
	go run cron.go

run-gin: 
	go run main.go

run: gen run-gin

build: 
	docker build -t gin-blog-docker-scratch .

scratch-compile: 
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-gin-example .

.PHONY: clean
clean:
	rm go-gin-example
	rm main