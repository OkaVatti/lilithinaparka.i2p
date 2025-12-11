<template>
  <div class="leaderboard-page">
    <header class="page-header">
      <NuxtLink :to="`/games/${slug}`" class="back-link">
        <FeatherIcon name="arrow-left" size="18" />
        <span>Back to Game</span>
      </NuxtLink>
      
      <h1>
        <FeatherIcon name="award" size="32" />
        {{ game?.name }} Leaderboard
      </h1>
    </header>
    
    <div v-if="gamesStore.loading" class="loading">
      <FeatherIcon name="loader" size="24" class="spin" />
      <span>Loading leaderboard...</span>
    </div>
    
    <div v-else-if="gamesStore.error" class="error">
      <FeatherIcon name="alert-circle" size="20" />
      <span>{{ gamesStore.error }}</span>
      <button @click="retry" class="btn mt-2">
        <FeatherIcon name="refresh-cw" size="18" />
        <span>Retry</span>
      </button>
    </div>
    
    <div v-else class="leaderboard-content">
      <div class="leaderboard-filters">
        <select v-model="timeFilter" @change="fetchLeaderboard">
          <option value="all">All Time</option>
          <option value="today">Today</option>
          <option value="week">This Week</option>
          <option value="month">This Month</option>
        </select>
        
        <button @click="fetchLeaderboard" class="btn btn-secondary">
          <FeatherIcon name="refresh-cw" size="18" />
          <span>Refresh</span>
        </button>
      </div>
      
      <div v-if="topThree.length > 0" class="podium">
        <div v-for="(score, index) in topThree" :key="score.id" class="podium-place" :class="`place-${index + 1}`">
          <div class="podium-rank">
            <FeatherIcon :name="getRankIcon(index)" size="32" />
            <span class="rank-number">{{ index + 1 }}</span>
          </div>
          <div class="podium-info">
            <h3>{{ score.alias }}</h3>
            <p class="score">{{ score.score.toLocaleString() }}</p>
            <p v-if="score.level > 1" class="level">Level {{ score.level }}</p>
          </div>
        </div>
      </div>
      
      <div class="leaderboard-table">
        <div class="table-header">
          <span class="col-rank">Rank</span>
          <span class="col-name">Player</span>
          <span class="col-score">Score</span>
          <span class="col-level">Level</span>
          <span class="col-date">Date</span>
        </div>
        
        <div v-if="remainingScores.length > 0" class="table-body">
          <div v-for="(score, index) in remainingScores" :key="score.id" class="table-row">
            <span class="col-rank">{{ index + 4 }}</span>
            <span class="col-name">{{ score.alias }}</span>
            <span class="col-score">{{ score.score.toLocaleString() }}</span>
            <span class="col-level">{{ score.level }}</span>
            <span class="col-date">{{ formatDate(score.created_at) }}</span>
          </div>
        </div>
        
        <div v-else class="no-scores">
          <FeatherIcon name="award" size="48" />
          <p>No scores yet. Be the first to play!</p>
        </div>
      </div>
      
      <div class="leaderboard-stats">
        <div class="stat-card">
          <h3>Total Players</h3>
          <p class="stat-value">{{ totalPlayers }}</p>
        </div>
        <div class="stat-card">
          <h3>Highest Score</h3>
          <p class="stat-value">{{ highestScore.toLocaleString() }}</p>
        </div>
        <div class="stat-card">
          <h3>Average Score</h3>
          <p class="stat-value">{{ Math.round(averageScore).toLocaleString() }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useGamesStore } from '~~/stores/games'

const route = useRoute()
const gamesStore = useGamesStore()

const slug = route.params.slug as string
const timeFilter = ref('all')

const game = computed(() => gamesStore.currentGame)
const scores = computed(() => gamesStore.leaderboard)

const topThree = computed(() => scores.value.slice(0, 3))
const remainingScores = computed(() => scores.value.slice(3))

const totalPlayers = computed(() => {
  const uniquePlayers = new Set(scores.value.map((s: { alias: any }) => s.alias))
  return uniquePlayers.size
})

const highestScore = computed(() => {
  return scores.value.length > 0 ? scores.value[0].score : 0
})

const averageScore = computed(() => {
  if (scores.value.length === 0) return 0
  const total = scores.value.reduce((sum: any, s: { score: any }) => sum + s.score, 0)
  return total / scores.value.length
})

const getRankIcon = (index: number) => {
  const icons = ['award', 'award', 'award']
  return icons[index]
}

const formatDate = (dateStr: string) => {
  const date = new Date(dateStr)
  const now = new Date()
  const diffDays = Math.floor((now.getTime() - date.getTime()) / (1000 * 60 * 60 * 24))
  
  if (diffDays === 0) return 'Today'
  if (diffDays === 1) return 'Yesterday'
  if (diffDays < 7) return `${diffDays}d ago`
  
  return date.toLocaleDateString('en-US', {
    month: 'short',
    day: 'numeric',
    year: date.getFullYear() !== now.getFullYear() ? 'numeric' : undefined
  })
}

const fetchLeaderboard = async () => {
  await Promise.all([
    gamesStore.fetchGameBySlug(slug),
    gamesStore.fetchLeaderboard(slug, 100)
  ])
}

