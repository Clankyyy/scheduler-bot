build: 
	@go build -C cmd/scheduler-bot -o ../../bin/scheduler-bot

run: build
	@./bin/scheduler-bot

test: 
	@go test -v ./...
