# EcoLink Core - Secure Development Makefile

.PHONY: help setup dev build test security clean

# Default target
help: ## Show available commands
	@echo "EcoLink Core - Available Commands:"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

setup: ## Initial project setup
	@echo "🔧 Setting up EcoLink Core..."
	go mod tidy
	go mod download
	@echo "✅ Setup complete"

dev: ## Start development environment
	@echo "🚀 Starting development environment..."
	docker-compose up --build

build: ## Build for production
	@echo "🏗️  Building for production..."
	docker-compose -f docker-compose.prod.yml build

test: ## Run all tests
	@echo "🧪 Running tests..."
	go test ./tests/unit/... -v
	go test ./tests/integration/... -v
	@echo "✅ Tests completed"

test-coverage: ## Run tests with coverage
	@echo "📊 Running tests with coverage..."
	go test ./tests/unit/... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	@echo "✅ Coverage report generated: coverage.html"

security: ## Run security checks
	@echo "🔒 Running security checks..."
	@echo "Checking for vulnerabilities..."
	go list -json -deps ./... | nancy sleuth
	@echo "Running static analysis..."
	staticcheck ./...
	@echo "Checking dependencies..."
	go mod verify
	@echo "✅ Security checks completed"

lint: ## Run code linting
	@echo "🔍 Running linters..."
	gofmt -l .
	go vet ./...
	staticcheck ./...
	@echo "✅ Linting completed"

clean: ## Clean build artifacts and containers
	@echo "🧹 Cleaning up..."
	docker-compose down -v
	docker system prune -f
	go clean -cache
	rm -f coverage.out coverage.html
	@echo "✅ Cleanup completed"

migrate: ## Run database migrations (when implemented)
	@echo "🗄️  Running database migrations..."
	# TODO: Implement database migrations
	@echo "✅ Migrations completed"

audit: ## Run comprehensive security audit
	@echo "🔍 Running security audit..."
	@echo "Dependency vulnerabilities:"
	go list -json -deps ./... | nancy sleuth
	@echo "Static analysis:"
	staticcheck ./...
	@echo "Go security checker:"
	gosec ./...
	@echo "✅ Security audit completed"

benchmark: ## Run performance benchmarks
	@echo "⚡ Running benchmarks..."
	go test -bench=. -benchmem ./tests/unit/...
	@echo "✅ Benchmarks completed"

docker-security: ## Scan Docker images for vulnerabilities
	@echo "🐳 Scanning Docker images..."
	docker run --rm -v /var/run/docker.sock:/var/run/docker.sock \
		-v $(PWD):/src aquasec/trivy image ecolink-backend:latest
	@echo "✅ Docker security scan completed"