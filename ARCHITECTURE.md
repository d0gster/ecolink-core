# 🏗️ EcoLink - Documentação de Arquitetura

## 📋 Visão Geral

EcoLink foi desenvolvido seguindo rigorosamente os princípios de **Clean Architecture**, **SOLID** e **Clean Code**, estabelecidos por Robert C. Martin (Uncle Bob) e Martin Fowler. Esta documentação detalha as decisões arquiteturais e padrões implementados.

## 🎯 Princípios Arquiteturais

### Clean Architecture (Hexagonal)

```
┌─────────────────────────────────────────────────────────────┐
│                    🌐 External Interfaces                   │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────────────┐  │
│  │   Web UI    │  │  REST API   │  │    Database         │  │
│  │ (SvelteKit) │  │   (Gin)     │  │ (Memory/Firestore)  │  │
│  └─────────────┘  └─────────────┘  └─────────────────────┘  │
└─────────────────────────────────────────────────────────────┘
           │                │                      │
┌─────────────────────────────────────────────────────────────┐
│                 🔌 Interface Adapters                       │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────────────┐  │
│  │  Controllers│  │  Presenters │  │    Gateways         │  │
│  │  (Handlers) │  │   (JSON)    │  │  (Repository)       │  │
│  └─────────────┘  └─────────────┘  └─────────────────────┘  │
└─────────────────────────────────────────────────────────────┘
           │                │                      │
┌─────────────────────────────────────────────────────────────┐
│                   💼 Application Business Rules             │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────────────┐  │
│  │   Use Cases │  │  Services   │  │    Interfaces       │  │
│  │ (Link CRUD) │  │(Link/User)  │  │   (Database)        │  │
│  └─────────────┘  └─────────────┘  └─────────────────────┘  │
└─────────────────────────────────────────────────────────────┘
           │                │                      │
┌─────────────────────────────────────────────────────────────┐
│                    🏛️ Enterprise Business Rules             │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────────────┐  │
│  │   Entities  │  │   Models    │  │    Domain Logic     │  │
│  │ (Link/User) │  │ (Structs)   │  │   (Validation)      │  │
│  └─────────────┘  └─────────────┘  └─────────────────────┘  │
└─────────────────────────────────────────────────────────────┘
```

### SOLID Principles

#### 🔸 Single Responsibility Principle (SRP)
- **Handlers**: Apenas processamento HTTP
- **Services**: Apenas lógica de negócio
- **Models**: Apenas representação de dados
- **Database**: Apenas persistência

#### 🔸 Open/Closed Principle (OCP)
- **Database Interface**: Extensível para novos adapters
- **Middleware**: Composável e extensível
- **Services**: Abertos para extensão via interfaces

#### 🔸 Liskov Substitution Principle (LSP)
- **Database Implementations**: Memory e Firestore são intercambiáveis
- **Auth Providers**: Google OAuth pode ser substituído

#### 🔸 Interface Segregation Principle (ISP)
- **Database Interface**: Métodos específicos por funcionalidade
- **Service Interfaces**: Contratos mínimos e focados

#### 🔸 Dependency Inversion Principle (DIP)
- **High-level modules**: Dependem de abstrações
- **Dependency Injection**: Implementado via constructors

## 🏛️ Estrutura do Backend (Go)

### Organização de Diretórios

```
backend/
├── cmd/                    # Application entry points
│   └── main.go            # Main application
├── internal/              # Private application code
│   ├── config/           # Configuration management
│   ├── handlers/         # HTTP handlers (Controllers)
│   ├── middleware/       # HTTP middleware
│   ├── models/          # Domain models (Entities)
│   └── services/        # Business logic (Use Cases)
├── pkg/                  # Public library code
│   ├── database/        # Database interfaces and implementations
│   └── utils/           # Utility functions
└── go.mod              # Go modules
```

### Camadas e Responsabilidades

#### 🎯 Domain Layer (`internal/models/`)
```go
// Link entity - Core business object
type Link struct {
    ID          string    `json:"id"`
    ShortCode   string    `json:"short_code"`
    OriginalURL string    `json:"original_url"`
    UserID      string    `json:"user_id"`
    CreatedAt   time.Time `json:"created_at"`
    Clicks      int       `json:"clicks"`
    QRCode      string    `json:"qr_code"`
}

// Business rules validation
func (l *Link) Validate() error {
    if l.OriginalURL == "" {
        return errors.New("original URL is required")
    }
    // Additional validation logic
    return nil
}
```

#### 💼 Application Layer (`internal/services/`)
```go
// LinkService - Use cases implementation
type LinkService struct {
    db database.Database
}

// CreateLink - Use case: Create shortened link
func (s *LinkService) CreateLink(url, userID string) (*models.LinkResponse, error) {
    // 1. Validate input
    // 2. Generate short code
    // 3. Create QR code
    // 4. Save to database
    // 5. Return response
}
```

#### 🔌 Interface Adapters (`internal/handlers/`)
```go
// LinkHandler - HTTP controller
type LinkHandler struct {
    linkService *services.LinkService
}

// CreateLink - HTTP handler
func (h *LinkHandler) CreateLink(c *gin.Context) {
    // 1. Parse HTTP request
    // 2. Call use case
    // 3. Format HTTP response
}
```

#### 🌐 Infrastructure Layer (`pkg/database/`)
```go
// Database interface - Port
type Database interface {
    SaveLink(link *models.Link) error
    GetLink(code string) (*models.Link, error)
    // Other methods...
}

// MemoryDatabase - Adapter implementation
type MemoryDatabase struct {
    links map[string]*models.Link
    mutex sync.RWMutex
}
```

