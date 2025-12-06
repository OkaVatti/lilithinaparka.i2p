<!-- app/pages/index.vue -->
<template>
  <div class="home-page">
    <!-- Hero Section -->
    <section class="hero-section mb-12">
      <div class="text-center">
        <h1 class="text-4xl md:text-6xl font-bold text-accent mb-4 font-mono">
          lilthinaparka.i2p
        </h1>
        <p class="text-xl text-text-secondary mb-8 font-mono">
          > personal blog-site + art gallery + portfolio + whatever i need it to be
        </p>
        
        <div class="inline-block p-4 border border-accent/20 rounded-lg bg-bg-secondary/50">
          <pre class="text-xs md:text-sm text-text-secondary font-mono">
$ whoami
name: Lilith Parker
username: @lilithinaparka
age: 20
bio: "chronically online cat girl who likes programming, music, and burritos"
location: "Earth, Sol System, Milky Way Galaxy"
          </pre>
        </div>
      </div>
    </section>
    
    <!-- Quick Stats -->
    <div class="grid grid-cols-2 md:grid-cols-4 gap-4 mb-12">
      <div class="stat-card">
        <div class="stat-number">50+</div>
        <div class="stat-label font-mono">Blog Posts</div>
      </div>
      <div class="stat-card">
        <div class="stat-number">100+</div>
        <div class="stat-label font-mono">Art Pieces</div>
      </div>
      <div class="stat-card">
        <div class="stat-number">10+</div>
        <div class="stat-label font-mono">Projects</div>
      </div>
      <div class="stat-card">
        <div class="stat-number">5+</div>
        <div class="stat-label font-mono">Games</div>
      </div>
    </div>
    
    <!-- Recent Content Sections -->
    <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
      <!-- Recent Blog Posts -->
      <section class="content-section">
        <h2 class="section-title">
          <span class="text-accent">$</span> recent blog posts
        </h2>
        <div class="space-y-4">
          <BlogPostCard
            v-for="post in recentPosts"
            :key="post.id"
            :post="post"
            class="hover:translate-x-2 transition-transform duration-200"
          />
        </div>
        <NuxtLink to="/blog" class="section-link">
          view all posts →
        </NuxtLink>
      </section>
      
      <!-- Recent Art -->
      <section class="content-section">
        <h2 class="section-title">
          <span class="text-accent">$</span> recent art
        </h2>
        <div class="grid grid-cols-2 gap-3">
          <ArtCard
            v-for="art in recentArt"
            :key="art.id"
            :art="art"
            class="hover:scale-105 transition-transform duration-200"
          />
        </div>
        <NuxtLink to="/gallery" class="section-link">
          browse gallery →
        </NuxtLink>
      </section>
      
      <!-- Recent Bluesky Posts -->
      <section class="content-section">
        <h2 class="section-title">
          <span class="text-accent">$</span> recent bluesky
        </h2>
        <div class="space-y-4">
          <BskyPostCard
            v-for="post in recentBsky"
            :key="post.id"
            :post="post"
          />
        </div>
        <NuxtLink to="/social" class="section-link">
          view social →
        </NuxtLink>
      </section>
    </div>
    
    <!-- Featured Projects -->
    <section class="mt-12 content-section">
      <h2 class="section-title">
        <span class="text-accent">$</span> featured projects
      </h2>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <ProjectCard
          v-for="project in featuredProjects"
          :key="project.id"
          :project="project"
          class="hover:-translate-y-1 transition-all duration-300"
        />
      </div>
      <NuxtLink to="/projects" class="section-link">
        view all projects →
      </NuxtLink>
    </section>
    
    <!-- Tech Stack -->
    <section class="mt-12 p-6 rounded-xl bg-bg-secondary/30 border border-accent/10">
      <h2 class="text-2xl font-bold mb-6 font-mono text-accent">
        $ tech stack
      </h2>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
        <div>
          <h3 class="text-lg font-bold mb-3 font-mono text-accent">Frontend</h3>
          <ul class="space-y-2 font-mono">
            <li class="flex items-center gap-2">
              <span class="text-accent">→</span>
              <span>Deno</span>
            </li>
            <li class="flex items-center gap-2">
              <span class="text-accent">→</span>
              <span>Nuxt 4</span>
            </li>
            <li class="flex items-center gap-2">
              <span class="text-accent">→</span>
              <span>Pinia</span>
            </li>
            <li class="flex items-center gap-2">
              <span class="text-accent">→</span>
              <span>TailwindCSS v4</span>
            </li>
            <li class="flex items-center gap-2">
              <span class="text-accent">→</span>
              <span>TypeScript</span>
            </li>
          </ul>
        </div>
        <div>
          <h3 class="text-lg font-bold mb-3 font-mono text-accent">Backend</h3>
          <ul class="space-y-2 font-mono">
            <li class="flex items-center gap-2">
              <span class="text-accent">→</span>
              <span>Go</span>
            </li>
            <li class="flex items-center gap-2">
              <span class="text-accent">→</span>
              <span>GORM</span>
            </li>
            <li class="flex items-center gap-2">
              <span class="text-accent">→</span>
              <span>Echo Framework v4</span>
            </li>
            <li class="flex items-center gap-2">
              <span class="text-accent">→</span>
              <span>SQLite</span>
            </li>
          </ul>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
