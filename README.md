# 🏢 Sistema de Gerenciamento de Estabelecimentos

Uma aplicação fullstack moderna construída com Nuxt 3 (frontend) e Golang (backend) para gerenciamento eficiente de estabelecimentos comerciais e suas lojas.

![Preview da Aplicação](https://via.placeholder.com/800x400?text=Dashboard+Preview) <!-- Adicione uma imagem real posteriormente -->

## ✨ Funcionalidades Principais

### 🚀 Sistema de Autenticação

- Login simulado com controle de acesso
- Gerenciamento de sessão do usuário

### 📊 Gestão Completa

- Cadastro e edição de estabelecimentos comerciais
- Controle de múltiplas lojas por estabelecimento
- Dashboard analítico com visão geral

### � UX Avançado

- Interface responsiva (mobile, tablet, desktop)
- Notificações interativas com SweetAlert2
- Navegação intuitiva

## 🛠 Stack Tecnológica

### Frontend

- [Nuxt 3](https://nuxt.com) - Framework Vue.js para SSR
- [Vite](https://vitejs.dev/) - Build tool ultrarrápido
- [Tailwind CSS](https://tailwindcss.com) - Framework CSS utility-first
- [SweetAlert2](https://sweetalert2.github.io) - Alertas modais elegantes

### Backend

- [Golang](https://go.dev) - Linguagem de programação eficiente
- [Echo](https://echo.labstack.com) - Framework web (se aplicável)
- Docker - Containerização do serviço

## 🚀 Instalação e Execução

### Pré-requisitos

- Node.js 18+
- Go 1.20+ (apenas para desenvolvimento backend)
- Docker e Docker Compose
- npm ou yarn

### 1. Clone o repositório

```bash
git clone https://github.com/gabryel-araujo/teste-snet.git
cd teste-snet
```

### 2. Instale as dependências

```bash
npm install
npm run dev
```

O projeto estará disponível em http://localhost:3000

### 3. Inicie o back-end

Abra uma nova aba do seu terminal e acesse a pasta do back-end

```bash
cd back-go-snet
docker compose up --build
```

O docker automaticamente criará os scripts do banco de dados, mas um arquivo de referencia estará disponivel na pasta /back-go-snet/db/script.sql

### Estrutura do projeto

```
teste-snet/
├── .nuxt/                  # Arquivos gerados automaticamente pelo Nuxt
├── .output/                # Build de produção gerado pelo Nuxt
├── api/                    # Arquivo de Configuração da api
├── assets/                 # Arquivos estáticos do css
│   ├── css/                # Estilos globais
│
├── back-go-snet/           # Aplicação backend em Golang
│   ├── db/                 # Arquivo de configuração do banco de dados
│   ├── handler/            # Lógica dos endpoints
│   ├── models/             # Definições de estruturas de dados
│   ├── repository/         # Consultas SQL
│   ├── service/            # Utilização da regra de negócio
│   ├── docker-compose.yml  # Script do docker para iniciar as dependencias
│   ├── Dockerfile          # Script de configuração do docker
│   ├── go.mod              # Dependências do Go
│   ├── go.sum              #
│   └── main.go             # Ponto de entrada da aplicação
│
├── components/             # Componentes Vue reutilizáveis
│
├── layouts/                # Layouts da aplicação
│   ├── default.vue         # Layout padrão
│   └── login.vue           # Layout para páginas de autenticação
│
├── middleware/             # Middlewares de autenticação/roteamento
│   └── fake-auth.ts        # Middleware fake de autenticação
│
├── node_modules/           # Dependências do Node.js
├── pages/                  # Páginas/rotas da aplicação
│   ├── app/                # Páginas do dashboard
│   │   ├── index.vue       # Dashboard principal
│   │   ├── sobre.vue       # Página da documentação html
│   │
│   └── index.vue           # Página inicial
│
├── public/                 # Arquivos públicos (acessíveis diretamente)
│
├── schemas/                # Esquemas de validação
│   ├── estabSchema.ts      # Validação de estabelecimentos
│   ├── loginSchema.ts      # Validação de estabelecimentos
│   └── storeSchema.ts      # Validação de lojas
│
├── server/                 # Configurações do servidor (Nuxt)
│
├── types/                  # Tipos TypeScript
│
├── .gitignore              # Arquivos ignorados pelo Git
├── app.vue                 # Componente raiz da aplicação
├── nuxt.config.ts          # Configuração do Nuxt
├── package-lock.json       # Versões exatas das dependências
├── package.json            # Configuração do projeto e scripts
└── README.md               # Documentação do projeto
```