## 🎨 Estrutura do Frontend (SvelteKit)

### Organização de Diretórios

```
frontend/src/
├── lib/                   # Shared libraries
│   ├── auth/             # Authentication logic
│   ├── stores/           # State management
│   ├── components/       # Reusable components
│   └── api.js           # API client
├── routes/               # Page components (File-based routing)
│   ├── auth/            # Authentication pages
│   ├── dashboard/       # User dashboard
│   └── +layout.svelte   # Root layout
└── app.html             # HTML template
```

### Padrões Implementados

#### 🔄 State Management (Reactive Stores)
```javascript
// auth.js - Authentication store
import { writable } from 'svelte/store';

export const user = writable(null);
export const pendingLink = writable(null);

export const auth = {
    async loginWithProvider(provider) {
        // OAuth flow implementation
    },
    
    logout() {
        // Clean logout with state cleanup
    },
    
    setUser(userData) {
        // Secure user data persistence
    }
};
```

#### 🔐 Authentication Flow
```javascript
// google.js - OAuth 2.0 implementation
export function getGoogleAuthUrl() {
    const params = new URLSearchParams({
        client_id: GOOGLE_CLIENT_ID,
        redirect_uri: REDIRECT_URI,
        response_type: 'code',
        scope: 'openid email profile'
    });
    
    return `https://accounts.google.com/oauth/authorize?${params}`;
}

export async function exchangeCodeForTokens(code) {
    // RFC 6749 compliant token exchange
}
```

## 🔐 Segurança e Autenticação

### OAuth 2.0 Implementation

#### Authorization Code Flow
1. **Authorization Request**: Redirect to Google
2. **Authorization Grant**: User consent
3. **Authorization Code**: Callback with code
4. **Access Token**: Exchange code for token
5. **Protected Resource**: API calls with token

#### Security Measures
- **CSRF Protection**: State parameter validation
- **Token Security**: Secure storage and transmission
- **Session Management**: Proper cleanup on logout
- **Input Validation**: All inputs sanitized
- **CORS Configuration**: Proper origin validation

## 📊 Qualidade e Métricas

### Code Quality Metrics
- **Cyclomatic Complexity**: < 10 per function
- **Function Length**: < 50 lines
- **File Length**: < 500 lines
- **Test Coverage**: > 80%

### Performance Benchmarks
- **API Response Time**: < 100ms
- **Database Query Time**: < 50ms
- **Frontend Load Time**: < 2s
- **Bundle Size**: < 500KB

### Clean Code Principles Applied

#### 🧹 Meaningful Names
```go
// ❌ Bad
func p(u string) string { ... }

// ✅ Good
func generateShortCode(originalURL string) string { ... }
```

#### 🧹 Small Functions
```go
// ✅ Single responsibility
func validateURL(url string) error {
    if url == "" {
        return errors.New("URL cannot be empty")
    }
    
    if !isValidURL(url) {
        return errors.New("invalid URL format")
    }
    
    return nil
}
```

#### 🧹 Error Handling
```go
// ✅ Explicit error handling
func (s *LinkService) CreateLink(url, userID string) (*models.LinkResponse, error) {
    if err := validateURL(url); err != nil {
        return nil, fmt.Errorf("validation failed: %w", err)
    }
    
    link, err := s.buildLink(url, userID)
    if err != nil {
        return nil, fmt.Errorf("failed to build link: %w", err)
    }
    
    return s.saveAndRespond(link)
}
```

## 🧪 Testing Strategy

### Test Pyramid
```
    ┌─────────────────┐
    │   E2E Tests     │  ← Few, High-level
    │   (Playwright)  │
    ├─────────────────┤
    │ Integration     │  ← Some, API + DB
    │ Tests (Go)      │
    ├─────────────────┤
    │   Unit Tests    │  ← Many, Fast
    │ (Go + Vitest)   │
    └─────────────────┘
```

### Test Categories
- **Unit Tests**: Individual functions and methods
- **Integration Tests**: API endpoints with database
- **Component Tests**: Svelte components
- **E2E Tests**: Complete user workflows

## 🚀 Deployment Architecture

### Production Environment
```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Load Balancer │    │   Cloud Run     │    │   Firestore     │
│   (Cloud LB)    │◄──►│   (Backend)     │◄──►│   (Database)    │
└─────────────────┘    └─────────────────┘    └─────────────────┘
         │                       │
┌─────────────────┐    ┌─────────────────┐
│   CDN           │    │   Cloud Storage │
│   (Frontend)    │    │   (Assets)      │
└─────────────────┘    └─────────────────┘
```

### Scalability Considerations
- **Horizontal Scaling**: Stateless backend design
- **Database Sharding**: Prepared for data partitioning
- **Caching Strategy**: Redis for high-frequency data
- **CDN Integration**: Global asset distribution

## 📚 Referências e Inspirações

### Livros Aplicados
- **Clean Architecture** - Robert C. Martin
- **Clean Code** - Robert C. Martin
- **The Clean Coder** - Robert C. Martin
- **Refactoring** - Martin Fowler
- **Domain-Driven Design** - Eric Evans

### Padrões Implementados
- **Repository Pattern**: Database abstraction
- **Dependency Injection**: Loose coupling
- **Factory Pattern**: Object creation
- **Strategy Pattern**: Algorithm selection
- **Observer Pattern**: Reactive state management

---

**Arquiteto**: Danilo Monteiro  
**Princípios**: Clean Architecture, SOLID, DDD  
**Qualidade**: Code Review, Testing, Documentation  
**Mentoria**: Robert C. Martin, Martin Fowler