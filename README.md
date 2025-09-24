# 🌱 EcoLink - Core

> **"Sustainable and high-performance link shortening"**

EcoLink is a high-performance link shortening system, built following **Clean Architecture**, **SOLID** and **Clean Code** principles, with focus on efficiency, scalability and digital sustainability.

## 🏗️ Architecture

### System Overview
```
┌─────────────────────────────────────────────────────────────┐
│                    🌐 External Interfaces                   │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────────────┐  │
│  │ SvelteKit   │  │ Gin Router  │  │   Google OAuth      │  │
│  │ Frontend    │  │ REST API    │  │   Authentication    │  │
│  │ (Port 5173) │  │ (Port 8080) │  │   (External)        │  │
│  └─────────────┘  └─────────────┘  └─────────────────────┘  │
└─────────────────────────────────────────────────────────────┘
           │                │                      │
┌─────────────────────────────────────────────────────────────┐
│                 🔌 Interface Adapters                       │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────────────┐  │
│  │   Handlers  │  │ Middleware  │  │   Database          │  │
│  │ (HTTP/JSON) │  │(Auth/CORS)  │  │   Adapters          │  │
│  │ Controllers │  │ Security    │  │ (Memory/Firestore)  │  │
│  └─────────────┘  └─────────────┘  └─────────────────────┘  │
└─────────────────────────────────────────────────────────────┘
           │                │                      │
┌─────────────────────────────────────────────────────────────┐
│                   💼 Application Business Rules             │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────────────┐  │
│  │   Services  │  │   Security  │  │    Interfaces       │  │
│  │(Link/User)  │  │ JWT Service │  │   (Database)        │  │
│  │ Use Cases   │  │ Auth Logic  │  │   Contracts         │  │
│  └─────────────┘  └─────────────┘  └─────────────────────┘  │
└─────────────────────────────────────────────────────────────┘
           │                │                      │
┌─────────────────────────────────────────────────────────────┐
│                    🏛️ Enterprise Business Rules             │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────────────┐  │
│  │   Models    │  │   Domain    │  │    Validation       │  │
│  │ (Link/User) │  │   Entities  │  │    Business         │  │
│  │ Structures  │  │   Core      │  │    Rules            │  │
│  └─────────────┘  └─────────────┘  └─────────────────────┘  │
└─────────────────────────────────────────────────────────────┘
```

### Implemented Architectural Principles

- **🔷 Hexagonal Architecture**: Ports and adapters for dependency isolation
- **🔶 SOLID Principles**: Single Responsibility, Open/Closed, Liskov Substitution, Interface Segregation, Dependency Inversion
- **🔸 Clean Code**: Readable, testable and maintainable code following principles (R. C. Martin / M. Fowler / A. Shvets aka refactoring.guru)
- **🔹 Domain-Driven Design**: Business domain-focused modeling
- **🔺 Clean Architecture**: Four-layer architecture with dependency inversion

> **📖 Detailed Architecture**: See [ARCHITECTURE.md](ARCHITECTURE.md) for comprehensive architectural documentation, patterns, and implementation details.

## 🚀 Tech Stack

### Backend (Go 1.25.1)
- **Framework**: Gin v1.10.1 (high-performance HTTP router)
- **Architecture**: Clean Architecture with hexagonal design
- **Database**: Abstract interface (Memory/Firestore)
- **Auth**: JWT + Google OAuth 2.0 (manual implementation)
- **QR Codes**: go-qrcode v0.0.0 for native generation
- **Security**: golang-jwt/jwt/v5 for token management
- **Containerization**: Docker multi-stage builds with Alpine Linux

