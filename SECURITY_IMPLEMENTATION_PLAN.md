# 🔒 Security Implementation Plan - EcoLink

**Methodology**: Clean Architecture + Security by Design  
**Priority**: Complex architectural changes → Point implementations  
**Timeline**: 8 weeks (320 hours)

---

## 🏗️ Phase 1: Security Architectural Refactoring (Weeks 1-3)

### **1.1 Security Layer Architecture (40h)**

#### **Problem**: Dispersed security, no dedicated layer
#### **Solution**: Implement Security Domain Layer

```
backend/
├── internal/
│   ├── security/           # 🆕 New security layer
│   │   ├── domain/         # Security entities
│   │   │   ├── session.go
│   │   │   ├── token.go
│   │   │   └── audit.go
│   │   ├── services/       # Security services
│   │   │   ├── crypto_service.go
│   │   │   ├── session_service.go
│   │   │   ├── audit_service.go
│   │   │   └── validation_service.go
│   │   ├── middleware/     # Security middleware
│   │   │   ├── csrf.go
│   │   │   ├── rate_limit.go
│   │   │   ├── input_validation.go
│   │   │   └── security_headers.go
│   │   └── interfaces/     # Security contracts
│   │       ├── crypto.go
│   │       ├── session.go
│   │       └── audit.go
```

#### **Core Implementation**

```go
// internal/security/domain/session.go
type Session struct {
    ID          string    `json:"id"`
    UserID      string    `json:"user_id"`
    Token       string    `json:"-"` // Never serialize
    CSRFToken   string    `json:"csrf_token"`
    CreatedAt   time.Time `json:"created_at"`
    ExpiresAt   time.Time `json:"expires_at"`
    IPAddress   string    `json:"ip_address"`
    UserAgent   string    `json:"user_agent"`
    IsActive    bool      `json:"is_active"`
    LastAccess  time.Time `json:"last_access"`
}

// internal/security/interfaces/crypto.go
type CryptoService interface {
    GenerateSecureCode(input string, length int) (string, error)
    HashPassword(password string) (string, error)
    VerifyPassword(password, hash string) bool
    GenerateCSRFToken() (string, error)
    ValidateCSRFToken(token, expected string) bool
    EncryptSensitiveData(data []byte) ([]byte, error)
    DecryptSensitiveData(data []byte) ([]byte, error)
}

// internal/security/services/crypto_service.go
type cryptoService struct {
    secretKey []byte
    pepper    string
}

func (c *cryptoService) GenerateSecureCode(input string, length int) (string, error) {
    // 🔒 Replace MD5 with SHA-256 + crypto/rand
    randomBytes := make([]byte, 16)
    if _, err := rand.Read(randomBytes); err != nil {
        return "", fmt.Errorf("failed to generate random bytes: %w", err)
    }
    
    combined := append([]byte(input), randomBytes...)
    hash := sha256.Sum256(combined)
    
    return base62Encode(hash[:length]), nil
}
```

### **1.2 Session Management Architecture (32h)**

#### **Problem**: Insecure session storage (localStorage)
#### **Solution**: Secure Session Management with Redis

```go
// internal/security/services/session_service.go
type SessionService struct {
    redis      redis.Client
    crypto     CryptoService
    audit      AuditService
    config     SessionConfig
}

type SessionConfig struct {
    TokenTTL        time.Duration
    RefreshTTL      time.Duration
    MaxSessions     int
    SecureCookies   bool
    SameSite        http.SameSite
    CSRFProtection  bool
}

func (s *SessionService) CreateSession(userID, ipAddress, userAgent string) (*Session, error) {
    // Limit sessions per user
    if err := s.enforceSessionLimit(userID); err != nil {
        return nil, err
    }
    
    session := &Session{
        ID:         s.generateSessionID(),
        UserID:     userID,
        Token:      s.generateSecureToken(),
        CSRFToken:  s.crypto.GenerateCSRFToken(),
        CreatedAt:  time.Now(),
        ExpiresAt:  time.Now().Add(s.config.TokenTTL),
        IPAddress:  ipAddress,
        UserAgent:  userAgent,
        IsActive:   true,
        LastAccess: time.Now(),
    }
    
    // Store in Redis with TTL
    if err := s.storeSession(session); err != nil {
        return nil, err
    }
    
    // Audit log
    s.audit.LogSessionCreated(session.UserID, session.IPAddress)
    
    return session, nil
}
```

