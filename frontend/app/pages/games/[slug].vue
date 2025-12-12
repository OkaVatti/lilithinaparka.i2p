<template>
  <div class="game-page">
    <div v-if="gamesStore.loading" class="loading">
      <FeatherIcon name="loader" size="24" class="spin" />
      <span>Loading game...</span>
    </div>
    
    <div v-else-if="gamesStore.error" class="error">
      <FeatherIcon name="alert-circle" size="20" />
      <span>{{ gamesStore.error }}</span>
      <NuxtLink to="/games" class="btn mt-2">
        <FeatherIcon name="arrow-left" size="18" />
        <span>Back to Games</span>
      </NuxtLink>
    </div>
    
    <div v-else-if="game" class="game-container">
      <header class="game-header">
        <NuxtLink to="/games" class="back-link">
          <FeatherIcon name="arrow-left" size="18" />
          <span>Back to Games</span>
        </NuxtLink>
        
        <div class="game-info">
          <h1>{{ game.name }}</h1>
          <p class="game-description">{{ game.description }}</p>
          
          <div class="game-meta">
            <span class="meta-item">
              <FeatherIcon name="tag" size="16" />
              {{ game.category }}
            </span>
            <span class="meta-item">
              <FeatherIcon name="user" size="16" />
              {{ game.min_players }}-{{ game.max_players }} Players
            </span>
            <span v-if="game.multiplayer_supported" class="meta-item">
              <FeatherIcon name="users" size="16" />
              Multiplayer
            </span>
            <span class="meta-item">
              <FeatherIcon name="code" size="16" />
              v{{ game.version }}
            </span>
          </div>
        </div>
      </header>
      
      <div class="game-content">
        <div class="game-canvas-container">
          <canvas 
            ref="gameCanvas" 
            :width="canvasWidth" 
            :height="canvasHeight"
            @click="handleCanvasClick"
            @mousemove="handleCanvasMouseMove"
            @keydown="handleKeyDown"
            @keyup="handleKeyUp"
            tabindex="0"
          ></canvas>
          
          <div v-if="!gameStarted" class="game-overlay">
            <div class="start-screen">
              <h2>{{ game.name }}</h2>
              <p>{{ gameConfig?.instructions || 'Click to start' }}</p>
              <button @click="startGame" class="btn btn-large">
                <FeatherIcon name="play" size="20" />
                <span>Start Game</span>
              </button>
            </div>
          </div>
          
          <div v-if="gameOver" class="game-overlay">
            <div class="game-over-screen">
              <h2>Game Over!</h2>
              <p class="final-score">Score: {{ currentScore }}</p>
              
              <div v-if="isHighScore" class="high-score-message">
                <FeatherIcon name="award" size="24" />
                <p>New High Score!</p>
              </div>
              
              <div class="score-submit">
                <input 
                  v-model="playerAlias" 
                  type="text" 
                  placeholder="Enter your name"
                  maxlength="20"
                  @keyup.enter="submitScore"
                />
                <button @click="submitScore" class="btn" :disabled="submittingScore">
                  <FeatherIcon name="send" size="18" />
                  <span>{{ submittingScore ? 'Submitting...' : 'Submit Score' }}</span>
                </button>
              </div>
              
              <div class="game-over-actions">
                <button @click="restartGame" class="btn btn-primary">
                  <FeatherIcon name="refresh-cw" size="18" />
                  <span>Play Again</span>
                </button>
                <NuxtLink to="/games" class="btn btn-secondary">
                  <FeatherIcon name="arrow-left" size="18" />
                  <span>Back to Games</span>
                </NuxtLink>
              </div>
            </div>
          </div>
        </div>
        
        <aside class="game-sidebar">
          <div class="game-stats">
            <h3>Current Game</h3>
            <div class="stat">
              <span class="stat-label">Score:</span>
              <span class="stat-value">{{ currentScore }}</span>
            </div>
            <div class="stat">
              <span class="stat-label">Level:</span>
              <span class="stat-value">{{ currentLevel }}</span>
            </div>
            <div v-if="gameTime > 0" class="stat">
              <span class="stat-label">Time:</span>
              <span class="stat-value">{{ formatTime(gameTime) }}</span>
            </div>
          </div>
          
          <div v-if="gameConfig?.controls" class="game-controls">
            <h3>Controls</h3>
            <div class="control-list">
              <div v-for="(key, action) in gameConfig.controls" :key="action" class="control-item">
                <span class="control-action">{{ formatAction(action) }}:</span>
                <kbd class="control-key">{{ key }}</kbd>
              </div>
            </div>
          </div>
          
          <div class="game-actions">
            <button v-if="gameStarted && !gameOver" @click="pauseGame" class="btn">
              <FeatherIcon :name="gamePaused ? 'play' : 'pause'" size="18" />
              <span>{{ gamePaused ? 'Resume' : 'Pause' }}</span>
            </button>
            <button @click="restartGame" class="btn btn-secondary">
              <FeatherIcon name="refresh-cw" size="18" />
              <span>Restart</span>
            </button>
          </div>
          
          <div v-if="game.has_leaderboard" class="mini-leaderboard">
            <h3>Top Scores</h3>
            <div v-if="topScores.length > 0" class="score-list">
              <div v-for="(score, index) in topScores.slice(0, 5)" :key="score.id" class="score-item">
                <span class="rank">{{ index + 1 }}</span>
                <span class="alias">{{ score.alias }}</span>
                <span class="score">{{ score.score.toLocaleString() }}</span>
              </div>
            </div>
            <NuxtLink :to="`/games/${game.slug}/leaderboard`" class="view-all-link">
              View Full Leaderboard
            </NuxtLink>
          </div>
        </aside>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useGamesStore } from '~~/stores/games'