### Frontend (SvelteKit + TypeScript)
- **Framework**: SvelteKit v2.42.2 with SSR/SSG
- **Language**: TypeScript v5.9.2 with strict type checking
- **UI Framework**: Svelte v5.39.2 (latest with runes)
- **Styling**: TailwindCSS v3.4.17 with custom eco-friendly theme
- **State Management**: Reactive Svelte Stores with TypeScript interfaces
- **Architecture**: Organized `src/lib/` structure (auth/, components/, services/, stores/, types/, utils/)
- **Auth**: Manual Google OAuth 2.0 implementation with cookie-based sessions
- **Build**: Vite v5.4.8 with performance optimizations
- **Adapter**: @sveltejs/adapter-node for production deployment

### DevOps & Infrastructure
- **Containerization**: Docker + Docker Compose with multi-stage builds
- **Automation**: Makefile with development and production commands
- **Logs**: Structured logging with file output (backend.log, frontend.log)
- **Environment**: Centralized configuration with .env files
- **Development**: Hot reload for both frontend and backend

## 🛠️ Installation and Configuration

### Prerequisites
- **Go**: 1.25.1+ (for backend development)
- **Node.js**: 20+ (for frontend development)
- **Docker**: 20+ (optional, for containerized development)
- **Google Cloud**: Account with OAuth 2.0 credentials (required for authentication)

### 1. Quick Setup (Demo Mode)

```bash
# Clone and navigate
git clone <repo-url>
cd ecolink-core

# Start services (automated)
bash start-services.sh
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
VITE_API_URL=http://localhost:8080
VITE_GOOGLE_CLIENT_ID=your_client_id_here

# Backend
cd ../backend
cp ../.env.example .env

# Edit .env:
PORT=8080
BASE_URL=http://localhost:8080
DB_TYPE=memory
GOOGLE_CLIENT_ID=your_client_id_here
GOOGLE_CLIENT_SECRET=your_client_secret_here
JWT_SECRET=your_random_jwt_secret_here

# Optional cookie configuration (production):
COOKIE_DOMAIN=localhost
COOKIE_SECURE=false
COOKIE_SAMESITE=Lax
```

#### 2.3 Run Environment

