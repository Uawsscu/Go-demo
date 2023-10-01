###################
# Docker commands
up:
	docker-compose up -d

down:
	docker-compose down --remove-orphans

clean:
	sudo rm -rf db/data

###################
# Run Main and Swagger
.PHONY: run

run:
	export PATH=$(shell go env GOPATH)/bin:$$PATH; \
	swag init -g main.go -o docs; \
	go run main.go