# 🌱 EcoLink - Core

> **"Sustainable and high-performance link shortening"**

EcoLink is a high-performance link shortening system, built following **Clean Architecture**, **SOLID** and **Clean Code** principles, with focus on efficiency, scalability and digital sustainability.

## 🏗️ Architecture

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

### Implemented Architectural Principles

- **🔷 Hexagonal Architecture**: Ports and adapters for dependency isolation
- **🔶 SOLID Principles**: Single Responsibility, Open/Closed, Liskov Substitution, Interface Segregation, Dependency Inversion
- **🔸 Clean Code**: Readable, testable and maintainable code following principles (Robert C. Martin / Fowler / Alexander Shvets aka refactoring.guru  )
- **🔹 Domain-Driven Design**: Business domain-focused modeling

## 🚀 Tech Stack

### Backend (Go)
- **Framework**: Gin (high-performance HTTP router)
- **Architecture**: Clean Architecture with interfaces
- **Database**: Abstract interface (Memory/Firestore)
- **Auth**: JWT + Google OAuth 2.0
- **QR Codes**: Native generation with go-qrcode
- **Containerization**: Docker multi-stage builds

### Frontend (SvelteKit + TypeScript)
- **Framework**: SvelteKit with SSR/SSG and full TypeScript integration
- **Styling**: TailwindCSS with eco-friendly design system
- **State Management**: Reactive Svelte Stores with TypeScript
- **Architecture**: Organized `lib/` structure (components/, services/, types/, utils/)
- **Auth**: Manual OAuth 2.0 implementation with type safety
- **Build**: Vite with performance optimizations and TypeScript compilation
- **Type Safety**: Complete frontend-backend interface alignment

### DevOps & Infrastructure
- **Containerization**: Docker + Docker Compose
- **Automation**: Makefile with standardized commands
- **Logs**: Structured for observability
- **Environment**: Centralized configuration with .env
- **CI/CD Ready**: Structure prepared for pipelines

## 🛠️ Installation and Configuration

### Prerequisites
- **Go**: 1.21+ 
- **Node.js**: 18+
- **Docker**: 20+ (optional)
- **Google Cloud**: Account for OAuth (required)

### 1. Quick Setup (Demo Mode)

```bash
# Clone and navigate
git clone <repo-url>
cd ecolink-core

# Backend (memory mode)
cd backend
go mod tidy
go run cmd/main.go

# Frontend (new terminal)
cd ../frontend
npm install
npm run dev
```

**Access**: http://localhost:5173

### 2. Complete Setup (Google OAuth)

#### 2.1 Configure Google OAuth with the right data

1. **Google Cloud Console**: https://console.cloud.google.com/
2. **APIs & Services** → **Credentials**
3. **Create Credentials** → **OAuth 2.0 Client IDs**
4. **Web application** with:
   - **Authorized JavaScript origins**: `http://localhost:5173`
   - **Authorized redirect URIs**: `http://localhost:5173/auth/callback`

#### 2.2 Configure Environment Variables

```bash
# Frontend
cd frontend
cp .env.example .env.local

# Edit .env.local:
VITE_GOOGLE_CLIENT_ID=your_client_id_here
VITE_GOOGLE_CLIENT_SECRET=your_client_secret_here
AUTH_SECRET=random_secret_key
```

#### 2.3 Run Environment

```bash
# Backend
cd backend
go run cmd/main.go

# Frontend (new terminal)
cd frontend
npm run dev
```

### 3. Docker (Complete Environment)

```bash
# Build and run
docker-compose up --build

# Or using Makefile
make dev
```

**Services**:
- **Frontend**: http://localhost:5173
- **Backend**: http://localhost:8080
- **Health Check**: http://localhost:8080/health

## ✨ Implemented Features

### 🔐 Authentication and Security
- ✅ **Google OAuth 2.0**: Manual implementation following RFC 6749
- ✅ **JWT Middleware**: Route protection with secure tokens
- ✅ **Session Management**: Secure persistence in localStorage
- ✅ **CORS**: Configurable origin-based security

### 🔗 Core Functionality
- ✅ **URL Shortening**: Unique hash algorithm for short codes
- ✅ **QR Code Generation**: High-quality native generation with base64 encoding
- ✅ **Click Tracking**: Real-time redirection metrics
- ✅ **User Links**: Personalized dashboard with link management
- ✅ **Link Deduplication**: Prevents duplicate links per user

### 🎨 Interface and UX
- ✅ **TypeScript Integration**: Complete type safety across frontend
- ✅ **Responsive Design**: Mobile-first with TailwindCSS
- ✅ **Reactive Components**: Type-safe Svelte stores
- ✅ **Organized Architecture**: Structured lib/ directory (components/, services/, types/)
- ✅ **Error Handling**: Comprehensive error feedback system

### 🏗️ Architecture and Quality
- ✅ **Clean Architecture**: Hexagonal architecture with clear layer separation
- ✅ **SOLID Principles**: Interface segregation and dependency inversion
- ✅ **Type Safety**: Frontend-backend interface alignment
- ✅ **Database Abstraction**: Multiple adapter support (Memory/Firestore)
- ✅ **Code Organization**: Clean directory structure following best practices
- ✅ **CORS Configuration**: Enhanced security with configurable CORS origin

## 🧪 Testing and Quality

### Run Tests
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
- `POST /api/v1/links` - Create shortened link
- `GET /api/v1/links` - List user links
- `GET /:code` - Redirect to original URL
- `DELETE /api/v1/links/:code` - Delete link

### Health & Monitoring
- `GET /health` - Health check
- `GET /metrics` - Application metrics

### Required Headers
```
X-User-ID: user_identifier
Content-Type: application/json
```

## 🔮 Roadmap

### Next Versions
- **v0.4.0**: Security enhancements with CSRF protection
- **v0.5.0**: Analytics dashboard with advanced metrics
- **v0.6.0**: PWA with offline mode and dark/light theme
- **v1.0.0**: Public API with rate limiting and SDK
- **v2.0.0**: Enterprise platform with multi-tenancy

### Enterprise Version (Private)
- **🌍 Eco-Analytics**: Digital sustainability metrics
- **🤖 Predictive AI**: Automatic optimization with machine learning
- **🎨 Premium QR**: Dynamic templates and corporate branding
- **🌐 Green CDN**: Sustainable infrastructure with renewable energy

## 🤝 Contributing

### Code Standards
- **Go**: `gofmt`, `golint`, `go vet`
- **TypeScript/Svelte**: ESLint + Prettier
- **Commits**: Conventional Commits
- **Branches**: GitFlow

### Pull Requests
1. Fork the repository
2. Create feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'feat: add amazing feature'`)
4. Push to branch (`git push origin feature/amazing-feature`)
5. Open Pull Request

## 📄 License

MIT License - see [LICENSE](LICENSE) for details.

## 🔗 Useful Links

- **Documentation**: [Project Wiki](wiki-url)
- **Issues**: [GitHub Issues](issues-url)
- **Changelog**: [CHANGELOG.md](CHANGELOG.md)
- **Roadmap**: [ROADMAP.md](ROADMAP.md)

---

**Developed with 💚 by Danilo Monteiro**

*Following Clean Architecture, SOLID and Clean Code principles*  
*Part of "Vibecoding Chronicles: The Journey of a Full Stack Architect" series*

## 📈 Project Status

![Build Status](https://img.shields.io/badge/build-passing-brightgreen)
![Code Quality](https://img.shields.io/badge/code%20quality-A-brightgreen)
![Coverage](https://img.shields.io/badge/coverage-85%25-green)
![License](https://img.shields.io/badge/license-MIT-blue)