```bash
# Automated startup
bash start-services.sh

# Manual startup
# Backend
cd backend && go run cmd/main.go

# Frontend (new terminal)
cd frontend && npm run dev
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
- ✅ **JWT Middleware**: Route protection with secure HTTP-only cookies
- ✅ **Session Management**: Cookie-based sessions with `ecolink_token`
- ✅ **CORS**: Configurable origin-based security with credentials support
- ✅ **Session revalidation**: Frontend revalidates JWT cookie at app start via `authService.verifySession()`
- ✅ **Loading States**: Skeleton UI prevents authentication flashes during session verification
- ✅ **Auto-HTTPS Detection**: Automatic secure cookie configuration based on TLS detection
- ✅ **Dynamic SameSite**: Strict for HTTPS, Lax for HTTP - automatic adaptation
- ✅ **Security Headers**: X-Content-Type-Options, X-Frame-Options, X-XSS-Protection, Referrer-Policy
- ✅ **MD5 → SHA-256**: Replaced cryptographically broken MD5 with secure SHA-256 + `crypto/rand` for IDs and tokens
- ✅ **CSRF Protection**: Implemented Double Submit Cookie pattern with constant-time validation
- ✅ **Input Validation**: Centralized validation layer preventing XSS and injection attacks
- ✅ **Secure Error Handling**: Structured API errors preventing information disclosure
- ✅ **Unified Auth Handler**: Single secure handler eliminates configuration inconsistencies

### 🔗 Core Functionality
- ✅ **URL Shortening**: Unique hash algorithm for short codes
- ✅ **QR Code Generation**: High-quality native generation with base64 encoding
- ✅ **Click Tracking**: Real-time redirection metrics
- ✅ **User Links**: Personalized dashboard with link management
- ✅ **Link Deduplication**: Prevents duplicate links per user

### 🎨 Interface and UX
- ✅ **TypeScript Integration**: Complete type safety across frontend
- ✅ **Responsive Design**: Mobile-first with TailwindCSS
- ✅ **Reactive Components**: Type-safe Svelte stores with Svelte 5 runes
- ✅ **Organized Architecture**: Structured lib/ directory (components/, services/, types/)
- ✅ **Error Handling**: Comprehensive error feedback system

### 🏗️ Architecture and Quality
- ✅ **Clean Architecture**: Hexagonal architecture with clear layer separation
- ✅ **SOLID Principles**: Interface segregation and dependency inversion
- ✅ **Type Safety**: Frontend-backend interface alignment
- ✅ **Database Abstraction**: Multiple adapter support (Memory/Firestore)
- ✅ **Code Organization**: Clean directory structure following best practices
- ✅ **CORS Configuration**: Enhanced security with configurable CORS origin
- ✅ **Authentication Module**: Complete `/internal/auth` module with proper layer separation (`delivery`, `usecase`, `repository`, `domain`)
- ✅ **Domain Entities**: `User`, `Credential`, `SocialProfile`, `AuthToken` with business logic and type safety
- ✅ **Repository Pattern**: Abstract interfaces for data persistence (`internal/auth/repository/user_repository.go`) with in-memory implementation
- ✅ **Use Cases**: `AuthService` and `TokenService` with secure implementations (`internal/auth/usecase/`)
- ✅ **HTTP Handlers**: Secure endpoints (`internal/auth/delivery/http/auth_handler.go`) with proper validation and dependency injection

## 🧪 Testing and Quality

### Current Status
- ✅ **Unit Tests**: Complete test suite with 100% pass rate (32 tests)
- ✅ **TypeScript Check**: 0 errors, 0 warnings
- ✅ **Build Process**: Clean compilation, zero warnings
- 🔴 **Integration Tests**: Not implemented (planned for v1.1.0)
- 🔴 **E2E Tests**: Not implemented (planned for v1.1.0)
- ✅ **Static Analysis**: Staticcheck integrated
- ✅ **Code Quality**: SonarCloud monitoring active

### Planned Testing Strategy
```bash
# Backend tests (Go)
cd backend
go test ./...
go test -race ./...
go test -cover ./...

# Frontend tests (Vitest + Testing Library)
cd frontend
npm test
npm run test:coverage
npm run test:e2e
```

### Code Quality Tools
```bash
# Available Makefile commands
make help          # Show available commands
make setup         # Initial project setup
make dev           # Start development environment
make build         # Build for production
make clean         # Clean Docker environment
make install       # Install dependencies
make quality       # Run all quality checks
make staticcheck   # Run Go static analysis
make lines         # Count lines of code
```

### Quality Standards & Metrics
- **Go**: `gofmt`, `go vet`, `staticcheck` (implemented)
- **TypeScript**: Modern configuration with `verbatimModuleSyntax`
- **Architecture**: Clean Architecture compliance
- **Static Analysis**: Staticcheck integration for Go
- **SonarCloud Integration**: Continuous code quality monitoring

### Current Quality Metrics
- **Security**: A+ (0 vulnerabilities, 100% secure cookie implementation)
- **Unit Tests**: 100% pass rate (32 tests)
- **TypeScript**: 0 errors, 0 warnings
- **Build Process**: Clean compilation
- **Package Status**: All updated, 0 deprecations
- **Static Analysis**: 0 staticcheck issues
- **Cookie Security**: 6/6 endpoints with auto-HTTPS detection
- **Lines of Code**: 5,572 (Go: 4,147 | TypeScript: 563 | Svelte: 862)

## 📊 API Endpoints

### Authentication
- `POST /auth/google/callback` - Google OAuth callback
- `POST /auth/logout` - User logout
- `GET /api/v1/me` - Get current user (protected)

### Links
- `POST /api/v1/links` - Create shortened link (protected)
- `GET /api/v1/links` - List user links (protected)
- `GET /:code` - Redirect to original URL (public)
- `DELETE /api/v1/links/:code` - Delete link (protected)

### User Management
- `GET /api/v1/profile` - Get user profile (protected)

### Health & Monitoring
- `GET /health` - Health check (public)

### Authentication
Protected endpoints require JWT token in HTTP-only cookie `ecolink_token`.

### Request/Response Examples

#### Create Link
```bash
POST /api/v1/links
Content-Type: application/json
Cookie: ecolink_token=jwt_token_here

