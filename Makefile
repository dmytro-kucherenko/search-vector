build:
	@echo "Building..."
	@go build -o bin/main main.go

run:
	@go run main.go

watch:
	@air -c air.toml
