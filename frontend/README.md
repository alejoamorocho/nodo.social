# NODO.SOCIAL Frontend

## Visión General
Frontend para la plataforma NODO.SOCIAL, construido con Next.js 13+ y TypeScript, siguiendo los principios de Clean Architecture y Domain-Driven Design.

## Arquitectura

### 1. Capas de la Aplicación

#### 1.1 Dominio (`src/domain/`)
- Modelos y entidades del negocio
- Interfaces y tipos base
- Reglas de negocio puras
- No depende de ninguna otra capa

#### 1.2 Aplicación (`src/application/`)
- Casos de uso
- Servicios de aplicación
- Estado global (store)
- Depende solo del dominio

#### 1.3 Infraestructura (`src/infrastructure/`)
- Implementaciones concretas
- Adaptadores de API
- Servicios externos
- Configuraciones
- Depende de dominio y aplicación

#### 1.4 Interfaz (`src/`)
- `app/`: Páginas y rutas (Next.js App Router)
- `components/`: Componentes de UI reutilizables
- `features/`: Módulos específicos de funcionalidad
- Depende de todas las capas anteriores

### 2. Estructura de Carpetas
```
src/
├── app/                    # Rutas y páginas (Next.js App Router)
│   ├── (auth)/            # Grupo de rutas autenticadas
│   └── (main)/            # Grupo de rutas principales
├── components/            # Componentes compartidos
│   ├── ui/               # Componentes base (botones, inputs, etc.)
│   └── layout/           # Componentes de estructura
├── features/             # Módulos de funcionalidad
│   ├── nodes/           # Feature de nodos
│   ├── fundraising/     # Feature de recaudación
│   └── profile/         # Feature de perfil
├── domain/              # Capa de dominio
│   ├── models/         # Entidades y modelos
│   └── interfaces/     # Contratos e interfaces
├── application/        # Capa de aplicación
│   ├── services/      # Servicios de aplicación
│   ├── store/         # Estado global
│   └── hooks/         # Hooks personalizados
└── infrastructure/    # Capa de infraestructura
    ├── api/          # Cliente API y adaptadores
    ├── config/       # Configuraciones
    └── utils/        # Utilidades generales
```

### 3. Principios y Buenas Prácticas

#### 3.1 Clean Code
- Nombres descriptivos y significativos
- Funciones pequeñas con responsabilidad única
- Máximo 3 parámetros por función
- Evitar comentarios obvios, código autodocumentado
- Tests unitarios para lógica de negocio

#### 3.2 Convenciones de Nombrado
- Componentes: PascalCase (`NodeCard.tsx`)
- Hooks: camelCase con prefix 'use' (`useNodeData.ts`)
- Servicios: PascalCase con suffix 'Service' (`NodeService.ts`)
- Interfaces: PascalCase (`Node`, no usar prefix 'I')
- Features: kebab-case para carpetas (`feature-name/`)

#### 3.3 Estado y Efectos Secundarios
- Hooks personalizados para lógica reutilizable
- Zustand para estado global simple
- React Query para estado del servidor
- Efectos secundarios aislados en servicios

## Configuración del Proyecto

### Requisitos
- Node.js 18+
- npm 9+

### Instalación
```bash
npm install
```

### Desarrollo
```bash
npm run dev
```

### Construcción
```bash
npm run build
```

### Tests
```bash
npm run test
```

## Estándares de Código
- ESLint para linting
- Prettier para formateo
- Husky para git hooks
- Commitlint para mensajes de commit
