run:
	@echo "Running application..."
	go run ./cmd/app/main.go

migrate-up:
	@echo "Applying database migrations..."
	go run ./cmd/migrate/main.go up

migrate-down:
	@echo "Reverting database migrations..."
	go run ./cmd/migrate/main.go down