run:
	@go run main.go
build:
	@go build -o bin/fileHandler main.go
run-prod:
	@ ./bin/fileHandler
