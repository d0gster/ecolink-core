# 🌱 EcoLink - Core

> **"Encurtamento de links sustentável e de alta performance"**

EcoLink é um sistema de encurtamento de links de alta performance, construído seguindo princípios de **Clean Architecture**, **SOLID** e **Clean Code**, com foco em eficiência, escalabilidade e sustentabilidade digital.

## 🏗️ Arquitetura

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   SvelteKit     │    │   Go Backend    │    │   Database      │
│   Frontend      │◄──►│   (Clean Arch)  │◄──►│   Interface     │
│                 │    │                 │    │                 │
└─────────────────┘    └─────────────────┘    └─────────────────┘
         │                       │                       │
         │              ┌─────────────────┐              │
         └─────────────►│  Google OAuth   │◄─────────────┘
                        │  Authentication │
                        └─────────────────┘
```

### Princípios Arquiteturais Implementados

- **🔷 Hexagonal Architecture**: Portas e adaptadores para isolamento de dependências
- **🔶 SOLID Principles**: Single Responsibility, Open/Closed, Liskov Substitution, Interface Segregation, Dependency Inversion
- **🔸 Clean Code**: Código legível, testável e manutenível seguindo Robert C. Martin
- **🔹 Domain-Driven Design**: Modelagem focada no domínio do negócio

## 🚀 Tech Stack

### Backend (Go)
- **Framework**: Gin (alta performance HTTP router)
- **Arquitetura**: Clean Architecture com interfaces
- **Database**: Interface abstrata (Memory/Firestore)
- **Auth**: JWT + Google OAuth 2.0
- **QR Codes**: Geração nativa com go-qrcode
- **Containerização**: Docker multi-stage builds

### Frontend (SvelteKit)
- **Framework**: SvelteKit com SSR/SSG
- **Styling**: TailwindCSS com design system eco-friendly
- **State Management**: Svelte Stores reativos
- **TypeScript**: Type safety completo
- **Auth**: OAuth 2.0 manual implementation
- **Build**: Vite com otimizações de performance

### DevOps & Infraestrutura
- **Containerização**: Docker + Docker Compose
- **Automação**: Makefile com comandos padronizados
- **Logs**: Estruturados para observabilidade
- **Environment**: Configuração centralizada com .env
- **CI/CD Ready**: Estrutura preparada para pipelines

## 🛠️ Instalação e Configuração

### Pré-requisitos
- **Go**: 1.21+ 
- **Node.js**: 18+
- **Docker**: 20+ (opcional)
- **Google Cloud**: Conta para OAuth (obrigatório)

### 1. Configuração Rápida (Demo Mode)

```bash
# Clone e navegue
git clone <repo-url>
cd ecolink-core

# Backend (modo memória)
cd backend
go mod tidy
go run cmd/main.go

# Frontend (novo terminal)
cd ../frontend
npm install
npm run dev
```

**Acesse**: http://localhost:5173

### 2. Configuração Completa (Google OAuth)

#### 2.1 Configure Google OAuth

1. **Google Cloud Console**: https://console.cloud.google.com/
2. **APIs & Services** → **Credentials**
3. **Create Credentials** → **OAuth 2.0 Client IDs**
4. **Web application** com:
   - **Authorized JavaScript origins**: `http://localhost:5173`
   - **Authorized redirect URIs**: `http://localhost:5173/auth/callback`

#### 2.2 Configure Variáveis de Ambiente

```bash
# Frontend
cd frontend
cp .env.example .env.local

# Edite .env.local:
VITE_GOOGLE_CLIENT_ID=seu_client_id_aqui
VITE_GOOGLE_CLIENT_SECRET=seu_client_secret_aqui
AUTH_SECRET=chave_secreta_aleatoria
```

#### 2.3 Execute o Ambiente

```bash
# Backend
cd backend
go run cmd/main.go

# Frontend (novo terminal)
cd frontend
npm run dev
```

### 3. Docker (Ambiente Completo)

```bash
# Build e execute
docker-compose up --build

# Ou usando Makefile
make dev
```

**Serviços**:
- **Frontend**: http://localhost:5173
- **Backend**: http://localhost:8080
- **Health Check**: http://localhost:8080/health

