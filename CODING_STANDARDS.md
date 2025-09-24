# 📋 EcoLink - Coding Standards

**Based on**: Clean Code (Robert C. Martin), SOLID Principles, Go Best Practices, TypeScript Guidelines  
**Version**: 1.0  
**Last Updated**: January 2024

---

## 🎯 Development Philosophy

### Fundamental Principles
- **Clean Code**: Readable, testable and maintainable code
- **SOLID Principles**: Design oriented to solid principles
- **DRY (Don't Repeat Yourself)**: Avoid code duplication
- **KISS (Keep It Simple, Stupid)**: Simplicity over complexity
- **YAGNI (You Aren't Gonna Need It)**: Implement only what's necessary

---

## 🔧 Backend (Go)

### File Structure

#### Directory Organization
```
backend/
├── cmd/                    # Application entry points
├── internal/              # Private application code
│   ├── config/           # Configuration management
│   ├── handlers/         # HTTP controllers
│   ├── middleware/       # HTTP middleware
│   ├── models/          # Domain entities
│   ├── security/        # Security services
│   └── services/        # Business logic
└── pkg/                  # Public code/libraries
    ├── database/        # Database interfaces and adapters
    └── utils/           # Utilities
```

### Naming Conventions

#### Files and Packages
```go
// ✅ Correct
package handlers
// auth_handler.go
// link_handler.go
// user_handler.go

// ❌ Incorrect
package AuthHandlers
// AuthHandler.go
// linkHandler.go
```

#### Variables and Functions
```go
// ✅ Correct - CamelCase for exported
func NewLinkService(db database.Database) *LinkService

// ✅ Correct - camelCase for private
func generateShortCode(url string) string

// ✅ Correct - Descriptive names
var googleClientID string
var jwtSecretKey []byte

// ❌ Incorrect - Generic names
var data string
var temp []byte
```

#### Constants
```go
// ✅ Correct
const (
    DefaultPort = "8080"
    JWTExpirationHours = 24
    MaxRetryAttempts = 3
)

// ❌ Incorrect
const port = "8080"
const JWT_EXPIRATION = 24
```

### Function Structure

#### Function Signatures
```go
// ✅ Correct - Clear parameters, error return
func (s *LinkService) CreateLink(originalURL, userID string) (*models.CreateLinkResponse, error) {
    // Implementation
}

// ✅ Correct - Validation at the beginning
func (h *LinkHandler) CreateLink(c *gin.Context) {
    var req models.CreateLinkRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    // Main logic
}

// ❌ Incorrect - Too many parameters
func CreateLink(url, userID, baseURL, qrSize, format string, enableTracking bool) error
```

#### Error Handling
```go
// ✅ Correct - Specific and contextualized error
func (s *LinkService) GetOriginalURL(shortCode string) (string, error) {
    link, err := s.db.GetLink(shortCode)
    if err != nil {
        return "", fmt.Errorf("failed to retrieve link %s: %w", shortCode, err)
    }
    
    if err := s.db.IncrementClicks(shortCode); err != nil {
        // Log error but don't fail the request
        log.Printf("Failed to increment clicks for %s: %v", shortCode, err)
    }
    
    return link.URL, nil
}

// ❌ Incorrect - Generic error
func (s *LinkService) GetOriginalURL(shortCode string) (string, error) {
    link, err := s.db.GetLink(shortCode)
    if err != nil {
        return "", err // Loses context
    }
    return link.URL, nil
}
```

### Structures and Interfaces

#### Struct Definition
```go
// ✅ Correct - Aligned JSON and Firestore tags
type Link struct {
    URL       string    `json:"originalUrl" firestore:"originalUrl"`
    Code      string    `json:"shortCode" firestore:"shortCode"`
    UserID    string    `json:"userId" firestore:"userId"`
    CreatedAt time.Time `json:"createdAt" firestore:"createdAt"`
    Clicks    int       `json:"clickCount" firestore:"clickCount"`
}

// ✅ Correct - Validation with binding tags
type CreateLinkRequest struct {
    URL string `json:"url" binding:"required,url"`
}
```

#### Interfaces
```go
// ✅ Correct - Focused and specific interface
type Database interface {
    SaveLink(link *models.Link) error
    GetLink(code string) (*models.Link, error)
    GetUserLinks(userID string) ([]*models.Link, error)
    IncrementClicks(code string) error
    DeleteLink(code string) error
    SaveUser(user *models.User) error
    GetUser(id string) (*models.User, error)
    GetUserByGoogleID(googleID string) (*models.User, error)
}

// ❌ Incorrect - Too broad interface
type Repository interface {
    Save(interface{}) error
    Get(interface{}) error
    Delete(interface{}) error
    // ... many generic methods
}
```

### Middleware and Handlers

#### Middleware Pattern
```go
// ✅ Correct - Focused middleware
func JWTAuthMiddleware(jwtService *security.JWTService) gin.HandlerFunc {
    return func(c *gin.Context) {
        token, err := c.Cookie("ecolink_token")
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
            c.Abort()
            return
        }

        claims, err := jwtService.ValidateToken(token)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        c.Set("user_id", claims.UserID)
        c.Set("user_email", claims.Email)
        c.Next()
    }
}
```

---

## 🎨 Frontend (TypeScript + Svelte)

### File Structure

#### Directory Organization
```
src/
├── lib/
│   ├── auth/              # Authentication logic
│   ├── components/        # Reusable components
│   ├── config/           # Configuration
│   ├── services/         # Services and API
│   ├── stores/           # Reactive state
│   ├── types/            # TypeScript interfaces
│   └── utils/            # Utilities
└── routes/               # Pages (file-based routing)
```

### TypeScript Conventions

#### Interfaces and Types
```typescript
// ✅ Correct - Descriptive interfaces
export interface User {
    id: string;
    name: string;
    email: string;
    picture: string;
}

export interface CreateLinkRequest {
    url: string;
}

export interface CreateLinkResponse {
    shortUrl: string;
    qrCode: string;
}

// ✅ Correct - Union types for states
export type AuthState = 'loading' | 'authenticated' | 'unauthenticated';

// ❌ Incorrect - Too generic types
interface Data {
    value: any;
    type: string;
}
```

#### Functions and Services
```typescript
// ✅ Correct - Function with explicit types
export async function verifySession(): Promise<User | null> {
    if (!browser) return null;

    isLoading.set(true);

    try {
        const response = await fetch(`${config.apiUrl}/api/v1/me`, {
            credentials: 'include'
        });

        if (!response.ok) {
            userStore.set(null);
            isAuthenticated.set(false);
            return null;
        }

        const data: User = await response.json();
        userStore.set(data);
        isAuthenticated.set(true);
        return data;
    } catch (err) {
        console.error('verifySession error:', err);
        userStore.set(null);
        isAuthenticated.set(false);
        return null;
    } finally {
        isLoading.set(false);
    }
}

// ❌ Incorrect - No types, no error handling
export async function verifySession() {
    const response = await fetch('/api/v1/me');
    const data = await response.json();
    return data;
}
```

### Svelte Components

#### Component Structure
```svelte
<!-- ✅ Correct - Well-structured component -->
<script lang="ts">
    import { onMount } from 'svelte';
    import type { User } from '$lib/types/user';
    
    export let user: User;
    export let showDropdown = false;
    
    let dropdownElement: HTMLElement;
    
    onMount(() => {
        // Initialization logic
    });
    
    function handleLogout() {
        // Specific logic
    }
</script>

<div class="user-dropdown" bind:this={dropdownElement}>
    <button on:click={() => showDropdown = !showDropdown}>
        <img src={user.picture} alt={user.name} />
        <span>{user.name}</span>
    </button>
    
    {#if showDropdown}
        <div class="dropdown-menu">
            <a href="/profile">Profile</a>
            <button on:click={handleLogout}>Logout</button>
        </div>
    {/if}
</div>

<style>
    .user-dropdown {
        position: relative;
    }
    
    .dropdown-menu {
        position: absolute;
        top: 100%;
        right: 0;
        background: white;
        border: 1px solid #e5e7eb;
        border-radius: 0.5rem;
        box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1);
    }
</style>
```

#### Stores (Reactive State)
```typescript
// ✅ Correct - Typed and focused stores
import { writable } from 'svelte/store';
import type { Link } from '../types/link';

export interface User {
    id: string;
    name: string;
    email: string;
    picture: string;
}

export const user = writable<User | null>(null);
export const isAuthenticated = writable(false);
export const isLoading = writable(true);
export const pendingLink = writable<Link | null>(null);

// ❌ Incorrect - Too generic store
export const appState = writable<any>({});
```

---

## 🗄️ Database

### Interface Pattern
```go
// ✅ Correct - Domain-specific interface
type Database interface {
    // Link operations
    SaveLink(link *models.Link) error
    GetLink(code string) (*models.Link, error)
    GetUserLinks(userID string) ([]*models.Link, error)
    IncrementClicks(code string) error
    DeleteLink(code string) error
    
    // User operations
    SaveUser(user *models.User) error
    GetUser(id string) (*models.User, error)
    GetUserByGoogleID(googleID string) (*models.User, error)
}
```

### Memory vs Firestore Implementation
```go
// ✅ Correct - Interchangeable implementations
type MemoryDB struct {
    links map[string]*models.Link
    users map[string]*models.User
    mutex sync.RWMutex
}

type FirestoreDB struct {
    client    *firestore.Client
    projectID string
}

// Both implement the same Database interface
```

---

## 🔒 Security

### Authentication
```go
// ✅ Correct - JWT with secure configuration
func (j *JWTService) GenerateToken(userID, email string) (string, error) {
    claims := &Claims{
        UserID: userID,
        Email:  email,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            Issuer:    j.issuer,
            Audience:  []string{j.audience},
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(j.secretKey)
}
```

### Secure Cookies
```go
// ✅ Correct - Cookie with security settings
cookie := &http.Cookie{
    Name:     "ecolink_token",
    Value:    jwtToken,
    Path:     "/",
    Domain:   h.cfg.CookieDomain,
    MaxAge:   3600 * 24 * 30,
    HttpOnly: true,
    Secure:   h.cfg.CookieSecure,
    SameSite: parseSameSite(h.cfg.CookieSameSite),
}
```

---

## 🧪 Testing (Planned for v0.5.0)

### Test Structure
```
backend/
├── internal/
│   ├── handlers/
│   │   ├── auth_handler.go
│   │   └── auth_handler_test.go
│   ├── services/
│   │   ├── link_service.go
│   │   └── link_service_test.go
└── pkg/
    └── database/
        ├── memory.go
        └── memory_test.go

frontend/
├── src/
│   ├── lib/
│   │   ├── services/
│   │   │   ├── authService.ts
│   │   │   └── authService.test.ts
└── tests/
    ├── unit/
    ├── integration/
    └── e2e/
```

### Test Patterns
```go
// ✅ Correct - Focused unit test
func TestLinkService_CreateLink(t *testing.T) {
    // Arrange
    db := database.NewMemoryDB()
    service := services.NewLinkService(db, "http://localhost:8080")
    
    // Act
    response, err := service.CreateLink("https://example.com", "user123")
    
    // Assert
    assert.NoError(t, err)
    assert.NotEmpty(t, response.ShortURL)
    assert.NotEmpty(t, response.QRCode)
}
```

---

## 📝 Documentation

### Code Comments
```go
// ✅ Correct - Explanatory comment
// GenerateShortCode creates a unique short code for the given URL
// using a combination of the URL content and random bytes for uniqueness
func GenerateShortCode(originalURL string) string {
    // Implementation
}

// ❌ Incorrect - Obvious comment
// This function generates a short code
func GenerateShortCode(originalURL string) string {
    // Implementation
}
```

### README and Documentation
- **README.md**: Overview, installation, basic usage
- **ARCHITECTURE.md**: Detailed architectural decisions
- **CHANGELOG.md**: Version history
- **CODING_STANDARDS.md**: This document
- **API Documentation**: OpenAPI/Swagger (planned)

---

## 🔧 Tools and Configuration

### Go Tools
```bash
# Formatting
go fmt ./...

# Linting (planned)
golangci-lint run

# Tests
go test ./...
go test -race ./...
go test -cover ./...
```

### Frontend Tools
```bash
# Type checking
npm run check

# Linting (planned)
npm run lint

# Tests (planned)
npm test
npm run test:coverage
```

### Docker
```dockerfile
# ✅ Correct - Optimized multi-stage build
FROM golang:1.21-alpine AS builder
RUN adduser -D -s /bin/sh appuser
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/main.go

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
USER appuser
COPY --from=builder /app/main /main
EXPOSE 8080
ENTRYPOINT ["/main"]
```

---

## 📊 Quality Metrics

### Quality Objectives
- **Test Coverage**: > 80%
- **Cyclomatic Complexity**: < 10 per function
- **Function Size**: < 50 lines
- **File Size**: < 500 lines
- **Code Duplication**: < 5%

### Code Review Checklist
- [ ] Follows naming conventions
- [ ] Functions have single responsibility
- [ ] Adequate error handling
- [ ] Explicit TypeScript types
- [ ] Unit tests included
- [ ] Documentation updated
- [ ] No code duplication
- [ ] Correct security configurations

---

**Maintained by**: Danilo Monteiro  
**Based on**: Clean Code, SOLID, Go Best Practices, TypeScript Guidelines  
**Version**: 1.0 - January 2024