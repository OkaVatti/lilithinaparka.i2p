# Lilith Parker's I2P Blog - Frontend

A minimal, terminal-inspired blog frontend built with Nuxt 4, Deno, and TailwindCSS v4.

## Tech Stack

- **Framework**: Nuxt 4
- **Runtime**: Deno
- **State Management**: Pinia
- **Styling**: TailwindCSS v4
- **Language**: TypeScript
- **Markdown**: marked

## Features

- Dynamic markdown rendering from SQLite backend
- Real-time blog post updates
- BlueSky social feed integration
- Profile page with static and dynamic info
- Terminal/retro aesthetic inspired by classic hacker sites
- Responsive design
- SEO-friendly

## Setup

1. Install dependencies:
```bash
npm install
```

2. Configure environment variables:
Create a `.env` file with:
```
NUXT_PUBLIC_API_BASE=http://localhost:8080/api
```

3. Run development server:
```bash
npm run dev
```

4. Build for production:
```bash
npm run build
```

5. Preview production build:
```bash
npm run preview
```

## Project Structure
```
frontend/
├── assets/
│   └── css/
│       └── main.css          # Global styles
├── components/
│   ├── Header.vue            # Site header
│   ├── Footer.vue            # Site footer
│   ├── BlogPostCard.vue      # Blog post preview
│   └── BskyPostCard.vue      # BlueSky post card
├── composables/
│   └── useMarkdown.ts        # Markdown rendering
├── layouts/
│   └── default.vue           # Default layout
├── pages/
│   ├── index.vue             # Home page
│   ├── blog/
│   │   ├── index.vue         # Blog list
│   │   └── [slug].vue        # Individual blog post
│   ├── social.vue            # BlueSky feed
│   └── profile.vue           # Profile page
├── stores/
│   ├── blog.ts               # Blog state management
│   ├── bsky.ts               # BlueSky state management
│   └── profile.ts            # Profile state management
├── types/
│   └── index.ts              # TypeScript types
├── app.vue                   # Root component
├── nuxt.config.ts            # Nuxt configuration
└── tailwind.config.js        # Tailwind configuration
```

## API Endpoints

The frontend expects the following backend endpoints:

- `GET /api/blog/posts` - Get all blog posts
- `GET /api/blog/posts/:slug` - Get post by slug
- `GET /api/blog/category/:category` - Get posts by category
- `GET /api/blog/tag/:tag` - Get posts by tag
- `GET /api/bsky/posts` - Get BlueSky posts
- `GET /api/profile` - Get profile information

## Styling

The site uses a terminal/retro aesthetic with:
- Monospace fonts (Courier New)
- Green accent color (#00ff00)
- Dark background (#0a0a0a)
- Minimal borders and boxes
- ASCII art elements

## Deployment

For I2P deployment:
1. Build the production bundle: `npm run build`
2. The `.output` directory contains the production build
3. Configure your I2P tunnel to point to the server
4. Ensure the backend API is accessible

## License

MIT