import BlogPostCard from '~/components/Posts/BlogPostCard.vue'
import ArtCard from '~/components/Posts/ArtCard.vue'
import BskyPostCard from '~/components/Posts/BskyPostCard.vue'
import ProjectCard from '~/components/Posts/ProjectCard.vue'

// Sample data - in production, this would come from your API
const recentPosts = [
  { id: 1, title: 'Building a decentralized social network', date: '2024-01-15', category: 'Technology', summary: 'How I built my own social network using Crystal and I2P', slug: 'decentralized-social-network' },
  { id: 2, title: 'The ethics of AI art generation', date: '2024-01-10', category: 'Philosophy', summary: 'Exploring the moral implications of AI-generated artwork', slug: 'ai-art-ethics' },
  { id: 3, title: 'My favorite burrito recipe', date: '2024-01-05', category: 'Casual', summary: 'A step-by-step guide to making the perfect breakfast burrito', slug: 'burrito-recipe' }
]

const recentArt = [
  { id: 1, title: 'Digital Cat', image: '/images/art/cat-digital.jpg', tags: ['digital', 'cat', 'anime'] },
  { id: 2, title: 'Cyberpunk City', image: '/images/art/cyberpunk-city.jpg', tags: ['cyberpunk', 'cityscape', 'neon'] },
  { id: 3, title: 'Forest Spirit', image: '/images/art/forest-spirit.jpg', tags: ['fantasy', 'nature', 'spirit'] },
  { id: 4, title: 'Binary Dreams', image: '/images/art/binary-dreams.jpg', tags: ['abstract', 'code', 'digital'] }
]

const recentBsky = [
  { id: 1, text: 'Just finished implementing end-to-end encryption for my chat app. Privacy matters! #privacy #encryption', author: '@lilithinaparka.bsky.social', likes: 42, reposts: 12, timestamp: '2h ago' },
  { id: 2, text: 'New art piece: "Digital Ghost in the Machine". Exploring the intersection of AI and consciousness.', author: '@lilithinaparka.bsky.social', likes: 89, reposts: 23, timestamp: '1d ago' },
  { id: 3, text: 'Working on a new open-source shell written in Crystal. Going to call it CRSH. Stay tuned! #programming #crystal', author: '@aVaOk.bsky.social', likes: 56, reposts: 18, timestamp: '2d ago' }
]

const featuredProjects = [
  { id: 1, name: 'CRSH / Crush', description: 'Advanced POSIX-compliant shell written in Crystal', tags: ['crystal', 'shell', 'cli'], status: 'active', stars: 128 },
  { id: 2, name: 'blueberry-ib', description: 'Custom imageboard software written in Go', tags: ['go', 'imageboard', 'web'], status: 'active', stars: 89 },
  { id: 3, name: 'Vex', description: 'Voxel-based sandbox game with magic system', tags: ['game', 'voxel', 'sandbox'], status: 'development', stars: 45 }
]
</script>

<style scoped>
.home-page {
  @apply max-w-7xl mx-auto;
}

.hero-section {
  @apply pt-8 pb-12 border-b border-accent/20;
}

.stat-card {
  @apply p-4 rounded-lg bg-bg-secondary/50 border border-accent/10 
         text-center transition-all duration-300 hover:border-accent/30;
}

.stat-number {
  @apply text-3xl font-bold text-accent font-mono mb-1;
}

.stat-label {
  @apply text-text-secondary text-sm;
}

.content-section {
  @apply p-6 rounded-xl bg-bg-secondary/30 border border-accent/10;
}

.section-title {
  @apply text-xl font-bold mb-4 font-mono text-accent;
}

.section-link {
  @apply inline-block mt-4 text-sm text-link hover:text-accent 
         transition-colors font-mono;
}
</style>