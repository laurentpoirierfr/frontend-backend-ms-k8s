IMAGE_NAME="frontend-app:latest"

docker-build:
	docker build -t ${IMAGE_NAME} .

docker-run:
	docker run -e BACKEND_ADDRESS='http:localhost:9090' --rm -p 8080:8080 ${IMAGE_NAME} 

docker-push:
	eval $(minikube docker-env) && minikube image load ${IMAGE_NAME}
	eval $(minikube docker-env) && minikube image list | grep ${IMAGE_NAME}