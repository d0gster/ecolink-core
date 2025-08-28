.PHONY: help dev build clean test

help: ## Mostra esta ajuda
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

dev: ## Inicia ambiente de desenvolvimento
	@echo "🌱 Iniciando EcoLink em modo desenvolvimento..."
	docker-compose up --build

dev-backend: ## Inicia apenas o backend
	@echo "🚀 Iniciando backend Go..."
	cd backend && go run cmd/main.go

dev-frontend: ## Inicia apenas o frontend
	@echo "⚡ Iniciando frontend SvelteKit..."
	cd frontend && npm run dev

build: ## Build para produção
	@echo "📦 Fazendo build para produção..."
	docker-compose build

clean: ## Limpa containers e imagens
	@echo "🧹 Limpando ambiente Docker..."
	docker-compose down --rmi all --volumes --remove-orphans

test: ## Executa testes
	@echo "🧪 Executando testes..."
	cd backend && go test ./...

install: ## Instala dependências
	@echo "📥 Instalando dependências..."
	cd backend && go mod tidy
	cd frontend && npm install

logs: ## Mostra logs dos containers
	docker-compose logs -f