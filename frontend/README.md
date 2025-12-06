# lilithinaparka.i2p - Interactive Portal Frontend

## Overview

This is the frontend application for a privacy-focused, TUI-inspired portal and blog built with Nuxt 4, powered by Deno, and accessible exclusively via the I2P network
It delivers a unique retro terminal aesthetic while integrating a dynamic blog, interactive game suite, art gallery, and social feed. The project emphasizes security, privacy, and a highly customizable user experience.

## ✨ Core Features
- Dynamic Content Rendering: Server-side rendering for SEO and performance, with client-side hydration for interactive features.
- Comprehensive Theming System: 12+ color schemes (Dracula, Nord, Gruvbox, etc.) and 4 seasonal themes with animated ASCII backgrounds.
- Interactive Game Suite: 8+ browser-based games (Bones, Crypt, Gems, etc.) with leaderboards, local save states, and I2P-based multiplayer where applicable.
- Integrated Social Feed: Seamless display of BlueSky posts fetched and cached via the backend API.
- Advanced Media Gallery: Masonry-style art gallery with EXIF data and adaptive video streaming for I2P constraints.
- Privacy-First Design: No tracking, anonymized analytics, and a strict Content Security Policy (CSP).

## 🛠️ Tech Stack

- Framework: Nuxt 4 (Vue 3)
- Runtime: Deno 2.5
- Package Manager: Deno (handles npm packages)
- Styling: TailwindCSS v4 via @tailwindcss/vite
- State Management: Pinia
- Language: TypeScript
- Markdown: marked with syntax highlighting

## 🚀 Getting Started

### Prerequisites
- Deno: Version 2.5.6 or higher.
- I2P Router (for testing): For local development, you may connect to a backend running on localhost.

### Installation & Setup

#### Clone and Navigate:
```bash

git clone <your-repository-url>
cd frontend
```

#### Install Dependencies with Deno:
```bash
# Deno will read the project's package.json and install dependencies[citation:1]
deno install
```

### Configure Environment:

#### Create a .env file in the project root:
```env
# API Configuration
NUXT_PUBLIC_API_BASE=http://localhost:8080/api
# Site Identity
NUXT_PUBLIC_SITE_NAME="lilithinaparka.i2p"
NUXT_PUBLIC_BSKY_HANDLE="@lilithinaparka.bsky.social"
# Default Theme
NUXT_PUBLIC_DEFAULT_THEME="dracula"
NUXT_PUBLIC_DEFAULT_COLORSCHEME="dark"
```

### Development

#### Start the development server with hot-reload:
```bash
deno run dev
```

The application will be available at http://localhost:3000.

### Building for Production

#### Build the optimized static site and server-side rendering bundle:
```bash
deno run build
```

#### The output will be in the .output directory. You can preview the production build locally:
```bash
deno run preview
```

## 📁 Project Structure

### A detailed overview of key directories to help you navigate the codebase:
```text
frontend/
├── app/                          # Main Nuxt application source
│   ├── assets/css/               # Global & modular CSS[citation:3]
│   ├── components/               # Vue components (Art, Blog, Games, etc.)
│   ├── composables/              # Vue composables (useApi, useTheme, etc.)
│   ├── layouts/                  # Layout components
│   └── pages/                    # File-based routing
├── stores/                       # Pinia stores for state management
├── types/                        # TypeScript type definitions
├── nuxt.config.ts                # Nuxt configuration[citation:9]
├── tailwind.config.js            # TailwindCSS design tokens[citation:3]
├── deno.json                     # Deno runtime configuration[citation:1]
└── app.vue                       # Root application component
```

## 🎨 Styling & Theming Guide

The application uses a Utility-First approach with TailwindCSS v4, extended with a custom design system

- Configuration: Define brand colors, fonts, and spacing in tailwind.config.js as semantic design tokens (e.g., primary-500 instead of #8b9fef).
- CSS Structure: Global styles are imported via apps/assets/css/main.css. Component-specific styles are co-located.
- Theme Implementation: Theming is managed by the useTheme composable and stores, which swap CSS classes and variables. Seasonal themes inject animated ASCII art backgrounds.
- Performance: Tailwind's Just-in-Time (JIT) compiler purges unused CSS automatically


## 📡 API Integration

The frontend communicates with the Go backend via RESTful endpoints. The useApi composable provides a centralized, typed client with error handling.

Example Blog API Call:
```typescript

// composables/useApi.ts
const { data: posts, refresh } = await useFetch('/api/blog/posts', {
  baseURL: runtimeConfig.public.apiBase,
  // ...options
});
```

## 🔒 Security Notes

    Content Security Policy (CSP): Strict headers are configured in nuxt.config.ts.

    Input Sanitization: All user-generated content rendered via v-html is sanitized with DOMPurify.

    Dependency Auditing: Regular deno audit runs are recommended to check for vulnerabilities.

## 🐳 Deployment (Docker)

A Dockerfile is provided for containerized deployment.

#### Build the image:
```bash
docker build -t lilithinaparka-frontend:latest .
```

#### Run the container:
```bash
docker run -p 3000:3000 --env-file .env lilithinaparka-frontend:latest
```

#### 🔍 Troubleshooting

- "Permission Denied" errors (Deno): Use the appropriate --allow- flags (e.g., --allow-net, --allow-read) when running scripts. Permissions are explicitly defined in deno.json
- Styles not loading: Ensure @tailwindcss/vite is correctly added as a Vite plugin in nuxt.config.ts and that main.css is imported
- API connection failed: Verify the backend server is running and the NUXT_PUBLIC_API_BASE URL in your .env file is correct.

