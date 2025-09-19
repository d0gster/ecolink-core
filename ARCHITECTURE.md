# 🏗️ EcoLink - Architecture Documentation

## 📋 Overview

EcoLink was developed strictly following **Clean Architecture**, **SOLID**, and **Clean Code** principles, established by Robert C. Martin (Uncle Bob) and Martin Fowler. This documentation details the architectural decisions and implemented patterns.

## 🎯 Architectural Principles

### Clean Architecture (Hexagonal) - Current Implementation

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

### SOLID Principles

#### 🔸 Single Responsibility Principle (SRP)
- **Handlers**: Only HTTP processing
- **Services**: Only business logic
- **Models**: Only data representation
- **Database**: Only persistence

#### 🔸 Open/Closed Principle (OCP)
- **Database Interface**: Extensible for new adapters
- **Middleware**: Composable and extensible
- **Services**: Open for extension via interfaces

#### 🔸 Liskov Substitution Principle (LSP)
- **Database Implementations**: Memory and Firestore are interchangeable
- **Auth Providers**: Google OAuth can be replaced

#### 🔸 Interface Segregation Principle (ISP)
- **Database Interface**: Specific methods per functionality
- **Service Interfaces**: Minimal and focused contracts

#### 🔸 Dependency Inversion Principle (DIP)
- **High-level modules**: Depend on abstractions
- **Dependency Injection**: Implemented via constructors

## 🏛️ Backend Structure (Go 1.21+)

### Current Directory Organization

```
backend/
├── cmd/
│   └── main.go                  # Application entry point & DI composition
├── internal/                    # Private application code
│   ├── config/
│   │   └── config.go           # Configuration management + cookie settings
│   ├── handlers/               # HTTP handlers (Interface Adapters)
│   │   ├── auth_handler.go     # Google OAuth + JWT authentication
│   │   ├── link_handler.go     # Link CRUD operations
│   │   └── user_handler.go     # User profile management
│   ├── middleware/             # Cross-cutting concerns
│   │   └── auth.go            # JWT middleware + security headers + CORS
│   ├── models/                 # Domain entities
│   │   ├── link.go            # Link entity + request/response DTOs
│   │   └── user.go            # User entity + Google OAuth DTOs
│   ├── security/               # Security services
│   │   └── jwt_service.go     # JWT generation & validation
│   └── services/               # Business logic (Use Cases)
│       ├── link_service.go    # Link shortening + QR generation + deduplication
│       └── user_service.go    # User management + Google OAuth integration
├── pkg/                        # Public library code
│   ├── database/              # Database abstraction layer
│   │   ├── firestore.go       # Firestore adapter (production)
│   │   ├── interface.go       # Database interface contract
│   │   └── memory.go          # In-memory adapter (development)
│   └── utils/
│       └── shortener.go       # URL shortening algorithm
├── .env                       # Environment variables
├── go.mod                     # Go modules (Go 1.21)
└── go.sum                     # Dependency checksums
```

### Layers and Responsibilities

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

## 🎨 Frontend Structure (SvelteKit + TypeScript)

### Current Directory Organization

```
frontend/
├── src/
│   ├── lib/                   # Shared libraries (TypeScript)
│   │   ├── auth/             # Authentication logic
│   │   │   ├── auth-guard.ts     # Route protection guards
│   │   │   ├── google-direct.ts  # Google OAuth implementation
│   │   │   └── guard.ts          # Guard utilities
│   │   ├── components/       # Reusable Svelte components
│   │   │   ├── BackgroundVideo.svelte
│   │   │   └── UserDropdown.svelte
│   │   ├── config/
│   │   │   └── env.ts           # Environment configuration
│   │   ├── services/         # API & business logic
│   │   │   ├── api.ts           # HTTP client
│   │   │   └── authService.ts   # Authentication service (SRP)
│   │   ├── stores/           # Reactive state management
│   │   │   └── auth.ts          # Authentication stores
│   │   ├── types/            # TypeScript interfaces
│   │   │   ├── link.ts          # Link-related types
│   │   │   ├── profileData.ts   # User profile types
│   │   │   ├── provider.ts      # OAuth provider types
│   │   │   └── result.ts        # API response types
│   │   └── utils/
│   │       └── constants.ts     # Application constants
│   ├── routes/               # SvelteKit file-based routing
│   │   ├── auth/            # Authentication pages
│   │   │   ├── callback/
│   │   │   │   └── google/
│   │   │   │       └── +page.svelte  # Google OAuth callback
│   │   │   └── +page.svelte     # Auth landing page
│   │   ├── dashboard/       # User dashboard
│   │   │   └── +page.svelte     # Links management
│   │   ├── profile/         # User profile
│   │   │   └── +page.svelte     # Profile management
│   │   ├── result/          # Link creation result
│   │   │   └── +page.svelte     # QR code display
│   │   ├── +layout.svelte   # Root layout + session rehydration
│   │   └── +page.svelte     # Homepage + link shortening form
│   ├── app.css              # Global styles + TailwindCSS
│   ├── app.d.ts             # TypeScript global declarations
│   └── app.html             # HTML template
├── static/                  # Static assets
│   ├── images/
│   └── videos/              # Background videos
├── .env.example             # Environment template
├── .env.local               # Local environment (gitignored)
├── package.json             # Dependencies + scripts
├── svelte.config.js         # SvelteKit configuration
├── tailwind.config.js       # TailwindCSS + eco theme
├── tsconfig.json            # TypeScript configuration
└── vite.config.ts           # Vite build configuration
```