### **1.3 Input Validation Architecture (24h)**

#### **Problem**: Dispersed and inconsistent validation
#### **Solution**: Centralized Validation Layer

```go
// internal/security/services/validation_service.go
type ValidationService struct {
    rules map[string][]ValidationRule
}

type ValidationRule interface {
    Validate(value interface{}) error
    Name() string
}

type URLValidationRule struct{}
func (r URLValidationRule) Validate(value interface{}) error {
    url, ok := value.(string)
    if !ok {
        return errors.New("value must be string")
    }
    
    // Whitelist of allowed schemes
    allowedSchemes := []string{"http", "https"}
    parsed, err := neturl.Parse(url)
    if err != nil {
        return fmt.Errorf("invalid URL format: %w", err)
    }
    
    if !contains(allowedSchemes, parsed.Scheme) {
        return errors.New("URL scheme not allowed")
    }
    
    // Blacklist of malicious domains
    if s.isBlacklistedDomain(parsed.Host) {
        return errors.New("domain not allowed")
    }
    
    return nil
}

// Automatic validation middleware
func ValidationMiddleware(validator *ValidationService) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Validate all inputs automatically
        if err := validator.ValidateRequest(c); err != nil {
            c.JSON(400, gin.H{"error": "validation failed", "details": err.Error()})
            c.Abort()
            return
        }
        c.Next()
    }
}
```

### **1.4 Audit & Monitoring Architecture (32h)**

#### **Problem**: No security auditing
#### **Solution**: Comprehensive Security Audit System

```go
// internal/security/domain/audit.go
type AuditEvent struct {
    ID          string                 `json:"id"`
    Timestamp   time.Time             `json:"timestamp"`
    UserID      string                `json:"user_id,omitempty"`
    SessionID   string                `json:"session_id,omitempty"`
    EventType   AuditEventType        `json:"event_type"`
    Resource    string                `json:"resource"`
    Action      string                `json:"action"`
    IPAddress   string                `json:"ip_address"`
    UserAgent   string                `json:"user_agent"`
    Success     bool                  `json:"success"`
    ErrorCode   string                `json:"error_code,omitempty"`
    Metadata    map[string]interface{} `json:"metadata,omitempty"`
    RiskScore   int                   `json:"risk_score"`
}

type AuditEventType string
const (
    EventTypeAuth       AuditEventType = "authentication"
    EventTypeAuthorization AuditEventType = "authorization"
    EventTypeDataAccess AuditEventType = "data_access"
    EventTypeDataModification AuditEventType = "data_modification"
    EventTypeSecurityViolation AuditEventType = "security_violation"
)

// internal/security/services/audit_service.go
type AuditService struct {
    storage    AuditStorage
    analyzer   ThreatAnalyzer
    alerting   AlertingService
}

func (a *AuditService) LogSecurityEvent(event *AuditEvent) error {
    // Enrich event with context
    event.ID = a.generateEventID()
    event.Timestamp = time.Now()
    event.RiskScore = a.analyzer.CalculateRiskScore(event)
    
    // Store event
    if err := a.storage.Store(event); err != nil {
        return err
    }
    
    // Real-time threat analysis
    if event.RiskScore > 80 {
        a.alerting.SendSecurityAlert(event)
    }
    
    return nil
}
```

---

## 🛡️ Phase 2: Security Middleware Stack (Weeks 3-4)