import type { GameScore } from '~~/types'

const route = useRoute()
const gamesStore = useGamesStore()
const { apiFetch } = useApi()

const gameCanvas = ref<HTMLCanvasElement | null>(null)
const game = computed(() => gamesStore.currentGame)
const gameConfig = ref<any>(null)

const gameStarted = ref(false)
const gameOver = ref(false)
const gamePaused = ref(false)
const currentScore = ref(0)
const currentLevel = ref(1)
const gameTime = ref(0)
const playerAlias = ref('')
const submittingScore = ref(false)
const isHighScore = ref(false)
const topScores = ref<GameScore[]>([])

const canvasWidth = ref(800)
const canvasHeight = ref(600)

let gameLoop: number | null = null
let gameState: any = null

const startGame = () => {
  gameStarted.value = true
  gameOver.value = false
  gamePaused.value = false
  currentScore.value = 0
  currentLevel.value = 1
  gameTime.value = 0
  
  initializeGame()
  startGameLoop()
  
  if (gameCanvas.value) {
    gameCanvas.value.focus()
  }
}

const pauseGame = () => {
  gamePaused.value = !gamePaused.value
  if (gamePaused.value) {
    stopGameLoop()
  } else {
    startGameLoop()
  }
}

const restartGame = () => {
  stopGameLoop()
  gameOver.value = false
  startGame()
}

const initializeGame = () => {
  // Initialize game-specific state based on game.slug
  const canvas = gameCanvas.value
  if (!canvas) return
  
  const ctx = canvas.getContext('2d')
  if (!ctx) return
  
  // Clear canvas
  ctx.fillStyle = '#000000'
  ctx.fillRect(0, 0, canvas.width, canvas.height)
  
  // Game-specific initialization would go here
  gameState = {
    entities: [],
    player: null,
    // Add game-specific state
  }
}

const startGameLoop = () => {
  let lastTime = Date.now()
  
  const loop = () => {
    if (gamePaused.value) return
    
    const currentTime = Date.now()
    const deltaTime = (currentTime - lastTime) / 1000
    lastTime = currentTime
    
    updateGame(deltaTime)
    renderGame()
    
    gameLoop = requestAnimationFrame(loop)
  }
  
  gameLoop = requestAnimationFrame(loop)
}

const stopGameLoop = () => {
  if (gameLoop) {
    cancelAnimationFrame(gameLoop)
    gameLoop = null
  }
}

const updateGame = (deltaTime: number) => {
  if (!gameState) return
  
  gameTime.value += deltaTime
  
  // Game-specific update logic would go here
  // This is a placeholder - actual game logic depends on the game type
}

