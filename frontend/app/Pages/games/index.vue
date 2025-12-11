<template>
  <div class="games-page">
    <header class="page-header">
      <h1>Game Suite</h1>
      <p>A collection of retro-inspired browser games</p>
    </header>
    
    <section class="games-filters">
      <div class="filter-group">
        <label>Filter by Category:</label>
        <select v-model="selectedCategory" @change="filterGames">
          <option value="">All Categories</option>
          <option value="Dice">Dice</option>
          <option value="Puzzle">Puzzle</option>
          <option value="Match-3">Match-3</option>
          <option value="Arcade">Arcade</option>
        </select>
      </div>
      
      <div class="search-group">
        <input 
          v-model="searchQuery" 
          type="text" 
          placeholder="Search games..." 
          @input="filterGames"
        />
        <FeatherIcon name="search" size="18" />
      </div>
    </section>
    
    <section class="featured-game" v-if="featuredGame">
      <div class="featured-content">
        <h2>Featured: {{ featuredGame.name }}</h2>
        <p>{{ featuredGame.description }}</p>
        <NuxtLink :to="`/games/${featuredGame.slug}`" class="btn btn-large">
          <FeatherIcon name="play" size="20" />
          <span>Play Now</span>
        </NuxtLink>
      </div>
      <div class="featured-image">
        <img :src="getGameImage(featuredGame.slug)" :alt="featuredGame.name" />
      </div>
    </section>
    
    <section class="games-grid">
      <div v-if="gamesStore.loading" class="loading">
        <FeatherIcon name="loader" size="24" class="spin" />
        <span>Loading games...</span>
      </div>
      
      <div v-else-if="gamesStore.error" class="error">
        <FeatherIcon name="alert-circle" size="24" />
        <span>{{ gamesStore.error }}</span>
        <button @click="retryLoading" class="btn mt-2">
          Retry
        </button>
      </div>
      
      <template v-else>
        <GameCard
          v-for="game in filteredGames"
          :key="game.id"
          :game="game"
        />
      </template>
      
      <div v-if="!gamesStore.loading && filteredGames.length === 0" class="no-games">
        <FeatherIcon name="package" size="48" />
        <p>No games found matching your criteria.</p>
      </div>
    </section>
    
    <section v-if="recentScores.length > 0" class="recent-scores">
      <h2>Recent High Scores</h2>
      <div class="scores-list">
        <div v-for="score in recentScores" :key="score.id" class="score-item">
          <span class="score-alias">{{ score.alias }}</span>
          <span class="score-game">{{ getGameName(score.game_id) }}</span>
          <span class="score-value">{{ score.score.toLocaleString() }}</span>
          <span class="score-time">{{ formatTimeAgo(score.created_at) }}</span>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
import { useGamesStore } from '~~/stores/games'
import { useApi } from '~/composables/useApi'
import type { Game, GameScore } from '~~/types'

const gamesStore = useGamesStore()
const { apiFetch } = useApi()

const selectedCategory = ref('')
const searchQuery = ref('')
const recentScores = ref<GameScore[]>([])
const gamesMap = ref<Record<number, string>>({})

const games = computed(() => gamesStore.games)

const featuredGame = computed(() => {
  return games.value.find((game: { slug: string }) => game.slug === 'gems') || games.value[0]
})

const filteredGames = computed(() => {
  return games.value.filter((game: { category: string; name: string; description: string; tags: string }) => {
    const matchesCategory = !selectedCategory.value || game.category === selectedCategory.value
    const matchesSearch = !searchQuery.value || 
      game.name.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      game.description.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      game.tags.toLowerCase().includes(searchQuery.value.toLowerCase())
    
    return matchesCategory && matchesSearch
  })
})

const getGameImage = (slug: string) => {
  const images = {
    'bones': '/images/games/bones.png',
    'crypt': '/images/games/crypt.png',
    'gems': '/images/games/gems.png',
    'default': '/images/games/default.png'
  }
  return images[slug as keyof typeof images] || images.default
}

const getGameName = (gameId: number) => {
  return gamesMap.value[gameId] || 'Unknown Game'
}

const formatTimeAgo = (dateStr: string) => {
  const date = new Date(dateStr)
  const now = new Date()
  const diffMs = now.getTime() - date.getTime()
  const diffMins = Math.floor(diffMs / 60000)
  
  if (diffMins < 60) return `${diffMins}m ago`
  if (diffMins < 1440) return `${Math.floor(diffMins / 60)}h ago`
  return `${Math.floor(diffMins / 1440)}d ago`
}

const filterGames = () => {
  // Real-time filtering handled by computed property
}