### **2.1 CSRF Protection Middleware (16h)**

```go
// internal/security/middleware/csrf.go
type CSRFMiddleware struct {
    sessionService *SessionService
    config         CSRFConfig
}

type CSRFConfig struct {
    TokenLength    int
    HeaderName     string
    FormFieldName  string
    CookieName     string
    Secure         bool
    SameSite       http.SameSite
}

func (m *CSRFMiddleware) Protect() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Skip for safe methods
        if c.Request.Method == "GET" || c.Request.Method == "HEAD" || c.Request.Method == "OPTIONS" {
            c.Next()
            return
        }
        
        session, exists := c.Get("session")
        if !exists {
            c.JSON(401, gin.H{"error": "authentication required"})
            c.Abort()
            return
        }
        
        // Extract CSRF token
        token := m.extractCSRFToken(c)
        if token == "" {
            c.JSON(403, gin.H{"error": "CSRF token required"})
            c.Abort()
            return
        }
        
        // Validate token
        sess := session.(*Session)
        if !m.sessionService.ValidateCSRFToken(sess.ID, token) {
            c.JSON(403, gin.H{"error": "invalid CSRF token"})
            c.Abort()
            return
        }
        
        c.Next()
    }
}
```

### **2.2 Rate Limiting Middleware (20h)**

```go
// internal/security/middleware/rate_limit.go
type RateLimitMiddleware struct {
    redis   redis.Client
    config  RateLimitConfig
}

type RateLimitConfig struct {
    RequestsPerMinute int
    BurstSize         int
    WindowSize        time.Duration
    KeyGenerator      func(*gin.Context) string
}

func (m *RateLimitMiddleware) Limit() gin.HandlerFunc {
    return func(c *gin.Context) {
        key := m.config.KeyGenerator(c)
        
        // Implement sliding window rate limiting
        allowed, remaining, resetTime, err := m.checkRateLimit(key)
        if err != nil {
            c.JSON(500, gin.H{"error": "rate limit check failed"})
            c.Abort()
            return
        }
        
        // Informative headers
        c.Header("X-RateLimit-Limit", strconv.Itoa(m.config.RequestsPerMinute))
        c.Header("X-RateLimit-Remaining", strconv.Itoa(remaining))
        c.Header("X-RateLimit-Reset", strconv.FormatInt(resetTime.Unix(), 10))
        
        if !allowed {
            c.JSON(429, gin.H{
                "error": "rate limit exceeded",
                "retry_after": resetTime.Sub(time.Now()).Seconds(),
            })
            c.Abort()
            return
        }
        
        c.Next()
    }
}
```

### **2.3 Security Headers Middleware (12h)**

```go
// internal/security/middleware/security_headers.go
func SecurityHeadersMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // OWASP Security Headers
        c.Header("X-Content-Type-Options", "nosniff")
        c.Header("X-Frame-Options", "DENY")
        c.Header("X-XSS-Protection", "1; mode=block")
        c.Header("Referrer-Policy", "strict-origin-when-cross-origin")
        c.Header("Permissions-Policy", "geolocation=(), microphone=(), camera=()")
        
        // CSP Header
        csp := "default-src 'self'; " +
               "script-src 'self' 'unsafe-inline' https://accounts.google.com; " +
               "style-src 'self' 'unsafe-inline'; " +
               "img-src 'self' data: https:; " +
               "connect-src 'self' https://oauth2.googleapis.com; " +
               "frame-src https://accounts.google.com"
        c.Header("Content-Security-Policy", csp)
        
        // HSTS (only on HTTPS)
        if c.Request.TLS != nil {
            c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
        }
        
        c.Next()
    }
}
```

---

## 🔐 Phase 3: Cryptographic Refactoring (Weeks 4-5)

### **3.1 Secure Hash Algorithm Migration (24h)**

