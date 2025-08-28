# 🌱 EcoLink - Core

> **"Encurtamento de links sustentável e de alta performance"**

EcoLink é um sistema de encurtamento de links de alta performance, construído com foco em eficiência, escalabilidade e sustentabilidade digital.

## 🚀 Tech Stack

### Backend
- **Go (Golang)** com Gin Framework
- **Cloud Firestore** para mapeamento de URLs
- **Cloud SQL PostgreSQL** para dados de usuários
- Geração de QR Codes nativa

### Frontend  
- **SvelteKit** com SSR/SSG
- **TailwindCSS** para styling
- **TypeScript** para type safety

### Infraestrutura
- **Google Cloud Platform (GCP)**
- **Cloud Run** para escalabilidade serverless
- **Auth0** para autenticação multi-plataforma
- **Docker** para containerização

## 🏗️ Arquitetura

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   SvelteKit     │    │   Go Backend    │    │   Firestore     │
│   Frontend      │◄──►│   (Gin API)     │◄──►│   Database      │
│                 │    │                 │    │                 │
└─────────────────┘    └─────────────────┘    └─────────────────┘
         │                       │                        │
         │              ┌─────────────────┐               │
         └──────────────►│     Auth0       │◄─────────────┘
                        │  Authentication │
                        └─────────────────┘
```

## 🛠️ Como Rodar Localmente

### Pré-requisitos
- Go 1.21+
- Node.js 18+
- Docker (opcional)

### Backend
```bash
cd backend
go mod tidy
go run main.go
```

### Frontend
```bash
cd frontend
npm install
npm run dev
```

### Docker (Ambiente Completo)
```bash
docker-compose up --build
```

## ✨ Features Implementadas

- ✅ Encurtamento de URLs com hash único
- ✅ Geração de QR Codes
- ✅ Redirecionamento rápido
- ✅ Dashboard básico
- ✅ Autenticação multi-plataforma
- ✅ Histórico de links

## 🔮 Features da Versão Completa (Privada)

A versão completa e proprietária deste projeto, atualmente em desenvolvimento privado, expande o core para incluir um ecossistema de features inovadoras:

### 🌍 Eco-Analytics Dashboard
- Cálculo da redução da pegada de carbono digital
- Métricas de eficiência energética por redirecionamento
- Relatórios de impacto ambiental personalizados

### 🤖 Motor de IA Preditiva
- Otimização automática de links baseada em padrões de uso
- Sugestão inteligente de UTMs para campanhas
- Análise preditiva de performance de links

### 🎨 QR Codes Dinâmicos Premium
- Templates sustentáveis com design eco-friendly
- Customização avançada com branding corporativo
- QR Codes que se adaptam ao contexto do usuário

### 🌐 Roteamento Geográfico Inteligente
- Redirecionamento baseado na localização do usuário
- Otimização de latência para minimizar consumo energético
- CDN verde com servidores alimentados por energia renovável

### 📊 Analytics Avançados
- Heatmaps de cliques em tempo real
- Análise de comportamento do usuário
- Integração com Google Analytics 4 e ferramentas de BI

### 🔐 Segurança Enterprise
- Proteção contra links maliciosos com IA
- Auditoria completa de acessos
- Compliance com LGPD/GDPR

## 🎯 Roadmap Público

- [ ] **v1.1**: Temas dark/light mode
- [ ] **v1.2**: API pública com rate limiting
- [ ] **v1.3**: Integração com Slack/Discord
- [ ] **v2.0**: PWA com funcionalidades offline

## 🤝 Contribuição

Este é um projeto de portfólio demonstrativo. Para colaborações ou interesse na versão completa, entre em contato.

## 📄 Licença

MIT License - veja [LICENSE](LICENSE) para detalhes.

---

**Desenvolvido com 💚 por Danilo Monteiro**
  
*Parte da série "Vibecoding Chronicles: A Jornada de um Arquiteto Full Stack"*