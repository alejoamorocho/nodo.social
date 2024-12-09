# NODO.SOCIAL - Social Impact Platform

A platform designed to connect people and entrepreneurs with social, environmental, and animal causes through thematic nodes.

## Project Overview

Nodo.social is a platform that facilitates meaningful connections between social initiatives and supporters, leveraging technology to maximize social impact.

## Core Features

### Node Management

- Publication of causes (social/environmental/animal)
- Update system
- Category-based feed
- Product approval system
- Impact metrics

### User Connections

- Node following
- User following
- Personalized feed
- Activity notifications

### Store Management

- Product catalog
- Node linking
- Support percentages
- Direct contact (WhatsApp)

### User Engagement

- Participation achievements
- Impact ranking
- Recognition badges

## Technical Stack

### Backend (Go/Firebase)

- Language: Go 1.20
- Architecture: Clean Architecture
- Infrastructure: Firebase Cloud Functions

#### Project Structure

```go
functions/
├── cmd/                    # Application entry points
├── domain/                # Business rules and entities
│   ├── models/            # Domain models
│   └── repositories/      # Repository interfaces
├── dto/                   # Data Transfer Objects
├── interfaces/            # Interface adapters
│   ├── http/             # HTTP Handlers
│   └── cloud/            # Cloud Functions
├── internal/              # Shared internal code
│   └── config/           # Centralized configuration
└── services/             # Business logic
```

#### Architecture Overview

1. **Domain Layer** (`domain/`)

   - Contains business entities and repository interfaces
   - Independent of external frameworks

1. **Service Layer** (`services/`)

   - Implements business logic
   - Uses repository interfaces from domain layer

1. **Interface Layer** (`interfaces/`)

   - HTTP handlers and Cloud Functions
   - Adapts external requests to internal services

1. **Infrastructure Layer** (`internal/`)

   - Configuration and external service implementations
   - Firebase and database implementations

### Frontend Technologies

- Framework: Next.js
- Language: TypeScript
- Styling: TailwindCSS
- State Management: React Context/Redux

### Cloud Infrastructure

- Authentication: Firebase Auth
- Database: Cloud Firestore
- Storage: Cloud Storage
- Functions: Cloud Functions
- CDN: Firebase Hosting

## Getting Started

### Prerequisites

- Go 1.20 or higher
- Node.js 16 or higher
- Firebase CLI
- Google Cloud SDK

### Installation

1. Clone the repository:

```bash
git clone https://github.com/kha0sys/nodo.social.git
cd nodo.social
```

1. Install backend dependencies:

```bash
cd functions
go mod tidy
```

1. Install frontend dependencies:

```bash
cd frontend
npm install
```

1. Set up environment variables:

```bash
cp .env.example .env
# Edit .env with your configuration
```

### Running Locally

1. Start the backend:

```bash
cd functions
go run cmd/main.go
```

1. Start the frontend:

```bash
cd frontend
npm run dev
```

## Development Guidelines

### Code Organization

- Follow Clean Architecture principles
- Use DTOs for data transfer
- Implement interfaces for external dependencies
- Keep business logic in services

### Development Standards

- Write unit tests for business logic
- Document public interfaces
- Use dependency injection
- Follow Go and TypeScript best practices

### Version Control

1. Create feature branches from `main`
1. Use conventional commits
1. Submit PRs for review
1. Merge to `main` after approval

## Testing

### Backend Testing

```bash
cd functions
go test ./...
```

### Frontend Testing

```bash
cd frontend
npm test
```

## Deployment

### Backend Deployment

```bash
cd functions
firebase deploy --only functions
```

### Frontend Deployment

```bash
cd frontend
npm run build
firebase deploy --only hosting
```

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