#### **Critical Problem**: MD5 is cryptographically broken
#### **Solution**: Migration to SHA-256 + Salt

```go
// internal/security/services/crypto_service.go
type SecureCrypto struct {
    pepper    []byte
    keyDerivation *argon2.Params
}

type argon2Params struct {
    Memory      uint32
    Iterations  uint32
    Parallelism uint8
    SaltLength  uint32
    KeyLength   uint32
}

func (c *SecureCrypto) GenerateShortCode(originalURL string) (string, error) {
    // Generate cryptographically secure salt
    salt := make([]byte, 16)
    if _, err := rand.Read(salt); err != nil {
        return "", fmt.Errorf("failed to generate salt: %w", err)
    }
    
    // Combine URL + salt + pepper
    combined := append([]byte(originalURL), salt...)
    combined = append(combined, c.pepper...)
    
    // SHA-256 hash
    hash := sha256.Sum256(combined)
    
    // Base62 encoding for URL-safe
    shortCode := base62Encode(hash[:6])
    
    // Check uniqueness in database
    if exists, err := c.checkCodeExists(shortCode); err != nil {
        return "", err
    } else if exists {
        // Retry with new salt
        return c.GenerateShortCode(originalURL)
    }
    
    return shortCode, nil
}

// Migration of existing data
func (c *SecureCrypto) MigrateExistingCodes() error {
    // 1. Find all links with MD5 codes
    // 2. Generate new secure codes
    // 3. Maintain temporary mapping
    // 4. Update gradually
    
    links, err := c.db.GetAllLinks()
    if err != nil {
        return err
    }
    
    for _, link := range links {
        newCode, err := c.GenerateShortCode(link.URL)
        if err != nil {
            continue // Log error, continue migration
        }
        
        // Keep old code for transition period
        migration := &CodeMigration{
            OldCode: link.Code,
            NewCode: newCode,
            LinkID:  link.ID,
            MigratedAt: time.Now(),
        }
        
        if err := c.db.SaveMigration(migration); err != nil {
            return err
        }
    }
    
    return nil
}
```

### **3.2 Encryption at Rest (20h)**

```go
// internal/security/services/encryption_service.go
type EncryptionService struct {
    key    []byte
    cipher cipher.AEAD
}

func NewEncryptionService(key []byte) (*EncryptionService, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }
    
    return &EncryptionService{
        key:    key,
        cipher: gcm,
    }, nil
}

func (e *EncryptionService) EncryptSensitiveData(data []byte) ([]byte, error) {
    nonce := make([]byte, e.cipher.NonceSize())
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        return nil, err
    }
    
    ciphertext := e.cipher.Seal(nonce, nonce, data, nil)
    return ciphertext, nil
}

// Apply to sensitive data
type EncryptedLink struct {
    ID              string `json:"id"`
    Code            string `json:"code"`
    EncryptedURL    []byte `json:"-"` // Encrypted at rest
    UserID          string `json:"user_id"`
    CreatedAt       time.Time `json:"created_at"`
    Clicks          int    `json:"clicks"`
}
```

---

## 🔍 Phase 4: Input Sanitization & Output Encoding (Weeks 5-6)

### **4.1 XSS Prevention Layer (20h)**

