IMAGE_NAME="bluehousemax/keycloack:latest"

DOCKER_HUB_USERNAME="bluehousemax"
DOCKER_HUB_TOKEN="dckr_pat_v4RjqlXOa4aj7MvzyOqNX8kdj0Y"

docker-build:
	docker build -t ${IMAGE_NAME} .

docker-push: docker-build
	docker login -u ${DOCKER_HUB_USERNAME} -p ${DOCKER_HUB_TOKEN}  
	docker image push ${IMAGE_NAME}

.PHONY: docker-build