
APP_NAME=user-service

test: swagger
	go test -v -cover -short ./...

swagger:
	swag init -g cmd/main.go -o ./api

generate:
	go generate

docker-build:
	docker build -t ${APP_NAME} .

docker-run:
	docker run --rm -p 9090:9090 ${APP_NAME}

.PHONY: gererate swagger test