```go
// internal/security/services/sanitization_service.go
type SanitizationService struct {
    htmlPolicy *bluemonday.Policy
    urlPolicy  *bluemonday.Policy
}

func NewSanitizationService() *SanitizationService {
    // Restrictive policy for HTML
    htmlPolicy := bluemonday.StrictPolicy()
    
    // Policy for URLs
    urlPolicy := bluemonday.NewPolicy()
    urlPolicy.AllowURLSchemes("http", "https")
    
    return &SanitizationService{
        htmlPolicy: htmlPolicy,
        urlPolicy:  urlPolicy,
    }
}

func (s *SanitizationService) SanitizeURL(url string) (string, error) {
    // Validate format
    parsed, err := neturl.Parse(url)
    if err != nil {
        return "", fmt.Errorf("invalid URL format: %w", err)
    }
    
    // Whitelist of schemes
    allowedSchemes := []string{"http", "https"}
    if !contains(allowedSchemes, parsed.Scheme) {
        return "", errors.New("URL scheme not allowed")
    }
    
    // Sanitize components
    sanitized := s.urlPolicy.Sanitize(url)
    
    // Check domain blacklist
    if s.isBlacklistedDomain(parsed.Host) {
        return "", errors.New("domain not allowed")
    }
    
    return sanitized, nil
}

// Automatic sanitization middleware
func (s *SanitizationService) SanitizeMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Sanitize query parameters
        for key, values := range c.Request.URL.Query() {
            for i, value := range values {
                values[i] = s.htmlPolicy.Sanitize(value)
            }
        }
        
        // Sanitize form data
        if c.Request.Method == "POST" || c.Request.Method == "PUT" {
            c.Request.ParseForm()
            for key, values := range c.Request.PostForm {
                for i, value := range values {
                    values[i] = s.htmlPolicy.Sanitize(value)
                }
            }
        }
        
        c.Next()
    }
}
```

### **4.2 Path Traversal Prevention (16h)**

```go
// internal/security/middleware/path_validation.go
func PathValidationMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Validate path parameters
        for _, param := range c.Params {
            if err := validatePathParam(param.Key, param.Value); err != nil {
                c.JSON(400, gin.H{"error": fmt.Sprintf("invalid %s parameter", param.Key)})
                c.Abort()
                return
            }
        }
        c.Next()
    }
}

func validatePathParam(key, value string) error {
    switch key {
    case "code":
        return validateShortCode(value)
    case "id":
        return validateID(value)
    default:
        return validateGenericParam(value)
    }
}

func validateShortCode(code string) error {
    // Only alphanumeric, 6 characters
    if !regexp.MustCompile(`^[a-zA-Z0-9]{6}$`).MatchString(code) {
        return errors.New("invalid short code format")
    }
    
    // Check dangerous characters
    dangerous := []string{"..", "/", "\\", "<", ">", "&", "'", "\""}
    for _, char := range dangerous {
        if strings.Contains(code, char) {
            return errors.New("dangerous characters detected")
        }
    }
    
    return nil
}
```

---

## 📊 Phase 5: Security Testing & Validation (Weeks 6-7)

### **5.1 Security Test Suite (32h)**

```go
// tests/security/security_test.go
func TestCSRFProtection(t *testing.T) {
    router := setupTestRouter()
    
    // Test without CSRF token
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("POST", "/api/v1/links", strings.NewReader(`{"url":"https://example.com"}`))
    router.ServeHTTP(w, req)
    
    assert.Equal(t, 403, w.Code)
    assert.Contains(t, w.Body.String(), "CSRF token required")
}

func TestRateLimiting(t *testing.T) {
    router := setupTestRouter()
    
    // Make multiple requests
    for i := 0; i < 100; i++ {
        w := httptest.NewRecorder()
        req, _ := http.NewRequest("POST", "/api/v1/links", nil)
        router.ServeHTTP(w, req)
        
        if i > 60 { // Limit of 60/min
            assert.Equal(t, 429, w.Code)
        }
    }
}

func TestXSSPrevention(t *testing.T) {
    maliciousInputs := []string{
        "<script>alert('xss')</script>",
        "javascript:alert('xss')",
        "data:text/html,<script>alert('xss')</script>",
    }
    
    for _, input := range maliciousInputs {
        sanitized, err := sanitizationService.SanitizeURL(input)
        assert.Error(t, err)
        assert.NotContains(t, sanitized, "<script>")
    }
}
```

### **5.2 Penetration Testing Automation (24h)**

