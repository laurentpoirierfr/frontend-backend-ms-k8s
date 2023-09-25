IMAGE_NAME="bluehousemax/backend-app:latest"

DOCKER_HUB_USERNAME="bluehousemax"
DOCKER_HUB_TOKEN="dckr_pat_v4RjqlXOa4aj7MvzyOqNX8kdj0Y"

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
	docker login -u ${DOCKER_HUB_USERNAME} -p ${DOCKER_HUB_TOKEN}  
	docker image push ${IMAGE_NAME}
	# eval $(minikube docker-env) && minikube image load ${IMAGE_NAME}
	# eval $(minikube docker-env) && minikube image list | grep ${IMAGE_NAME}


.PHONY: gererate swagger test


