docker-build:
	docker build -t frontend-app .

docker-run:
	docker run -e BACKEND_ADDRESS='http:localhost:9090' --rm -p 8080:8080 frontend-app 