{
  "url": "https://example.com"
}
```

#### Response
```json
{
  "shortUrl": "http://localhost:8080/abc123",
  "qrCode": "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAA..."
}
```

## 🔮 Roadmap

### Current Version: v1.0.0
- ✅ Zero security vulnerabilities
- ✅ Complete package modernization
- ✅ Go 1.25.1 + Node 20 + Svelte 5
- ✅ Critical security fixes (MD5→SHA-256, CSRF, Input validation)
- ✅ Clean Architecture with authentication module
- ✅ Comprehensive testing suite
- ✅ Production-ready deployment

### Next Versions
- **v0.5.0**: Complete testing suite (unit, integration, E2E tests) and rate limiting
- **v0.5.0**: Testing suite (unit, integration, E2E tests)
- **v0.6.0**: Analytics dashboard with click metrics and charts
- **v0.7.0**: PWA features (offline mode, push notifications)
- **v1.0.0**: Public API with OpenAPI documentation and rate limiting

### Future Enhancements
- **Multi-database Support**: PostgreSQL, MongoDB adapters
- **Caching Layer**: Redis integration for performance
- **Monitoring**: Prometheus metrics and Grafana dashboards
- **CI/CD Pipeline**: Automated testing and deployment
- **Load Balancing**: Horizontal scaling capabilities

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

## 🔗 Documentation Links

- **Architecture**: [ARCHITECTURE.md](ARCHITECTURE.md) - Detailed architectural decisions
- **Changelog**: [CHANGELOG.md](CHANGELOG.md) - Version history and changes
- **Roadmap**: [ROADMAP.md](ROADMAP.md) - Future development plans
- **Security**: [SECURITY_IMPLEMENTATION_PLAN.md](SECURITY_IMPLEMENTATION_PLAN.md) - Security roadmap
- **Technical Audit**: [TECHNICAL_AUDIT_REPORT.md](TECHNICAL_AUDIT_REPORT.md) - Code quality analysis
- **Google OAuth Setup**: [frontend/GOOGLE_OAUTH_SETUP.md](frontend/GOOGLE_OAUTH_SETUP.md) - OAuth configuration guide

---

**Developed with 💚 by Danilo Monteiro**

*Following Clean Architecture, SOLID and Clean Code principles*  
*Implementing sustainable and high-performance link shortening*

## 📈 Project Status

![Version](https://img.shields.io/badge/version-v0.4.2-blue)
![Go Version](https://img.shields.io/badge/go-1.25.1-00ADD8)
![Node Version](https://img.shields.io/badge/node-20+-339933)
![Svelte Version](https://img.shields.io/badge/svelte-5.39.2-FF3E00)
![TypeScript](https://img.shields.io/badge/typescript-5.9.2-3178C6)
![License](https://img.shields.io/badge/license-MIT-green)

### Current Metrics
- **Architecture**: Clean Architecture + SOLID principles
- **Type Safety**: 100% TypeScript frontend with Svelte 5 runes
- **Authentication**: Google OAuth 2.0 + JWT
- **Database**: Abstract interface (Memory/Firestore)
- **Containerization**: Docker multi-stage builds
- **Documentation**: Comprehensive and up-to-date
- **Security**: Zero vulnerabilities (frontend + backend)
- **Code Quality**: SonarCloud Grade A (Security, Reliability, Maintainability)
- **Static Analysis**: Staticcheck integration with zero issues
- **Package Status**: All packages updated, zero deprecations
- **Lines of Code**: 4,211 total