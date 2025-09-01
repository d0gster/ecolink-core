# 🔐 Configuração Auth0 - EcoLink

## 1. Criar Conta Auth0

1. Acesse [auth0.com](https://auth0.com) e crie uma conta gratuita
2. Crie um novo tenant (ex: `ecolink-dev`)

## 2. Configurar Aplicação SPA

1. No dashboard Auth0, vá em **Applications** > **Create Application**
2. Nome: `EcoLink Frontend`
3. Tipo: **Single Page Web Applications**
4. Tecnologia: **Vanilla JavaScript**

### Configurações da Aplicação:

**Allowed Callback URLs:**
```
http://localhost:3000/auth/callback,
https://your-domain.com/auth/callback
```

**Allowed Logout URLs:**
```
http://localhost:3000,
https://your-domain.com
```

**Allowed Web Origins:**
```
http://localhost:3000,
https://your-domain.com
```

## 3. Configurar API

1. Vá em **APIs** > **Create API**
2. Nome: `EcoLink API`
3. Identifier: `https://api.ecolink.com` (pode ser qualquer URL única)
4. Signing Algorithm: **RS256**

## 4. Configurar Provedores Sociais

### Google
1. **Connections** > **Social** > **Google**
2. Configure com suas credenciais do Google Cloud Console
3. Scopes: `openid profile email`

### GitHub
1. **Connections** > **Social** > **GitHub**
2. Configure com suas credenciais do GitHub OAuth App
3. Scopes: `user:email`

### Microsoft
1. **Connections** > **Social** > **Microsoft Account**
2. Configure com suas credenciais do Azure AD
3. Scopes: `openid profile email`

## 5. Variáveis de Ambiente

### Backend (.env)
```bash
AUTH0_DOMAIN=your-tenant.auth0.com
AUTH0_AUDIENCE=https://api.ecolink.com
```

### Frontend (.env.local)
```bash
VITE_AUTH0_DOMAIN=your-tenant.auth0.com
VITE_AUTH0_CLIENT_ID=your-spa-client-id
VITE_AUTH0_AUDIENCE=https://api.ecolink.com
```

## 6. Testar Configuração

1. Inicie o backend: `cd backend && go run cmd/main.go`
2. Inicie o frontend: `cd frontend && npm run dev`
3. Acesse `http://localhost:3000`
4. Clique em "Entrar" e teste o fluxo de autenticação

## 7. Customização (Opcional)

### Universal Login
- **Branding** > **Universal Login**
- Customize cores, logo e layout da página de login

### Rules/Actions
- Adicione metadados customizados ao token JWT
- Configure redirecionamentos condicionais

## 🔧 Troubleshooting

### Erro: "Invalid Callback URL"
- Verifique se a URL está exatamente igual nas configurações
- Certifique-se de não ter espaços ou caracteres extras

### Erro: "Access Denied"
- Verifique se a API está configurada corretamente
- Confirme se o `audience` está correto em ambos frontend e backend

### Erro: "Invalid Token"
- Verifique se o domínio Auth0 está correto
- Confirme se o algoritmo é RS256