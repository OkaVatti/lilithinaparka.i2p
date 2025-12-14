<template>
  <div class="game-card window">
    <div class="title-bar">
      <div class="title-bar-text">{{ game.name }}</div>
    </div>
    <div class="window-body">
      <div class="game-info">
        <p class="description">{{ game.description }}</p>
        
        <div class="game-meta">
          <div class="status-bar">
            <p class="status-bar-field">{{ game.category }}</p>
            <p class="status-bar-field">
              {{ game.min_players }}-{{ game.max_players }} Players
            </p>
          </div>
        </div>
        
        <div class="game-actions">
          <NuxtLink :to="`/games/${game.slug}`">
            <button>▶ Play Now</button>
          </NuxtLink>
          
          <NuxtLink v-if="game.has_leaderboard" :to="`/games/${game.slug}/leaderboard`">
            <button>🏆 Leaderboard</button>
          </NuxtLink>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Game } from '~~/types'

defineProps<{
  game: Game
}>()
</script>

<style scoped>
.game-card {
  height: 100%;
  display: flex;
  flex-direction: column;
}

.window-body {
  flex: 1;
  display: flex;
  flex-direction: column;
}

.game-info {
  flex: 1;
  display: flex;
  flex-direction: column;
}

.description {
  flex: 1;
  margin: 0 0 1rem 0;
  font-size: 0.9rem;
}

.game-meta {
  margin-bottom: 1rem;
}

.status-bar {
  display: flex;
  gap: 0.5rem;
}

.status-bar-field {
  margin: 0;
  padding: 0.25rem 0.5rem;
  font-size: 0.85rem;
}

.game-actions {
  display: flex;
  gap: 0.5rem;
  flex-wrap: wrap;
}

.game-actions button {
  flex: 1;
  min-width: 120px;
}
</style>