### Implemented Patterns

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

## 🔐 Security and Authentication

### Google OAuth 2.0 Implementation (Manual)

#### Implemented Authorization Code Flow
1. **Authorization Request**: Frontend redirects to `accounts.google.com`
2. **User Consent**: User authorizes access (openid, profile, email)
3. **Authorization Code**: Google redirects to `/auth/callback/google` with code
4. **Token Exchange**: Backend exchanges code for access_token via `oauth2.googleapis.com/token`
5. **User Info**: Backend fetches user data via `googleapis.com/oauth2/v2/userinfo`
6. **JWT Generation**: Backend generates JWT and sets HTTP-only cookie
7. **Session Management**: Frontend stores state in Svelte stores

#### Detailed Authentication Flow

```typescript
// Frontend: authService.ts
export function login(state?: string) {
  const params = new URLSearchParams({
    client_id: config.googleClientId,
    redirect_uri: `${window.location.origin}/auth/callback/google`,
    response_type: 'code',
    scope: 'openid profile email'
  });
  
  if (state) params.append('state', state);
  window.location.href = `https://accounts.google.com/o/oauth2/v2/auth?${params}`;
}

export async function verifySession(): Promise<User | null> {
  const response = await fetch(`${config.apiUrl}/api/v1/me`, {
    credentials: 'include' // Includes ecolink_token cookie
  });
  
  if (response.ok) {
    const user = await response.json();
    userStore.set(user);
    isAuthenticated.set(true);
    return user;
  }
  
  return null;
}
```

```go
// Backend: auth_handler.go
func (h *AuthHandler) GoogleCallback(c *gin.Context) {
    // 1. Extract code from request
    var req struct {
        Code        string `json:"code" binding:"required"`
        RedirectURI string `json:"redirect_uri" binding:"required"`
        State       string `json:"state,omitempty"`
    }
    
    // 2. Exchange code for access_token
    token, err := h.exchangeCodeForToken(req.Code, req.RedirectURI)
    
    // 3. Fetch user information
    userInfo, err := h.getUserInfo(token.AccessToken)
    
    // 4. Create/update user in system
    user, err := h.userService.CreateOrUpdateUser(
        userInfo.ID, userInfo.Name, userInfo.Email, userInfo.Picture)
    
    // 5. Generate JWT
    jwtToken, err := h.jwtService.GenerateToken(user.ID, user.Email)
    
    // 6. Set HTTP-only cookie with security settings
    cookie := &http.Cookie{
        Name:     "ecolink_token",
        Value:    jwtToken,
        Path:     "/",
        Domain:   h.cfg.CookieDomain,
        MaxAge:   3600 * 24 * 30, // 30 days
        HttpOnly: true,
        Secure:   h.cfg.CookieSecure,
        SameSite: parseSameSite(h.cfg.CookieSameSite),
    }
    
    http.SetCookie(c.Writer, cookie)
}
```

#### Implemented Security Measures
- **HTTP-Only Cookies**: JWT token stored in secure cookie
- **Cookie Configuration**: Configurable Domain, Secure flag and SameSite
- **JWT Validation**: Middleware validates token on protected routes
- **Session Rehydration**: Frontend revalidates session on page load
- **CORS Configuration**: Specific origin configured for security
- **Security Headers**: X-Content-Type-Options, X-Frame-Options, X-XSS-Protection
- **State Management**: Loading states prevent authentication flashes

#### Security Configuration

```bash
# Development
COOKIE_DOMAIN=localhost
COOKIE_SECURE=false
COOKIE_SAMESITE=Lax

# Production
COOKIE_DOMAIN=yourdomain.com
COOKIE_SECURE=true
COOKIE_SAMESITE=Strict
```

## 📊 Quality and Metrics

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

## 📚 References and Inspirations

### Applied Books
- **Clean Architecture** - Robert C. Martin
- **Clean Code** - Robert C. Martin
- **The Clean Coder** - Robert C. Martin
- **Refactoring** - Martin Fowler
- **Domain-Driven Design** - Eric Evans

### Implemented Patterns
- **Repository Pattern**: Database abstraction
- **Dependency Injection**: Loose coupling
- **Factory Pattern**: Object creation
- **Strategy Pattern**: Algorithm selection
- **Observer Pattern**: Reactive state management

---

**Architect**: Danilo Monteiro  
**Principles**: Clean Architecture, SOLID, DDD  
**Quality**: Code Review, Testing, Documentation  
**Mentorship**: Robert C. Martin, Martin Fowler, Alexander Shvets