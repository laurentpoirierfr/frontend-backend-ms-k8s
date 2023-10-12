IMAGE_NAME="localhost:5000/postgresql:latest"

docker-build:
	docker build -t ${IMAGE_NAME} .

docker-push: docker-build
	docker image push ${IMAGE_NAME}

.PHONY: docker-build docker-push