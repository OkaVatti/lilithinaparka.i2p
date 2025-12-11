<template>
  <article class="game-card card">
    <div class="game-image">
      <img 
        :src="gameImage" 
        :alt="game.name"
        @error="handleImageError"
      />
      <div v-if="game.multiplayer_supported" class="multiplayer-badge">
        <FeatherIcon name="users" size="14" />
        <span>Multiplayer</span>
      </div>
    </div>
    
    <div class="game-content">
      <h3 class="game-title">{{ game.name }}</h3>
      <p class="game-description">{{ game.description }}</p>
      
      <div class="game-meta">
        <span class="game-category">{{ game.category }}</span>
        <span class="players">
          <FeatherIcon name="user" size="14" />
          {{ game.min_players }}{{ game.max_players > 1 ? `-${game.max_players}` : '' }}
        </span>
        <span v-if="game.version" class="version">v{{ game.version }}</span>
      </div>
      
      <div class="game-actions">
        <NuxtLink :to="`/games/${game.slug}`" class="btn btn-primary">
          <FeatherIcon name="play" size="16" />
          <span>Play Now</span>
        </NuxtLink>
        
        <NuxtLink v-if="game.has_leaderboard" :to="`/games/${game.slug}/leaderboard`" class="btn btn-secondary">
          <FeatherIcon name="award" size="16" />
          <span>Leaderboard</span>
        </NuxtLink>
      </div>
    </div>
  </article>
</template>

<script setup lang="ts">
import type { Game } from '~~/types'

const props = defineProps<{
  game: Game
}>()

const gameImage = computed(() => {
  const images = {
    'bones': '/images/games/bones.png',
    'crypt': '/images/games/crypt.png',
    'gems': '/images/games/gems.png',
    'default': '/images/games/default.png'
  }
  
  return images[props.game.slug as keyof typeof images] || images.default
})

const handleImageError = (e: Event) => {
  const img = e.target as HTMLImageElement
  img.src = '/images/games/default.png'
}
</script>

<style scoped>
.game-card {
  transition: all 0.3s;
  overflow: hidden;
}

.game-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 8px 16px var(--theme-shadow);
}

.game-image {
  position: relative;
  width: 100%;
  height: 200px;
  overflow: hidden;
  background: linear-gradient(135deg, var(--theme-primary) 0%, var(--theme-secondary) 100%);
}

.game-image img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.3s;
}

.game-card:hover .game-image img {
  transform: scale(1.05);
}

.multiplayer-badge {
  position: absolute;
  top: 1rem;
  right: 1rem;
  display: flex;
  align-items: center;
  gap: 0.25rem;
  padding: 0.25rem 0.5rem;
  background: rgba(0, 0, 0, 0.8);
  border-radius: 4px;
  color: white;
  font-size: 0.8rem;
}

.game-content {
  padding: 1.5rem;
}

.game-title {
  margin: 0 0 0.5rem 0;
  font-size: 1.5rem;
  color: var(--theme-primary);
}

.game-description {
  margin: 0 0 1rem 0;
  color: var(--theme-fg);
  opacity: 0.8;
  line-height: 1.6;
}

.game-meta {
  display: flex;
  gap: 1rem;
  margin-bottom: 1.5rem;
  font-size: 0.9rem;
  color: var(--theme-fg);
  opacity: 0.6;
}

.game-meta span {
  display: flex;
  align-items: center;
  gap: 0.25rem;
}

.game-category {
  text-transform: uppercase;
  font-weight: bold;
  letter-spacing: 0.05em;
}

.game-actions {
  display: flex;
  gap: 0.75rem;
}

.game-actions .btn {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
}

.btn-primary {
  background: var(--theme-primary);
  color: var(--theme-bg);
}

.btn-secondary {
  background: var(--theme-border);
  color: var(--theme-fg);
}

@media (max-width: 768px) {
  .game-actions {
    flex-direction: column;
  }
}
</style>