.PHONY: help setup dev build clean test install

help: ## Mostra esta ajuda
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

setup: ## Configuração inicial do projeto
	@echo "🔧 Configurando projeto..."
	@cp -n .env.example .env || true
	@cp -n frontend/.env.example frontend/.env.local || true
	@echo "✅ Arquivos .env criados. Configure suas credenciais!"
	@echo "📚 Leia AUTH0_SETUP.md para configuração completa"

dev: ## Inicia ambiente de desenvolvimento completo
	@echo "🌱 Iniciando EcoLink em modo desenvolvimento..."
	docker-compose up --build

dev-backend: ## Inicia apenas o backend
	@echo "🚀 Iniciando backend Go..."
	cd backend && go mod tidy && go run cmd/main.go

dev-frontend: ## Inicia apenas o frontend
	@echo "⚡ Iniciando frontend SvelteKit..."
	cd frontend && npm install && npm run dev

build: ## Build para produção
	@echo "📦 Fazendo build para produção..."
	docker-compose build

clean: ## Limpa containers e imagens
	@echo "🧹 Limpando ambiente Docker..."
	docker-compose down --rmi all --volumes --remove-orphans

test: ## Executa testes
	@echo "🧪 Executando testes..."
	cd backend && go test ./...

staticcheck: ## Executa análise estática do código Go
	@echo "🔍 Executando staticcheck..."
	cd backend && go run honnef.co/go/tools/cmd/staticcheck@latest ./...

install: ## Instala dependências
	@echo "📥 Instalando dependências..."
	cd backend && go mod tidy
	cd frontend && npm install

quality: ## Executa verificações de qualidade
	@echo "✨ Verificando qualidade do código..."
	cd backend && go fmt ./...
	cd backend && go vet ./...
	cd backend && go run honnef.co/go/tools/cmd/staticcheck@latest ./...

logs: ## Mostra logs dos containers
	docker-compose logs -f

lines: ## Conta linhas de código do projeto
	@echo "📊 Linhas de código no projeto:"
	@echo "Go: $$(find . -name '*.go' | grep -v node_modules | xargs wc -l 2>/dev/null | tail -1 | awk '{print $$1}') linhas"
	@echo "TypeScript: $$(find . -name '*.ts' | grep -v node_modules | grep -v .svelte-kit | xargs wc -l 2>/dev/null | tail -1 | awk '{print $$1}') linhas"
	@echo "Svelte: $$(find . -name '*.svelte' | grep -v node_modules | grep -v .svelte-kit | xargs wc -l 2>/dev/null | tail -1 | awk '{print $$1}') linhas"
	@echo "Total: $$(find . -name '*.go' -o -name '*.ts' -o -name '*.svelte' | grep -v node_modules | grep -v .svelte-kit | xargs wc -l 2>/dev/null | tail -1 | awk '{print $$1}') linhas"