const renderGame = () => {
  const canvas = gameCanvas.value
  if (!canvas) return
  
  const ctx = canvas.getContext('2d')
  if (!ctx) return
  
  // Clear canvas
  ctx.fillStyle = '#000000'
  ctx.fillRect(0, 0, canvas.width, canvas.height)
  
  // Game-specific rendering would go here
  // Draw game state, entities, etc.
}

const handleCanvasClick = (e: MouseEvent) => {
  if (!gameStarted.value || gameOver.value || gamePaused.value) return
  
  const rect = gameCanvas.value?.getBoundingClientRect()
  if (!rect) return
  
  const x = e.clientX - rect.left
  const y = e.clientY - rect.top
  
  // Handle click events in game
}

const handleCanvasMouseMove = (e: MouseEvent) => {
  if (!gameStarted.value || gameOver.value || gamePaused.value) return
  
  const rect = gameCanvas.value?.getBoundingClientRect()
  if (!rect) return
  
  const x = e.clientX - rect.left
  const y = e.clientY - rect.top
  
  // Handle mouse move events in game
}

const handleKeyDown = (e: KeyboardEvent) => {
  if (!gameStarted.value || gameOver.value) return
  
  // Prevent default for game keys
  if (['ArrowUp', 'ArrowDown', 'ArrowLeft', 'ArrowRight', ' '].includes(e.key)) {
    e.preventDefault()
  }
  
  // Handle key press in game
  if (e.key === 'Escape') {
    pauseGame()
  }
}

const handleKeyUp = (e: KeyboardEvent) => {
  if (!gameStarted.value || gameOver.value) return
  
  // Handle key release in game
}

const endGame = () => {
  gameOver.value = true
  stopGameLoop()
  
  // Check if it's a high score
  if (topScores.value.length > 0) {
    isHighScore.value = currentScore.value > topScores.value[0].score
  }
  
  // Load player alias from localStorage
  if (process.client) {
    const savedAlias = localStorage.getItem('player_alias')
    if (savedAlias) {
      playerAlias.value = savedAlias
    }
  }
}

const submitScore = async () => {
  if (!playerAlias.value.trim() || submittingScore.value) return
  
  submittingScore.value = true
  
  try {
    await gamesStore.submitScore(
      game.value!.slug,
      playerAlias.value,
      currentScore.value,
      currentLevel.value
    )
    
    // Save alias for next time
    if (process.client) {
      localStorage.setItem('player_alias', playerAlias.value)
    }
    
    // Refresh leaderboard
    await fetchTopScores()
    
    // Show success message
    alert('Score submitted successfully!')
  } catch (error) {
    console.error('Failed to submit score:', error)
    alert('Failed to submit score. Please try again.')
  } finally {
    submittingScore.value = false
  }
}

const fetchTopScores = async () => {
  if (!game.value) return
  
  try {
    const scores = await apiFetch<GameScore[]>(
      `/games/${game.value.slug}/leaderboard?limit=10`
    )
    topScores.value = scores
  } catch (error) {
    console.error('Failed to fetch top scores:', error)
  }
}

const formatTime = (seconds: number) => {
  const mins = Math.floor(seconds / 60)
  const secs = Math.floor(seconds % 60)
  return `${mins}:${secs.toString().padStart(2, '0')}`
}

const formatAction = (action: string) => {
  return action.charAt(0).toUpperCase() + action.slice(1).replace('_', ' ')
}

onMounted(async () => {
  const slug = route.params.slug as string
  await gamesStore.fetchGameBySlug(slug)
  
  if (game.value) {
    try {
      gameConfig.value = JSON.parse(game.value.config)
      
      if (gameConfig.value.canvas_width) {
        canvasWidth.value = gameConfig.value.canvas_width
      }
      if (gameConfig.value.canvas_height) {
        canvasHeight.value = gameConfig.value.canvas_height
      }
    } catch (error) {
      console.error('Failed to parse game config:', error)
    }
    
    if (game.value.has_leaderboard) {
      await fetchTopScores()
    }
  }
})

onUnmounted(() => {
  stopGameLoop()
})

useHead(() => ({
  title: `${game.value?.name || 'Game'} - Games`,
  meta: [
    { name: 'description', content: game.value?.description || '' }
  ]
}))
</script>

