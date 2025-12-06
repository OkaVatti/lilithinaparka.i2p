<!-- app/components/Posts/ProjectCard.vue -->
<template>
  <a :href="project.github" target="_blank" class="project-card block" v-if="project.github">
    <div class="p-4 rounded-lg border border-border bg-bg-secondary/30 hover:border-accent/30 transition-all duration-300">
      <div class="flex items-start justify-between mb-3">
        <h3 class="text-lg font-bold text-accent font-mono">{{ project.name }}</h3>
        <span class="project-status" :class="project.status">
          {{ project.status }}
        </span>
      </div>
      <p class="text-text-secondary text-sm mb-4">{{ project.description }}</p>
      <div class="flex flex-wrap gap-1 mb-3">
        <span v-for="tag in project.tags.slice(0, 3)" :key="tag" class="project-tag">
          {{ tag }}
        </span>
        <span v-if="project.tags.length > 3" class="project-tag">+{{ project.tags.length - 3 }}</span>
      </div>
      <div class="flex items-center justify-between text-xs text-text-secondary font-mono">
        <div class="flex items-center gap-2">
          <span class="project-language">{{ project.language }}</span>
          <span>·</span>
          <span class="flex items-center gap-1">
            <span>★</span>
            <span>{{ project.stars }}</span>
          </span>
        </div>
        <span class="project-link">view on github →</span>
      </div>
    </div>
  </a>
  <div class="project-card" v-else>
    <div class="p-4 rounded-lg border border-border bg-bg-secondary/30">
      <div class="flex items-start justify-between mb-3">
        <h3 class="text-lg font-bold text-accent font-mono">{{ project.name }}</h3>
        <span class="project-status" :class="project.status">
          {{ project.status }}
        </span>
      </div>
      <p class="text-text-secondary text-sm mb-4">{{ project.description }}</p>
      <div class="flex flex-wrap gap-1">
        <span v-for="tag in project.tags.slice(0, 3)" :key="tag" class="project-tag">
          {{ tag }}
        </span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  project: {
    id: number
    name: string
    description: string
    tags: string[]
    status: 'active' | 'development' | 'maintenance'
    stars?: number
    language?: string
    github?: string
  }
}>()
</script>

<style scoped>
.project-status {
  @apply px-2 py-0.5 rounded-full text-xs font-mono;
}

.project-status.active {
  @apply bg-green-500/10 text-green-400 border border-green-500/20;
}

.project-status.development {
  @apply bg-yellow-500/10 text-yellow-400 border border-yellow-500/20;
}

.project-status.maintenance {
  @apply bg-blue-500/10 text-blue-400 border border-blue-500/20;
}

.project-tag {
  @apply px-2 py-0.5 rounded-full bg-bg-primary/50 text-xs text-text-secondary 
         border border-border;
}

.project-language {
  @apply px-2 py-0.5 rounded bg-bg-primary text-text-secondary;
}

.project-link {
  @apply text-link hover:text-accent transition-colors;
}
</style>