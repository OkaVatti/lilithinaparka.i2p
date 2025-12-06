<!-- app/pages/projects/index.vue -->
<template>
  <div class="projects-page">
    <!-- Projects Header -->
    <div class="projects-header mb-8">
      <h1 class="text-4xl font-bold text-accent mb-4 font-mono">$ projects</h1>
      <p class="text-text-secondary font-mono">
        > my software projects, tools, and experiments
      </p>
      
      <!-- Projects Stats -->
      <div class="flex flex-wrap gap-4 mt-6">
        <div class="project-stat">
          <div class="stat-number">{{ totalProjects }}</div>
          <div class="stat-label">Active Projects</div>
        </div>
        <div class="project-stat">
          <div class="stat-number">{{ totalLanguages }}</div>
          <div class="stat-label">Languages Used</div>
        </div>
        <div class="project-stat">
          <div class="stat-number">{{ totalStars.toLocaleString() }}+</div>
          <div class="stat-label">Total Stars</div>
        </div>
      </div>
    </div>
    
    <!-- Status Filter -->
    <div class="flex flex-wrap gap-2 mb-8">
      <button
        v-for="status in statuses"
        :key="status"
        @click="setActiveStatus(status)"
        class="status-filter"
        :class="{ 'active': activeStatus === status }"
      >
        {{ status }}
      </button>
    </div>
    
    <!-- Projects Grid -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-12">
      <ProjectCard
        v-for="project in filteredProjects"
        :key="project.id"
        :project="project"
        :detailed="true"
      />
    </div>
    
    <!-- Games Section -->
    <div class="games-section mt-12">
      <h2 class="text-2xl font-bold mb-6 font-mono text-accent">
        $ games
      </h2>
      <p class="text-text-secondary mb-6 font-mono">
        > interactive games and experiments
      </p>
      
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <GameCard
          v-for="game in games"
          :key="game.id"
          :game="game"
        />
      </div>
      
      <NuxtLink to="/projects/games" class="inline-block mt-6 text-sm text-link hover:text-accent transition-colors font-mono">
        view all games →
      </NuxtLink>
    </div>
    
    <!-- Projects by Language -->
    <div class="mt-12 p-6 rounded-xl bg-bg-secondary/50 border border-accent/10">
      <h3 class="text-xl font-bold mb-4 font-mono text-accent">$ projects by language</h3>
      <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
        <div v-for="lang in languages" :key="lang.name" class="language-card">
          <div class="flex items-center gap-2 mb-2">
            <span class="text-lg" :class="lang.icon">{{ lang.icon }}</span>
            <span class="font-bold">{{ lang.name }}</span>
          </div>
          <div class="text-text-secondary text-sm">
            {{ lang.count }} project{{ lang.count !== 1 ? 's' : '' }}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import ProjectCard from '~/components/Posts/ProjectCard.vue'
import GameCard from '~/components/Posts/GameCard.vue'
import { ref, computed } from 'vue'

definePageMeta({
  layout: 'gamee'
})

