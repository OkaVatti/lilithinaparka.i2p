<!-- app/components/Posts/GameCard.vue -->
<template>
  <div class="game-card">
    <div class="p-4 rounded-lg border border-border bg-bg-secondary/30 hover:border-accent/30 transition-all duration-300">
      <div class="flex items-start justify-between mb-3">
        <h3 class="text-lg font-bold text-accent font-mono">{{ game.name }}</h3>
        <span class="game-category">
          {{ game.category }}
        </span>
      </div>
      <p class="text-text-secondary text-sm mb-4">{{ game.description }}</p>
      <div class="flex flex-wrap gap-1 mb-3">
        <span v-for="tag in game.tags.slice(0, 3)" :key="tag" class="game-tag">
          {{ tag }}
        </span>
        <span v-if="game.tags.length > 3" class="game-tag">+{{ game.tags.length - 3 }}</span>
      </div>
      <div class="flex items-center justify-between text-xs text-text-secondary font-mono">
        <div class="flex items-center gap-4">
          <span class="flex items-center gap-1">
            <span>👤</span>
            <span>{{ game.players }}</span>
          </span>
          <span class="game-status" :class="game.status">
            {{ game.status }}
          </span>
        </div>
        <button v-if="game.status === 'playable'" class="play-button">
          play →
        </button>
        <span v-else class="text-text-secondary">coming soon</span>
      </div>
      <div v-if="game.features && game.features.length > 0" class="mt-3 pt-3 border-t border-border">
        <ul class="space-y-1">
          <li v-for="feature in game.features.slice(0, 2)" :key="feature" class="text-xs text-text-secondary">
            • {{ feature }}
          </li>
          <li v-if="game.features.length > 2" class="text-xs text-text-secondary">
            • +{{ game.features.length - 2 }} more features
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  game: {
    id: number
    name: string
    description: string
    category: string
    tags: string[]
    players: string
    status: 'playable' | 'development' | 'active'
    features?: string[]
  }
}>()
</script>

<style scoped>
.game-category {
  @apply px-2 py-0.5 rounded-full bg-bg-primary text-xs text-text-secondary 
         border border-border;
}

.game-tag {
  @apply px-2 py-0.5 rounded-full bg-bg-primary/50 text-xs text-text-secondary 
         border border-border;
}

.game-status {
  @apply px-2 py-0.5 rounded-full text-xs;
}

.game-status.playable {
  @apply bg-green-500/10 text-green-400;
}

.game-status.development {
  @apply bg-yellow-500/10 text-yellow-400;
}

.game-status.active {
  @apply bg-blue-500/10 text-blue-400;
}

.play-button {
  @apply px-3 py-1 rounded border border-accent/20 text-accent text-xs
         hover:bg-accent/10 transition-colors;
}
</style>