<style scoped>
.game-page {
  max-width: 1400px;
  margin: 0 auto;
}

.game-header {
  margin-bottom: 2rem;
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

.game-info h1 {
  font-size: 2.5rem;
  margin-bottom: 0.5rem;
  color: var(--theme-primary);
}

.game-description {
  font-size: 1.1rem;
  margin-bottom: 1rem;
  color: var(--theme-fg);
  opacity: 0.8;
}

.game-meta {
  display: flex;
  flex-wrap: wrap;
  gap: 1.5rem;
  font-size: 0.9rem;
  color: var(--theme-fg);
  opacity: 0.7;
}

.meta-item {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.game-content {
  display: grid;
  grid-template-columns: 1fr 300px;
  gap: 2rem;
}

.game-canvas-container {
  position: relative;
  background: #000;
  border: 2px solid var(--theme-border);
  border-radius: 8px;
  overflow: hidden;
}

canvas {
  display: block;
  width: 100%;
  height: auto;
  cursor: pointer;
}

canvas:focus {
  outline: 2px solid var(--theme-primary);
}

.game-overlay {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.9);
  display: flex;
  align-items: center;
  justify-content: center;
}

.start-screen,
.game-over-screen {
  text-align: center;
  padding: 2rem;
  max-width: 500px;
}

.start-screen h2,
.game-over-screen h2 {
  font-size: 2rem;
  margin-bottom: 1rem;
  color: var(--theme-primary);
}

.start-screen p {
  margin-bottom: 2rem;
  line-height: 1.6;
}

.final-score {
  font-size: 2.5rem;
  font-weight: bold;
  color: var(--theme-accent);
  margin-bottom: 1.5rem;
}

.high-score-message {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
  margin-bottom: 1.5rem;
  color: var(--theme-success);
  font-size: 1.2rem;
}

.score-submit {
  display: flex;
  gap: 0.5rem;
  margin-bottom: 1.5rem;
}

.score-submit input {
  flex: 1;
  padding: 0.75rem;
  background: var(--theme-bg);
  border: 1px solid var(--theme-border);
  border-radius: 4px;
  color: var(--theme-fg);
  font-family: inherit;
}

.score-submit input:focus {
  outline: none;
  border-color: var(--theme-primary);
}

.game-over-actions {
  display: flex;
  gap: 0.75rem;
  justify-content: center;
}

.game-sidebar {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

.game-stats,
.game-controls,
.mini-leaderboard {
  padding: 1.5rem;
  background: rgba(0, 0, 0, 0.2);
  border: 1px solid var(--theme-border);
  border-radius: 8px;
}

.game-stats h3,
.game-controls h3,
.mini-leaderboard h3 {
  margin-top: 0;
  margin-bottom: 1rem;
  color: var(--theme-primary);
}

.stat {
  display: flex;
  justify-content: space-between;
  padding: 0.75rem 0;
  border-bottom: 1px solid var(--theme-border);
}

.stat:last-child {
  border-bottom: none;
}

.stat-value {
  font-weight: bold;
  color: var(--theme-accent);
}

.control-list {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.control-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.control-key {
  padding: 0.25rem 0.5rem;
  background: rgba(189, 147, 249, 0.1);
  border: 1px solid var(--theme-primary);
  border-radius: 4px;
  font-family: 'Courier New', monospace;
  color: var(--theme-primary);
}

.game-actions {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.score-list {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
  margin-bottom: 1rem;
}

.score-item {
  display: grid;
  grid-template-columns: auto 1fr auto;
  gap: 0.5rem;
  padding: 0.5rem;
  background: rgba(189, 147, 249, 0.05);
  border-radius: 4px;
}

.rank {
  font-weight: bold;
  color: var(--theme-accent);
}

.alias {
  color: var(--theme-fg);
}

.score {
  font-weight: bold;
  color: var(--theme-primary);
}

.view-all-link {
  display: block;
  text-align: center;
  color: var(--theme-accent);
  font-size: 0.9rem;
}

@media (max-width: 1200px) {
  .game-content {
    grid-template-columns: 1fr;
  }
  
  .game-sidebar {
    grid-row: 1;
  }
}
</style>