## ✨ Features Implementadas

### 🔐 Autenticação e Segurança
- ✅ **Google OAuth 2.0**: Implementação manual seguindo RFC 6749
- ✅ **JWT Middleware**: Proteção de rotas com tokens seguros
- ✅ **Session Management**: Persistência segura no localStorage
- ✅ **CORS**: Configuração adequada para desenvolvimento/produção

### 🔗 Core Functionality
- ✅ **URL Shortening**: Algoritmo hash único para códigos curtos
- ✅ **QR Code Generation**: Geração nativa de alta qualidade
- ✅ **Click Tracking**: Métricas básicas de redirecionamentos
- ✅ **User Links**: Dashboard personalizado por usuário

### 🎨 Interface e UX
- ✅ **Design Responsivo**: Mobile-first com TailwindCSS
- ✅ **Componentes Reativos**: Svelte stores para estado global
- ✅ **Background Video**: Elemento visual eco-friendly
- ✅ **Error Handling**: Feedback visual para todas as ações

### 🏗️ Arquitetura e Qualidade
- ✅ **Clean Architecture**: Separação clara de responsabilidades
- ✅ **Interface Abstraction**: Database interface para múltiplos adapters
- ✅ **Error Handling**: Tratamento robusto em todas as camadas
- ✅ **Code Quality**: Seguindo princípios de Clean Code

## 🧪 Testes e Qualidade

### Executar Testes
```bash
# Backend
cd backend
go test ./...

# Frontend
cd frontend
npm test
```

### Code Quality
```bash
# Linting
make lint

# Format
make format

# Security scan
make security
```

## 📊 API Endpoints

### Links
- `POST /api/v1/links` - Criar link encurtado
- `GET /api/v1/links` - Listar links do usuário
- `GET /:code` - Redirecionar para URL original
- `DELETE /api/v1/links/:code` - Deletar link

### Health & Monitoring
- `GET /health` - Health check
- `GET /metrics` - Métricas da aplicação

### Headers Obrigatórios
```
X-User-ID: user_identifier
Content-Type: application/json
```

## 🔮 Roadmap

### Próximas Versões
- **v0.4.0**: Analytics Dashboard com métricas avançadas
- **v0.5.0**: PWA com modo offline e tema dark/light
- **v1.0.0**: API pública com rate limiting e SDK
- **v2.0.0**: Plataforma enterprise com multi-tenancy

### Versão Enterprise (Privada)
- **🌍 Eco-Analytics**: Métricas de sustentabilidade digital
- **🤖 IA Preditiva**: Otimização automática com machine learning
- **🎨 QR Premium**: Templates dinâmicos e branding corporativo
- **🌐 CDN Verde**: Infraestrutura sustentável com energia renovável

## 🤝 Contribuição

### Padrões de Código
- **Go**: `gofmt`, `golint`, `go vet`
- **JavaScript**: ESLint + Prettier
- **Commits**: Conventional Commits
- **Branches**: GitFlow

### Pull Requests
1. Fork o repositório
2. Crie feature branch (`git checkout -b feature/amazing-feature`)
3. Commit suas mudanças (`git commit -m 'feat: add amazing feature'`)
4. Push para branch (`git push origin feature/amazing-feature`)
5. Abra Pull Request

## 📄 Licença

MIT License - veja [LICENSE](LICENSE) para detalhes.

## 🔗 Links Úteis

- **Documentação**: [Wiki do Projeto](wiki-url)
- **Issues**: [GitHub Issues](issues-url)
- **Changelog**: [CHANGELOG.md](CHANGELOG.md)
- **Roadmap**: [ROADMAP.md](ROADMAP.md)

---

**Desenvolvido com 💚 por Danilo Monteiro**

*Seguindo princípios de Clean Architecture, SOLID e Clean Code*  
*Parte da série "Vibecoding Chronicles: A Jornada de um Arquiteto Full Stack"*

## 📈 Status do Projeto

![Build Status](https://img.shields.io/badge/build-passing-brightgreen)
![Code Quality](https://img.shields.io/badge/code%20quality-A-brightgreen)
![Coverage](https://img.shields.io/badge/coverage-85%25-green)
![License](https://img.shields.io/badge/license-MIT-blue)