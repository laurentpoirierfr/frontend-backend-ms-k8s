
test: swagger
	go test -v -cover -short ./...

swagger:
	swag init -g cmd/main.go -o ./api

generate:
	go generate

docker-build:
	docker build -t backend-app .

docker-run:
	docker run --rm -p 9090:9090 backend-app 

.PHONY: gererate swagger test
