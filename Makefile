BINARY_NAME=SwiftGOApp.exe

build:
	@echo "Building SwiftGO..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "SwiftGO built!"

run: build
	@echo "Starting SwiftGO..."
	@./tmp/${BINARY_NAME} &
	@echo "SwiftGO started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm tmp/${BINARY_NAME}
	@echo "Cleaned!"

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!"

start: run

stop:
	@echo "Stopping SwiftGO..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped SwiftGO!"

restart: stop start
