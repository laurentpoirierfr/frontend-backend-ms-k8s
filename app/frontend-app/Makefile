IMAGE_NAME="localhost:5000/frontend-app:latest"

docker-build:
	docker build -t ${IMAGE_NAME} .

docker-run:
	docker run -e BACKEND_ADDRESS='http:localhost:9090' --rm -p 8080:8080 ${IMAGE_NAME} 

docker-push: docker-build
	docker image push ${IMAGE_NAME}
