# IMAGE_NAME="localhost:5000/backend-app:latest"
IMAGE_NAME="registry.local/backend-app:latest"

test: swagger
	go test -v -cover -short ./...

swagger:
	swag init -g cmd/main.go -o ./api

generate:
	go generate

docker-build:
	docker build -t ${IMAGE_NAME} .

docker-run:
	docker run --rm -p 9090:9090 ${IMAGE_NAME}

docker-push: docker-build
	docker image push ${IMAGE_NAME}


.PHONY: gererate swagger test