const retry = () => {
  fetchLeaderboard()
}

onMounted(() => {
  fetchLeaderboard()
})

useHead(() => ({
  title: `${game.value?.name || 'Game'} Leaderboard - Games`,
  meta: [
    { name: 'description', content: `Top scores for ${game.value?.name || 'this game'}` }
  ]
}))
</script>

<style scoped>
.page-header {
  text-align: center;
  margin-bottom: 3rem;
}

.back-link {
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  margin-bottom: 1.5rem;
  padding: 0.5rem 1rem;
  border: 1px solid var(--theme-border);
  border-radius: 4px;
  color: var(--theme-accent);
  transition: all 0.2s;
}

.back-link:hover {
  border-color: var(--theme-primary);
  background: rgba(189, 147, 249, 0.1);
}

.page-header h1 {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 1rem;
  font-size: 2.5rem;
  color: var(--theme-primary);
  margin: 0;
}

.leaderboard-filters {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
  padding: 1rem;
  background: rgba(0, 0, 0, 0.2);
  border-radius: 8px;
  border: 1px solid var(--theme-border);
}

.leaderboard-filters select {
  padding: 0.5rem 1rem;
  background: var(--theme-bg);
  border: 1px solid var(--theme-border);
  border-radius: 4px;
  color: var(--theme-fg);
  font-family: inherit;
  cursor: pointer;
}

.podium {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 1rem;
  margin-bottom: 3rem;
}

.podium-place {
  padding: 2rem 1rem;
  background: rgba(0, 0, 0, 0.2);
  border: 2px solid var(--theme-border);
  border-radius: 8px;
  text-align: center;
  transition: all 0.2s;
}

.podium-place:hover {
  transform: translateY(-4px);
}

.podium-place.place-1 {
  border-color: #ffd700;
  background: rgba(255, 215, 0, 0.1);
  grid-row: 1;
  order: 2;
}

.podium-place.place-2 {
  border-color: #c0c0c0;
  background: rgba(192, 192, 192, 0.1);
  order: 1;
}

.podium-place.place-3 {
  border-color: #cd7f32;
  background: rgba(205, 127, 50, 0.1);
  order: 3;
}

.podium-rank {
  margin-bottom: 1rem;
}

.podium-place.place-1 .podium-rank svg {
  color: #ffd700;
}

.podium-place.place-2 .podium-rank svg {
  color: #c0c0c0;
}

.podium-place.place-3 .podium-rank svg {
  color: #cd7f32;
}

.rank-number {
  display: block;
  font-size: 1.5rem;
  font-weight: bold;
  margin-top: 0.5rem;
}

.podium-info h3 {
  margin: 0 0 0.5rem 0;
  font-size: 1.2rem;
  color: var(--theme-primary);
}

.podium-info .score {
  font-size: 1.8rem;
  font-weight: bold;
  color: var(--theme-accent);
  margin: 0.5rem 0;
}

.podium-info .level {
  font-size: 0.9rem;
  color: var(--theme-fg);
  opacity: 0.7;
  margin: 0;
}

.leaderboard-table {
  margin-bottom: 3rem;
}

.table-header,
.table-row {
  display: grid;
  grid-template-columns: 60px 1fr 120px 80px 120px;
  gap: 1rem;
  padding: 1rem;
  align-items: center;
}

.table-header {
  background: rgba(189, 147, 249, 0.2);
  border: 1px solid var(--theme-primary);
  border-radius: 8px 8px 0 0;
  font-weight: bold;
  color: var(--theme-primary);
}

.table-row {
  border-bottom: 1px solid var(--theme-border);
  transition: all 0.2s;
}

.table-row:hover {
  background: rgba(189, 147, 249, 0.05);
}

.col-rank {
  text-align: center;
  font-weight: bold;
}

.col-score {
  text-align: right;
  font-weight: bold;
  color: var(--theme-accent);
}

.col-level {
  text-align: center;
}

.col-date {
  text-align: right;
  font-size: 0.9rem;
  opacity: 0.7;
}

.no-scores {
  text-align: center;
  padding: 4rem 2rem;
  color: var(--theme-fg);
  opacity: 0.6;
}

.leaderboard-stats {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 1rem;
}

.stat-card {
  padding: 1.5rem;
  background: rgba(0, 0, 0, 0.2);
  border: 1px solid var(--theme-border);
  border-radius: 8px;
  text-align: center;
}

.stat-card h3 {
  margin: 0 0 0.5rem 0;
  font-size: 0.9rem;
  color: var(--theme-fg);
  opacity: 0.7;
  text-transform: uppercase;
  letter-spacing: 0.05em;
}

.stat-value {
  font-size: 2rem;
  font-weight: bold;
  color: var(--theme-primary);
  margin: 0;
}

@media (max-width: 768px) {
  .podium {
    grid-template-columns: 1fr;
  }
  
  .podium-place {
    order: initial !important;
  }
  
  .table-header,
  .table-row {
    grid-template-columns: 50px 1fr 80px;
    font-size: 0.9rem;
  }
  
  .col-level,
  .col-date {
    display: none;
  }
}
</style>