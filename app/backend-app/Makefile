IMAGE_NAME="backend-app:latest"

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

docker-push:
	eval $(minikube docker-env) && minikube image load ${IMAGE_NAME}
	eval $(minikube docker-env) && minikube cache list

.PHONY: gererate swagger test
