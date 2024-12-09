# Nodo Social

A Crowdfunding Social Platform and Collaborative Marketplace connecting people working on social, environmental, and animal causes with individuals and businesses who want to support them.

## Project Structure

```
nodo.social/
├── frontend/               # Next.js frontend application
│   ├── src/
│   │   ├── app/           # Next.js app directory
│   │   ├── components/    # Reusable components
│   │   ├── lib/          # Utilities and helpers
│   │   └── types/        # TypeScript type definitions
│   └── public/           # Static files
├── functions/             # Firebase Cloud Functions (Go)
│   ├── cmd/              # Command line tools
│   ├── internal/         # Internal packages
│   └── pkg/              # Public packages
└── config/               # Configuration files
```

## Technologies

- Frontend: Next.js with TypeScript
- Backend: Firebase Cloud Functions with Go
- Database: Firestore
- Authentication: Firebase Auth
- Storage: Firebase Storage
- Styling: TailwindCSS with Dracula theme colors

## Configuración del Entorno

### Frontend (Next.js)
1. Copia el archivo de configuración de entorno:
   ```bash
   cp frontend/.env.template frontend/.env.local
   ```
2. Actualiza las variables en `frontend/.env.local` con tus credenciales de Firebase:
   ```env
   NEXT_PUBLIC_FIREBASE_API_KEY=tu-api-key
   NEXT_PUBLIC_FIREBASE_AUTH_DOMAIN=tu-auth-domain
   NEXT_PUBLIC_FIREBASE_PROJECT_ID=tu-project-id
   ...
   ```

### Backend (Firebase Functions con Go)
1. Copia el archivo de configuración de entorno:
   ```bash
   cp functions/.env.template functions/.env
   ```
2. Actualiza las variables en `functions/.env` con tus credenciales de servicio de Firebase:
   ```env
   FIREBASE_PROJECT_ID=tu-project-id
   FIREBASE_PRIVATE_KEY_ID=tu-private-key-id
   FIREBASE_PRIVATE_KEY=tu-private-key
   ...
   ```

⚠️ **IMPORTANTE**: Nunca subas los archivos `.env` al control de versiones. Estos archivos contienen información sensible y ya están incluidos en `.gitignore`.

## Setup Instructions

1. Install dependencies:
   ```bash
   npm install
   ```

2. Set up Firebase credentials:
   - Copy `config/serviceAccountKey.template.json` to `config/serviceAccountKey.json`
   - Fill in your Firebase service account credentials

3. Start development server:
   ```bash
   npm run dev
   ```

## Clean Architecture

This project follows Clean Architecture principles:

- Domain Layer: Core business logic
- Application Layer: Use cases and application logic
- Interface Layer: Controllers and presenters
- Infrastructure Layer: External services and frameworks

## Contributing

Please follow the clean code principles and architecture guidelines when contributing to this project.
