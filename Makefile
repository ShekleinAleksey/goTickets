swag:
	swag init -g ./internal/app/app.go
migrate:
	migrate -path ./schema -database 'postgres://postgres:postgres@0.0.0.0:5432/postgres?sslmode=disable' up
build:
	docker-compose build
run:
	docker-compose up
create-m:
	migrate create -ext sql -dir ./schema -seq init