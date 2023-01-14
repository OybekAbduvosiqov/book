run:
	go run cmd/main.go
swag:
	swag init -g api/api.go -o api/docs
migration-up:
	migrate -path ./migrations/postgres/ -database 'postgres://oybek:oybek@localhost:5432/catalog?sslmode=disable' up