const fetchRecentScores = async () => {
  try {
    const scores = await apiFetch<GameScore[]>('/games/scores/recent?limit=10')
    recentScores.value = scores
    
    // Create game name mapping
    games.value.forEach((game: { id: string | number; name: string }) => {
      gamesMap.value[game.id] = game.name
    })
  } catch (error) {
    console.error('Failed to fetch recent scores:', error)
  }
}

const retryLoading = () => {
  gamesStore.fetchGames()
  fetchRecentScores()
}

onMounted(async () => {
  await Promise.all([
    gamesStore.fetchGames(),
    fetchRecentScores()
  ])
})
</script>

<style scoped>
.page-header {
  text-align: center;
  margin-bottom: 3rem;
}

.page-header h1 {
  font-size: 3rem;
  margin-bottom: 0.5rem;
  color: var(--theme-primary);
}

.page-header p {
  font-size: 1.2rem;
  color: var(--theme-fg);
  opacity: 0.8;
}

.games-filters {
  display: flex;
  gap: 2rem;
  margin-bottom: 2rem;
  padding: 1.5rem;
  background: rgba(0, 0, 0, 0.2);
  border-radius: 8px;
  border: 2px solid var(--theme-border);
}

.filter-group, .search-group {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.filter-group label {
  color: var(--theme-fg);
  font-weight: bold;
}

.filter-group select,
.search-group input {
  padding: 0.5rem 1rem;
  background: var(--theme-bg);
  border: 1px solid var(--theme-border);
  border-radius: 4px;
  color: var(--theme-fg);
  font-family: inherit;
}

.search-group {
  flex: 1;
  position: relative;
}

.search-group input {
  width: 100%;
  padding-right: 2.5rem;
}

.search-group svg {
  position: absolute;
  right: 1rem;
  top: 50%;
  transform: translateY(-50%);
  color: var(--theme-fg);
  opacity: 0.6;
}

.featured-game {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 3rem;
  margin-bottom: 3rem;
  padding: 2rem;
  background: linear-gradient(135deg, rgba(139, 147, 233, 0.1) 0%, rgba(189, 147, 249, 0.1) 100%);
  border-radius: 12px;
  border: 2px solid var(--theme-primary);
}

.featured-content h2 {
  font-size: 2rem;
  margin-bottom: 1rem;
  color: var(--theme-primary);
}

.featured-content p {
  margin-bottom: 1.5rem;
  font-size: 1.1rem;
  line-height: 1.6;
}

.btn-large {
  padding: 1rem 2rem;
  font-size: 1.1rem;
}

.featured-image {
  border-radius: 8px;
  overflow: hidden;
  border: 3px solid var(--theme-border);
}

.featured-image img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.games-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 2rem;
  margin-bottom: 3rem;
}

.loading {
  grid-column: 1 / -1;
  text-align: center;
  padding: 3rem;
  color: var(--theme-accent);
}

.loading svg.spin {
  animation: spin 1s linear infinite;
  margin-bottom: 1rem;
}

@keyframes spin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}

.error {
  grid-column: 1 / -1;
  text-align: center;
  padding: 2rem;
}

.no-games {
  grid-column: 1 / -1;
  text-align: center;
  padding: 4rem 2rem;
  color: var(--theme-fg);
  opacity: 0.6;
}

.recent-scores {
  margin-top: 3rem;
  padding-top: 2rem;
  border-top: 2px solid var(--theme-border);
}

.recent-scores h2 {
  margin-bottom: 1.5rem;
  color: var(--theme-primary);
}

.scores-list {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.score-item {
  display: grid;
  grid-template-columns: 1fr 1fr auto auto;
  gap: 1rem;
  padding: 0.75rem 1rem;
  background: rgba(0, 0, 0, 0.2);
  border-radius: 4px;
  border: 1px solid var(--theme-border);
  transition: all 0.2s;
}

.score-item:hover {
  border-color: var(--theme-primary);
  background: rgba(189, 147, 249, 0.1);
}

.score-alias {
  font-weight: bold;
  color: var(--theme-primary);
}

.score-game {
  color: var(--theme-fg);
  opacity: 0.8;
}

.score-value {
  color: var(--theme-accent);
  font-weight: bold;
  min-width: 80px;
  text-align: right;
}

.score-time {
  color: var(--theme-fg);
  opacity: 0.6;
  font-size: 0.9rem;
  min-width: 80px;
  text-align: right;
}

@media (max-width: 768px) {
  .featured-game {
    grid-template-columns: 1fr;
    text-align: center;
  }
  
  .games-filters {
    flex-direction: column;
  }
  
  .score-item {
    grid-template-columns: 1fr 1fr;
    grid-template-rows: auto auto;
    gap: 0.5rem;
  }
  
  .score-value,
  .score-time {
    text-align: left;
  }
}
</style>