// Sample projects data
const projects = [
  { id: 1, name: 'CRSH / Crush', description: 'Advanced POSIX-compliant shell written in Crystal', tags: ['crystal', 'shell', 'cli'], status: 'active', stars: 128, language: 'Crystal', github: 'https://github.com/okavatti/crsh' },
  { id: 2, name: 'blueberry-ib', description: 'Custom imageboard software written in Go', tags: ['go', 'imageboard', 'web'], status: 'active', stars: 89, language: 'Go', github: 'https://github.com/okavatti/blueberry-ib' },
  { id: 3, name: 'Vex', description: 'Voxel-based sandbox game with magic system', tags: ['game', 'voxel', 'sandbox'], status: 'development', stars: 45, language: 'C++', github: 'https://github.com/okavatti/vex' },
  { id: 4, name: 'bottletype', description: 'Simple markdown document editor built with Tauri', tags: ['tauri', 'markdown', 'editor'], status: 'active', stars: 67, language: 'Rust', github: 'https://github.com/okavatti/bottletype' },
  { id: 5, name: 'Catacomb', description: 'Locally hosted private credential manager', tags: ['security', 'password', 'local'], status: 'active', stars: 92, language: 'Go', github: 'https://github.com/okavatti/catacomb' },
  { id: 6, name: 'chaos.cr', description: 'Hashing algorithms with chaotic attractors', tags: ['cryptography', 'chaos', 'hashing'], status: 'active', stars: 56, language: 'Crystal', github: 'https://github.com/okavatti/chaos.cr' },
  { id: 7, name: '[enter]', description: 'A startpage for superusers', tags: ['startpage', 'productivity', 'web'], status: 'active', stars: 103, language: 'JavaScript', github: 'https://github.com/okavatti/enter' },
  { id: 8, name: 'lithium.cr', description: 'libsodium rewrite for Crystal', tags: ['cryptography', 'crystal', 'security'], status: 'maintenance', stars: 34, language: 'Crystal', github: 'https://github.com/okavatti/lithium.cr' },
  { id: 9, name: 'Litin', description: 'Init system designed for simplicity', tags: ['init', 'system', 'linux'], status: 'development', stars: 28, language: 'C', github: 'https://github.com/okavatti/litin' },
  { id: 10, name: 'Sxrf', description: 'Simple startpage for daily use', tags: ['startpage', 'minimal', 'web'], status: 'active', stars: 42, language: 'TypeScript', github: 'https://github.com/okavatti/sxrf' }
]

const games = [
  { id: 1, name: 'AstroClicker', description: 'Space-themed idle clicker game', category: 'Idle', tags: ['space', 'clicker', 'idle'], players: 'Single', status: 'playable' },
  { id: 2, name: 'CatRunner', description: 'ASCII-art dino runner clone', category: 'Arcade', tags: ['ascii', 'runner', 'cats'], players: 'Single', status: 'playable' },
  { id: 3, name: 'Chess', description: 'Multiplayer chess with AI', category: 'Strategy', tags: ['chess', 'ai', 'multiplayer'], players: '1-2', status: 'active' },
  { id: 4, name: 'Bones', description: 'Inscryption-like card game', category: 'Card', tags: ['cards', 'strategy', 'tabletop'], players: '1-2', status: 'development' }
]

const statuses = ['All', 'active', 'development', 'maintenance']
const activeStatus = ref('All')

const filteredProjects = computed(() => {
  if (activeStatus.value === 'All') return projects
  return projects.filter(project => project.status === activeStatus.value)
})

const totalProjects = computed(() => projects.length)
const totalStars = computed(() => projects.reduce((sum, project) => sum + project.stars, 0))

const languages = computed(() => {
  const langCounts: Record<string, number> = {}
  projects.forEach(project => {
    langCounts[project.language] = (langCounts[project.language] || 0) + 1
  })
  
  return Object.entries(langCounts).map(([name, count]) => ({
    name,
    count,
    icon: getLanguageIcon(name)
  })).sort((a, b) => b.count - a.count)
})

const totalLanguages = computed(() => languages.value.length)

const setActiveStatus = (status: string) => {
  activeStatus.value = status
}

const getLanguageIcon = (lang: string) => {
  const icons: Record<string, string> = {
    'Crystal': '💎',
    'Go': '🐹',
    'C++': '🔧',
    'Rust': '🦀',
    'C': '⚙️',
    'JavaScript': '📜',
    'TypeScript': '📘'
  }
  return icons[lang] || '📝'
}
</script>

<style scoped>
.projects-page {
  @apply max-w-7xl mx-auto;
}

.projects-header {
  @apply pb-8 border-b border-accent/20;
}

.project-stat {
  @apply px-6 py-3 rounded-lg bg-bg-secondary/30 border border-accent/10;
}

.project-stat .stat-number {
  @apply text-2xl font-bold text-accent font-mono;
}

.project-stat .stat-label {
  @apply text-text-secondary text-sm mt-1;
}

.status-filter {
  @apply px-4 py-2 rounded-full border border-accent/20 text-sm font-mono
         text-text-secondary hover:text-accent hover:border-accent/40
         transition-all duration-200;
}

.status-filter.active {
  @apply bg-accent/10 text-accent border-accent/40;
}

.games-section {
  @apply p-6 rounded-xl bg-bg-secondary/30 border border-accent/10;
}

.language-card {
  @apply p-4 rounded-lg bg-bg-secondary/30 border border-accent/5;
}
</style>