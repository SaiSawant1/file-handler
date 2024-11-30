run:
	@go run main.go
build:
	@go build -o bin/fileHandler main.go
build-win:
	 @GOOS=windows GOARCH=amd64 go build -o bin/fileHandler.exe main.go
run-prod:
	@ ./bin/fileHandler ./source ./destination

test:
	@go test ./... -v


