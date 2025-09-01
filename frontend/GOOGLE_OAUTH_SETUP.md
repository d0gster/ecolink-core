# 🔐 Configuração Google OAuth - EcoLink

## 📋 Passos para Configurar

### 1. **Google Cloud Console**
1. Acesse: https://console.cloud.google.com/
2. Crie um novo projeto ou selecione existente
3. Vá em **APIs & Services** > **Credentials**

### 2. **Configurar OAuth 2.0**
1. Clique em **+ CREATE CREDENTIALS** > **OAuth 2.0 Client IDs**
2. Escolha **Web application**
3. Configure:
   - **Name**: EcoLink Local
   - **Authorized JavaScript origins**: 
     - `http://localhost:5173`
   - **Authorized redirect URIs**:
     - `http://localhost:5173/auth/callback`

### 3. **Copiar Credenciais**
1. Copie **Client ID** e **Client Secret**
2. Cole no arquivo `.env.local`:

```env
GOOGLE_CLIENT_ID=seu_client_id_aqui
GOOGLE_CLIENT_SECRET=seu_client_secret_aqui
AUTH_SECRET=uma_chave_secreta_aleatoria_aqui
```

### 4. **Gerar AUTH_SECRET**
```bash
openssl rand -base64 32
```

### 5. **Reiniciar Servidor**
```bash
npm run dev
```

## ✅ Teste
1. Acesse http://localhost:5173
2. Tente encurtar um link
3. Clique em "Continuar com Google"
4. Faça login com sua conta Google

## 🔧 Troubleshooting
- Verifique se as URLs estão corretas no Google Console
- Confirme se as variáveis de ambiente estão carregadas
- Reinicie o servidor após alterar `.env.local`