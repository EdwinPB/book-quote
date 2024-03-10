.PHONY: tailwind-watch
tailwind-watch:
	./bin/tailwindcss -i ./internal/assets/application.css -o ./public/application.css --watch

.PHONY: templ-generate
templ-generate:
	templ generate

.PHONY: templ-watch
templ-watch:
	templ generate --watch --proxy="http://localhost:3000"

.PHONY: dev
dev:
	@echo "Starting development server..."
	@go run cmd/dev/main.go