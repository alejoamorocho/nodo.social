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

```markdown
nodo.social/
├── .firebaserc             # Firebase project configuration
├── config/                 # Configuration files
├── firebase.json          # Firebase service configuration
├── firestore.indexes.json # Firestore indexes configuration
├── firestore.rules        # Firestore security rules
├── frontend/              # Frontend application
├── functions/             # Backend services (Go)
│   ├── domain/           # Business rules and entities
│   │   ├── models/       # Domain models (User, Node, Feed)
│   │   ├── dto/         # Data Transfer Objects
│   │   └── repositories/ # Repository interfaces
│   ├── infrastructure/   # External services implementation
│   │   └── firebase/    # Firebase implementations
│   ├── interfaces/       # Interface adapters
│   │   ├── http/        # HTTP handlers and routes
│   │   └── cloud/       # Cloud Functions
│   ├── services/        # Application business logic
│   │   └── user/        # User-related services
│   └── go.mod           # Go module definition
├── package.json          # Node.js project configuration
├── storage.rules         # Storage security rules
└── update_imports.ps1    # Import update script
```

#### Architecture Overview

The project follows a Clean Architecture pattern with distinct layers:

### Backend Architecture (Go)

1. **Domain Layer** (`functions/domain/`)
   - Contains core business entities and logic
   - Defines repository interfaces
   - Independent of external frameworks
   - Houses business rules and validation
   - DTOs for data transformation and validation
   - Models: User, Node, Feed, etc.

2. **Services Layer** (`functions/services/`)
   - Implements core business logic
   - Orchestrates domain objects
   - Handles use case implementation
   - Maintains business rule integrity
   - User service for authentication and profile management
   - Node service for content management

3. **Interface Layer** (`functions/interfaces/`)
   - HTTP Handlers: REST API endpoints
   - Base handler with common functionality
   - User handler for profile operations
   - Node handler for content operations
   - Cloud Functions: Firebase function handlers
   - Adapts external requests to internal services
   - Handles request/response transformations

4. **Infrastructure Layer**
   - Firebase integration
   - Database implementations
   - External service adapters
   - Configuration management

### Cloud Infrastructure

The project utilizes Firebase services for its cloud infrastructure:

- **Firebase Auth**: User authentication and authorization
- **Cloud Firestore**: NoSQL database for data storage
- **Cloud Storage**: File and media storage
- **Cloud Functions**: Serverless compute platform
- **Firebase Hosting**: Web application hosting and CDN

### Data Flow

1. External requests come through HTTP endpoints or Firebase triggers
2. Interface layer adapters process and validate requests
3. Services layer executes business logic
4. Domain layer enforces business rules
5. Infrastructure layer handles external service communication

### Frontend Technologies

- Framework: Next.js
- Language: TypeScript
- Styling: TailwindCSS
- State Management: React Context/Redux

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

2. Install backend dependencies:

```bash
cd functions
go mod tidy
go mod vendor  # Para asegurar consistencia en las dependencias
```

3. Install frontend dependencies:

```bash
cd frontend
npm install
```

4. Configure Firebase:

```bash
npm install -g firebase-tools
firebase login
firebase init
```

5. Set up environment variables:

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

2. Start Firebase emulators (opcional):

```bash
firebase emulators:start
```

3. Start the frontend:

```bash
cd frontend
npm run dev
```

4. Run tests:

```bash
cd functions
go test ./...
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
