IMAGE_NAME="bluehousemax/frontend-app:latest"

DOCKER_HUB_USERNAME="bluehousemax"
DOCKER_HUB_TOKEN="dckr_pat_v4RjqlXOa4aj7MvzyOqNX8kdj0Y"

docker-build:
	docker build -t ${IMAGE_NAME} .

docker-run:
	docker run -e BACKEND_ADDRESS='http:localhost:9090' --rm -p 8080:8080 ${IMAGE_NAME} 

docker-push: docker-build
	docker login -u ${DOCKER_HUB_USERNAME} -p ${DOCKER_HUB_TOKEN}  
	docker image push ${IMAGE_NAME}
	# eval $(minikube docker-env) && minikube image load ${IMAGE_NAME}
	# eval $(minikube docker-env) && minikube image list | grep ${IMAGE_NAME}