DATABASE_NAME="my-app"
DATABASE_URL="postgresql://root:secret@localhost:5432/${DATABASE_NAME}?sslmode=disable"

createdb:
    docker exec -it postgres12 createdb --username=root --owner=root ${DATABASE_NAME}

dropdb:
    docker exec -it postgres12 dropdb ${DATABASE_NAME}

migrate-init:
	migrate create -ext sql -dir db/migrations -seq init_schema

migrate-up:
	migrate -path db/migrations -database ${DATABASE_URL} -verbose up


migrate-down:
	migrate -path db/migrations -database ${DATABASE_URL} -verbose down	

.PHONY: createdb dropdb migrateup migratedown	