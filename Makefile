install_postgres:
	docker pull postgres

create_db:
	docker run --name=tages-files-db -e POSTGRES_PASSWORD='1234' -p 5433:5432 -d --rm postgres

schema_up:
	migrate -path ./schema -database 'postgres://postgres:1234@localhost:5433/postgres?sslmode=disable' up

schema_down:
	migrate -path ./schema -database 'postgres://postgres:1234@localhost:5433/postgres?sslmode=disable' down

swag:
	swag init -g cmd/main.go

test:
	go test -v ./...

grpcgen:
	protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        internal/api/api.proto

run:
	make grpcgen && make install_postgres && make create_db && make schema_up && go run .