```go
// tests/security/pentest_automation.go
type SecurityTestSuite struct {
    baseURL string
    client  *http.Client
}

func (s *SecurityTestSuite) TestSQLInjection() {
    payloads := []string{
        "'; DROP TABLE links; --",
        "' OR '1'='1",
        "' UNION SELECT * FROM users --",
    }
    
    for _, payload := range payloads {
        resp := s.makeRequest("GET", "/api/v1/links?search="+url.QueryEscape(payload))
        assert.NotEqual(s.T(), 500, resp.StatusCode, "SQL injection vulnerability detected")
    }
}

func (s *SecurityTestSuite) TestDirectoryTraversal() {
    payloads := []string{
        "../../../etc/passwd",
        "..\\..\\..\\windows\\system32\\config\\sam",
        "%2e%2e%2f%2e%2e%2f%2e%2e%2fetc%2fpasswd",
    }
    
    for _, payload := range payloads {
        resp := s.makeRequest("GET", "/"+payload)
        assert.NotEqual(s.T(), 200, resp.StatusCode, "Directory traversal vulnerability")
    }
}
```

---

## 🚀 Phase 6: Deployment Security (Weeks 7-8)

### **6.1 Infrastructure Security (24h)**

```yaml
# docker/security/Dockerfile.secure
FROM golang:1.21-alpine AS builder

# Security: Non-root user
RUN adduser -D -s /bin/sh appuser

# Security: Minimal dependencies
RUN apk add --no-cache ca-certificates git

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/main.go

# Production image
FROM scratch

# Security: CA certificates for HTTPS
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Security: Non-root user
COPY --from=builder /etc/passwd /etc/passwd
USER appuser

COPY --from=builder /app/main /main

# Security: Non-privileged port
EXPOSE 8080

ENTRYPOINT ["/main"]
```

### **6.2 Security Monitoring (20h)**

```go
// internal/security/monitoring/security_monitor.go
type SecurityMonitor struct {
    metrics    prometheus.Registerer
    alerting   AlertManager
    thresholds SecurityThresholds
}

type SecurityThresholds struct {
    FailedAuthAttempts int
    RateLimitViolations int
    SuspiciousPatterns  []string
}

func (m *SecurityMonitor) MonitorSecurityEvents() {
    // Security metrics
    failedAuthCounter := prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "security_failed_auth_total",
            Help: "Total failed authentication attempts",
        },
        []string{"ip", "user_agent"},
    )
    
    rateLimitCounter := prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "security_rate_limit_violations_total",
            Help: "Total rate limit violations",
        },
        []string{"ip", "endpoint"},
    )
    
    m.metrics.MustRegister(failedAuthCounter, rateLimitCounter)
}
```

---

## 📈 Timeline and Resources

### **Detailed Schedule**

| Week | Phase | Hours | Deliverables |
|------|-------|-------|-------------|
| 1-2 | Security Architecture | 80h | Security Layer, Session Management |
| 3 | Middleware Stack | 40h | CSRF, Rate Limiting, Headers |
| 4-5 | Crypto Refactoring | 44h | SHA-256 Migration, Encryption |
| 5-6 | Input/Output Security | 36h | XSS Prevention, Path Validation |
| 6-7 | Security Testing | 56h | Test Suite, Penetration Testing |
| 7-8 | Deployment Security | 44h | Infrastructure, Monitoring |

**Total**: 300 hours (37.5 business days)

### **Required Resources**

- **1 Security Architect** (Senior): 120h
- **2 Backend Developers** (Mid/Senior): 120h
- **1 DevOps Engineer**: 40h
- **1 Security Tester**: 20h

### **Success Criteria**

- ✅ Zero Critical/High vulnerabilities
- ✅ 100% security test coverage
- ✅ OWASP Top 10 compliance
- ✅ Performance impact < 10%
- ✅ Security audit approved

This plan prioritizes the most complex architectural changes that will have the greatest impact on security and project structure, following Clean Architecture and Security by Design principles.