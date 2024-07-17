build:
	go build main.go

run:
	go run main.go


migrate:
	migrate -path ./schema -database 'postgresql://postgres:admin@localhost:5432/postgres' up
